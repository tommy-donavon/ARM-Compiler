package main

import (
	"compiler/compiler"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("invalid number of arguments")
	}
	flag := args[0]
	input := args[1]
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	a, err := compiler.BuildCompiler(flag, string(b))
	if err != nil {
		log.Fatal(err.Error())
	}
	n, err := os.Create("input.a")
	if err != nil {
		log.Fatal(err)
	}
	_, err = n.WriteString(a.ToString())
	if err != nil {
		log.Fatal(err)
	}
}
