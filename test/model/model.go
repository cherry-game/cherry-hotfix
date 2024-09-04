package model

import (
	"github.com/cherry-game/cherry-hotfix/test/model/m1"
)

type Foo struct {
	String string
	M1Int  m1.M1
}

func (p *Foo) Hello() string {
	return p.String
}

func (p *Foo) ExecFibonacciSum(n int) int {
	return FibonacciSum(n)
}

func FibonacciSum(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	var a, b int = 0, 1
	for i := 2; i <= n; i++ {
		b = a + b
		a = b - a //将b的值还原城之前的值
		b = b % 1000000007
	}
	return b
}
