package main

import (
	"fmt"
	"path/filepath"
	"regexp"
)

func main() {
	a := "/a/b.md"
	e := filepath.Base(a)
	fmt.Println(a, e)
	v := regexp.MustCompile(".md$")
	vv := v.ReplaceAll([]byte(filepath.Base(a)), []byte(""))
	fmt.Println(string(vv))
}
