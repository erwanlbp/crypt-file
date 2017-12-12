package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var crypt bool
var decrypt bool
var file string
var pass string

func init() {
	flag.BoolVar(&crypt, "c", false, "To crypt the file")
	flag.BoolVar(&decrypt, "d", false, "To decrypt the file")
	flag.StringVar(&file, "f", "", "The file to crypt/decrypt")
	flag.StringVar(&pass, "p", "", "The password to crypt/decrypt")
	flag.Parse()
}

func checkArgs() error {
	if len(file) == 0 {
		return errors.New("No file provided")
	}
	if _, err := os.Stat(file); err != nil && os.IsNotExist(err) {
		return err
	}
	if !crypt && !decrypt {
		return errors.New("No action provided")
	}
	if crypt && decrypt {
		return errors.New("Only one action please, -c or -d")
	}
	return nil
}

func main() {
	if err := checkArgs(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}



}
