package test

import (
	"fmt"
	"github.com/Ming0706/base"
	"testing"
)

func TestAdd(t *testing.T) {
	r := base.Add(1, 2)
	if r !=3 {

		t.Errorf("Add(1, 2) failed. Got %d, expected 3.", r)
	}

}

func TestMain(m *testing.M) {
	fmt.Println("begin")
	m.Run()
	fmt.Println("end")
}
