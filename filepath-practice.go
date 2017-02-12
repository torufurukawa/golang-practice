package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	p, _ := os.Getwd()
	fmt.Println("current dir:", p)

	abs, _ := filepath.Abs(p)
	fmt.Println("abs:", abs)

	fmt.Println("base:", filepath.Base(p))
	fmt.Println("dir:", filepath.Dir(p))
	fmt.Println("vol:", filepath.VolumeName(p))
	fmt.Println("to slash:", filepath.ToSlash(p))
}
