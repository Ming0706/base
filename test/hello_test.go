package test

import "testing"

func TestHelloWorld(t *testing.T) {
	println("hello test...")
	t.Log("hello world")
}

/*
命令：go test hello_test.go
输出：
ok      command-line-arguments  0.054s

命令：go test -v hello_test.go		//显示在附加参数中添加了-v, 可以让测试时显示详细的流程
输出：
=== RUN   TestHelloWorld
hello test...
    hello_test.go:7: hello world	//文件名:代码处于第几行:内容
--- PASS: TestHelloWorld (0.00s)
PASS
ok      command-line-arguments  0.065s
*/