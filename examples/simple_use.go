package main

import (
	"fmt"
	"github.com/akazwz/go-i18n/i18n"
)

// Test in English is
// Test in Chinese is 测试
func main() {
	i := i18n.I18n{}
	i.SetLang("en")
	i.SetTransFilePath("./config/i18n")
	i.SetTransFileType("yaml")

	enStr := i.Trans("test.name").ToStr()

	i.SetLang("zh")
	zhStr := i.Trans("test.name").ToStr()

	fmt.Printf("Test in English is %s\n", enStr)
	fmt.Printf("Test in Chinese is %s\n", zhStr)
}
