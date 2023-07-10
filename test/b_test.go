package test

import (
	"fmt"
	digui "github.com/liweiww/test/t"
	"testing"
)

func TestName(t *testing.T) {
	digui.Print()
}

func TestSay(t *testing.T) {
	digui.Say()
	digui.Dog()
	digui.Dog()
	digui.Dog()
	fmt.Print("dadsa")
}
