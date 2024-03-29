package main

import (
	"fmt"
	"log"

	"github.com/ios116/go-envdir/envdir"
	flag "github.com/spf13/pflag"
)

var dirname string
var prog string

func init() {
	flag.StringVarP(&dirname, "directory", "d", ".", "a directory with an environment values")
	flag.StringVarP(&prog, "programma", "p", "", "program to run")
}
func main() {
	flag.Parse()
	b, err := envdir.RunWithEnv(dirname, prog)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(b))
}
