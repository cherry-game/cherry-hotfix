package hotfix_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cherry-game/cherry-hotfix/hotfix"
	"github.com/cherry-game/cherry-hotfix/test/model"

	"github.com/cherry-game/cherry-hotfix/symbols"
)

func TestFixFooHelloFunc(t *testing.T) {
	foo := &model.Foo{
		String: "foo",
	}

	// 初始执行Hello(),并打印结果
	fmt.Printf("[Init]  foo:{%p}, Hello():{%v}\n", foo, foo.Hello())

	// 模拟Hello()被调用
	for i := 0; i < 1000; i++ {
		go func(foo *model.Foo) {
			for {
				foo.Hello()
				time.Sleep(1 * time.Millisecond)
			}
		}(foo)
	}

	var (
		filePath = "./_patch_files/foo.go.patch" // 补丁脚本的路径
		evalText = "foo.GetPatch()"              // 补丁脚本内执行的函数名
	)

	// 加载补丁函数foo.GetPatch()
	patches, err := hotfix.ApplyFunc(filePath, evalText, symbols.Symbols)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印已被替换的foo.Hello()
	fmt.Printf("[Patch] foo:{%p}, Hello():{%v}\n", foo, foo.Hello())

	// 执行重置
	patches.Reset()

	// 打印函数
	fmt.Printf("[Reset] foo:{%p}, Hello():{%v}\n", foo, foo.Hello())
}

func BenchmarkHotfix(b *testing.B) {
	foo := &model.Foo{
		String: "foo",
	}
	var (
		filePath = "./_patch_files/foo_fib.go" // 补丁脚本的路径
		evalText = "foo.GetPatch()"            // 补丁脚本内执行的函数名
	)

	// 加载补丁函数foo.GetPatch()
	_, err := hotfix.ApplyFunc(filePath, evalText, symbols.Symbols)
	if err != nil {
		fmt.Println(err)
		return
	}

	b.Run("github.com/cherry-game/cherry-hotfix", func(b *testing.B) {

		b.StartTimer()
		for i := 0; i < b.N; i++ {
			// foo.ExecFibonacciSum(10000)
		}

		b.StopTimer()
	})

}
