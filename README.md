# Go 学习仓库（7 天上手）

环境：Go 1.27.3 darwin/arm64，GOROOT=`~/sdk/go`，GOPATH=`~/go`

## 已就绪

- [x] Go 1.27.3 arm64（官方 tarball，SHA256 已校验）
- [x] PATH / GOPATH / GOPROXY 写入 `~/.zshrc`
- [x] `hello-world.go` 跑通，`02-syntax` 语法对照示例跑通
- [x] gopls v0.24.0 + dlv 已安装（`~/go/bin`，Cursor/VS Code 的 Go 扩展自动识别）

## 7 天进度

> 已按目标 JD（教育科技 · 全栈，Go 为主 / Python 为辅）对齐：Day2 框架锁定 Gin；Day3 加 MySQL 索引 + Redis 缓存三件套；Day6 项目选教育 H5 场景服务。

| Day | 主题 | 内容 | 验收 | 状态 |
|---|---|---|---|---|
| 1 | 语法 + 工具链 | Go by Example 中文版手敲一遍；`go mod/build/run/test/fmt/vet` | CLI 小工具 | ⬜ |
| 2 | HTTP | `net/http` 手写路由 + 中间件，再上 Gin | 3 个路由跑通 | ⬜ |
| 3 | 存储 | sqlx/GORM + MySQL，go-redis，事务、连接池、ctx 传参 | 完整 CRUD + 缓存 | ⬜ |
| 4 | 并发 | goroutine / channel / select / WaitGroup / errgroup | 并发抓取器 | ⬜ |
| 5 | 工程化 | `%w` 包装、zap 日志、viper 配置、testify 单测 | 服务项目补齐三件套 | ⬜ |
| 6 | 实战 | 完整项目一口气写完 | 可跑项目 | ⬜ |
| 7 | 上线 | Docker 多阶段构建、优雅退出、pprof | 线上可访问 | ⬜ |

## 常用命令速查

```bash
go mod init <module>     # 初始化模块（每个项目一次）
go run .                 # 直接跑
go build -o app .        # 编译产物
go test ./...            # 跑测试
go vet ./...             # 静态检查（必跑）
go fmt ./...             # 格式化（Go 官方风格统一，别争缩进）
go mod tidy              # 整理依赖
go get <pkg>             # 加依赖
go tool pprof ...        # 性能分析
```

## 中文资源（已验证可访问）

| 站点 | 用途 |
|---|---|
| https://gobyexample-cn.github.io | 语法速刷，半天刷完（Day 1 主力） |
| https://liwenzhou.com/posts/go/golang-menu | 中文最系统入门，配 B 站视频 |
| https://www.topgoer.com | 标准库速查字典 |
| https://geektutu.com | 7 天实现 Web 框架 Gee / ORM（Day 4-7 主力） |
| https://studygolang.com/pkgdoc | 标准库中文文档 |
| https://draveness.me/golang | 《Go 语言设计与实现》，第 2-3 周进阶 |

## 目录

- `hello-world.go` — 环境验证（你的第一个 Go 程序）
- `02-syntax/` — 语法对照示例（struct/interface/error/goroutine/channel/context，含 Node 概念注释）
