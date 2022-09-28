package test

import (
	"fmt"
	"testing"
)

//当需要终止当前测试用例时，可以使用 FailNow
func TestFailNow(t *testing.T) {
	fmt.Println("before fail")
	t.FailNow()
	//println("after fail")		//该语句因 FailNow 终止运行而不会运行, 不注释会有警告
}

//当出现错误只标记不终止时，可以用 Fail
func TestFail(t *testing.T) {
	fmt.Println("before fail")
	t.Fail()
	fmt.Println("after fail")
}

/*
命令：go test -v fail_test.go	// go test 指定文件时默认执行文件内的所有测试用例
输出：
=== RUN   TestFailNow
before fail
--- FAIL: TestFailNow (0.00s)
=== RUN   TestFail
before fail
after fail
--- FAIL: TestFail (0.00s)
FAIL
FAIL    command-line-arguments  0.364s
FAIL
*/