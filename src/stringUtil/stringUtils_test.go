package stringUtil

import (
    "testing"
)

func TestIsNum(t *testing.T) {
	t.Log(IsNum("sfd"))
	t.Log(IsNum("123"))
	t.Log(IsNum("123.3"))
}

func TestSplit(t *testing.T) {
	for _,s:=range Split("sfd,,ff,ffs",","){
		t.Log(s)
	}
}
