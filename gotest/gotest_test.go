package gotest

import "testing"

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("除法函数测试没通过")
		return
	}
	t.Log("第一个测试通过了")
}

func Test_Division_2(t *testing.T) {
	t.Error("就是不通过")
}

func TestAdd(t *testing.T) {
	sum := Add(2, 3)
	if sum == 5 {
		t.Log("测试通过")
		return
	}
	t.Error("测试失败")
}
