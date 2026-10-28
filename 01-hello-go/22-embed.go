package main

import "fmt"

type base struct {
    num int
}

// 注意：这里用的是「值接收者」。
// 值接收者 → base 和 *base 都拥有 describe 方法
// （通过嵌入提升后，container 和 *container 也都拥有）
func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

// container 匿名嵌入了 base（只写类型名，不写字段名）
// 效果：base 的字段和方法会被「提升（promoted）」到 container 上，
// 所以 co.num、co.describe() 看起来就像 container 自己拥有的一样。
type container struct {
    base
    str string
}

func main() {

    co := container{
        base: base{
            num: 1,
        },
        str: "some name",
    }

    fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

    fmt.Println("also num:", co.base.num)

    fmt.Println("describe:", co.describe())

    // ============================================================
    // 接口（interface）：定义一份「方法清单」，实现多态
    // ------------------------------------------------------------
    // ① type describer interface { describe() string }
    //    接口只声明方法签名，不写实现。
    //    任何类型只要有 describe() string，就「属于」这个接口。
    //
    //    这里写在 main 函数内部 → 局部类型，只在 main 内可见
    //    （Go 允许函数内声明类型；实际项目里接口一般定义在包级，
    //     方便复用和被其他类型实现）
    //
    // ② var d describer = co
    //    Go 的接口是「隐式实现」：不需要 implements 关键字。
    //    只要 container 的方法集里有 describe() string，编译器就自动
    //    认为它实现了接口 —— 结构化类型 / 鸭子类型，且是编译期检查。
    //
    //    container 自己没写 describe，但因为嵌入了 base，
    //    base 的 describe() 被提升到了 container 上，所以满足接口。
    //
    //    若不满足，编译期直接报错（不是运行时才炸）：
    //      cannot use co (variable of type container) as describer value:
    //      missing method describe
    //
    // ③ d.describe()
    //    接口变量 d 内部是一个二元组 (动态类型, 动态值)，
    //    此时为 (container, co 的副本)。
    //    调用时根据动态类型找到真正的方法执行 —— 即动态分发（多态）。
    // ============================================================
    type describer interface {
        describe() string
    }

    var d describer = co
    fmt.Println("describer:", d.describe())

    // ------------------------------------------------------------
    // 与「值接收者 / 指针接收者」的关联（很关键）：
    //   第 10 行的 describe 是值接收者 → container 和 *container 都实现接口
    //   var d describer = co     ✅
    //   var d describer = &co    ✅
    //
    //   如果改成指针接收者 func (b *base) describe() string：
    //   只有 *container 拥有该方法，值不满足接口
    //   var d describer = co     ❌ 编译失败
    //   var d describer = &co    ✅
    // ------------------------------------------------------------

    // ------------------------------------------------------------
    // 接口的实际用途：解耦。函数只依赖「能 describe」这件事，
    // 不关心具体是什么类型。
    //
    //   func printIt(d describer) {
    //       fmt.Println(d.describe())
    //   }
    //   printIt(co)        // container 可以
    //   printIt(base{2})   // base 自己也可以
    //
    // 把接口还原成具体类型（类型断言 / 类型分支）：
    //   c := d.(container)          // 断言，类型不符会 panic
    //   c, ok := d.(container)      // 安全写法，ok=false 表示不是该类型
    //   switch v := d.(type) {      // 类型分支
    //   case container:
    //   case base:
    //   }
    // ------------------------------------------------------------
}

// ============================================================
// JS / TS 对照
//   - JS 没有接口，全靠约定（鸭子类型），且没有编译期检查，
//     传错对象要等运行时才报 d.describe is not a function
//   - TS 的 interface 是结构化的，跟 Go 最接近；
//     区别是 Go 完全不需要 implements 关键字
// ============================================================
