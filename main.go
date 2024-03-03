package main

import (
	"fmt"
	"github.com/yuin/gopher-lua"
)

func main() {
	Lua := lua.NewState()

	// 注册标准库
	Lua.OpenLibs(L)

	// 运行一个简单的Lua脚本
	if err := Lua.DoString(`print("Hello, Lua from Golang!")`); err != nil {
		fmt.Println("Error running Lua script:", err)
		return
	}
}