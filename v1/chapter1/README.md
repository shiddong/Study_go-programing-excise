## Chapter1 Go开发环境搭建与Hello world

### Go 的安装
一般有几种方法安装：
1. Go 源码安装
2. Go 标准包安装
3. 第三方工具安装，比如mac上的 brew 也可安装

### Go env
1. GOPATH: 工作目录
2. GOROOT: 安装目录
3. GOBIN: 工作目录下的bin目录
4. GOEXE: 生成的可执行文件的后缀，mac和linux下为空字符串，windows下为".exe"

根据约定，GOPATH目录下需要建立3个目录：
1. bin: 用来存放编译后生成的可执行文件
2. pkg: 存放编译后生成的包文件
3. src: 项目源码所在目录

### Go 常用命令
1. go get: 获取远程包，如果从github上下载包，需要先安装git
2. go run: 直接运行程序
3. go build: 测试编译，检查是否有编译错误
4. go fmt: 格式化源码，有些编译器比如 VS Code 会自动启动
5. go install: 编译包文件并编译整个程序
6. go test: 运行测试程序，测试程序的默认格式是 HelloWorld_test.go这个以 _test.go结尾的文件
7. go doc/godoc: 查看文档，比如本地搭建一个文档服务器：godoc -http=localhost:8080

### Go 程序的项目结构