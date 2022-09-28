package test

import "testing"

func TestA(t *testing.T) {
	t.Log("A")
}
func TestAK(t *testing.T) {
	t.Log("AK47")
}
func TestB(t *testing.T) {
	t.Log("B")
}
func TestC(t *testing.T) {
	t.Log("C")
}

/*
命令: go test -v -run TestA select_test.go	//使用 -run 参数选择需要的测试用例单独执行
输出:
=== RUN   TestA
    select_test.go:6: A
--- PASS: TestA (0.00s)
=== RUN   TestAK
    select_test.go:9: AK47		//TestA 和 TestAK 的测试用例都被执行，原因是-run跟随的测试用例的名称支持正则表达式
--- PASS: TestAK (0.00s)
PASS
ok      command-line-arguments  0.348s

命令: go test -v -run TestA$ select_test.go  // 使用正则表达式使只执行 TestA 的测试用例
输出:
=== RUN   TestA
    select_test.go:6: A
--- PASS: TestA (0.00s)
PASS
ok      command-line-arguments  0.052s
*/