package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := "/hoge/fuga.txt"
	fmt.Println(filepath.Base(p))
	fmt.Println(filepath.Ext(p))
	fmt.Println(strings.Replace(filepath.Base(p), filepath.Ext(p), "", -1))
}
