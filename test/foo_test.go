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
	foo1 := &model.Foo{
		String: "foo1",
	}

	// 初始执行Hello(),并打印结果
	fmt.Printf("[Init]  foo1:{%p}, Hello():{%v}\n", foo1, foo1.Hello())

	// 模拟Hello()被调用
	for i := 0; i < 1000; i++ {
		go func(foo *model.Foo) {
			for {
				foo.Hello()
				time.Sleep(1 * time.Millisecond)
			}
		}(foo1)
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

	// 打印已被替换的foo1.Hello()
	fmt.Printf("[Patch] foo1:{%p}, Hello():{%v}\n", foo1, foo1.Hello())

	// 执行重置
	patches.Reset()

	// 打印函数
	fmt.Printf("[Reset] foo1:{%p}, Hello():{%v}\n", foo1, foo1.Hello())
}
