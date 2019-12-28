go test 需要加上引用的源文件

`go test -v link_list_test.go link_list.go`

安装包

```bash
GOPATH=current_project_dir go get github.com/stretchr/testify/assert
```

#### 性能评估

`go test -bench=.`

bench的值：

- 表示运行所有benchmark函数
- 也可以直接写benchmark函数的名字，指定要运行benchmark的函数,如: bench=BenchmarkConcatStringByAdd

研究代码块有多少次内存分配：

`go test -bench=. -benchmem`

- -bench=<相关benchmark测试>
- Windows 下使用go test命令行时， -bench=. 应写为 -bench="."

#### BDD 
让业务领域的专家参与开发
你知道Story Card背面应该写什么吗？
背面应该写该Stroy应该如何被验收
用业务领域的语言来描述
"Given - When - Then"
Given a user is creating an account
When they specify an insecure password
Then they see a message, "Passwords must be at least 6 characters long with at least one letter, one number, and one symbol."

启动 WEB UI
`$GOPATH/bin/goconvey`

#### 反射

提高了程序的灵活性
降低了程序的可读性
降低了程序的性能