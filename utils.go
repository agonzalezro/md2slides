package main

import (
	"fmt"
	"log"
)

func ifErrFatal(err error) {
	if err != nil {
		fmt.Println("This shouldn't error, please open an issue attaching the following log:")
		fmt.Println("https://github.com/agonzalezro/md2slides/issues")
		fmt.Println()
		log.Fatal(err)
	}
}

func contains(keys []string, s string) bool {
	for _, k := range keys {
		if k == s {
			return true
		}
	}
	return false
}
