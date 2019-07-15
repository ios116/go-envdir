package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

var diranme string

func init() {
	flag.StringVarP(&dirname, "dir", "d", ".", "dir with envairment values")
}
func main() {
	flag.Parse()
	fmt.Println(dirname)
}
