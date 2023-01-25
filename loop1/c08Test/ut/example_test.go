package ut

import "testing"

func TestSomething(t *testing.T) {
	t.Log("something")
	t.Errorf("something error")
	t.Fatal("fatal")
	t.Log("this line can not be reached below `t.Fatal()`")
}
