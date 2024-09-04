package foo

import (
	"reflect"

	"github.com/cherry-game/cherry-hotfix/hotfix"
	"github.com/cherry-game/cherry-hotfix/test/model"
)

func GetPatch() *hotfix.FuncPatch {
	// fmt.Println("[Patch] invoke GetPatch()")

	fn := func(foo *model.Foo, n int) int {
		foo.M1Int.Int = 1
		// fmt.Println("======FibonacciSum")
		return foo.HotFibonacciSum(n)
	}

	return &hotfix.FuncPatch{
		StructType: reflect.TypeOf(&model.Foo{}),
		// FuncName:   "ExecFibonacciSum",
		FuncValue: reflect.ValueOf(fn),
	}
}
