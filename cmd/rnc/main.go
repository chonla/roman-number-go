package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"

	"github.com/pedroalbanese/roman-number-go"
)

func main() {
	flag.Parse()
	if isInt(os.Args[1]) {
		r := roman.NewRoman()
		s, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Println(err)
		}
		n := r.ToRoman(s)
		fmt.Println(n)
	} else {
		r := roman.NewRoman()
		n := r.ToNumber(os.Args[1])
		fmt.Println(n)
	}
}

func isInt(s string) bool {
    for _, c := range s {
        if !unicode.IsDigit(c) {
            return false
        }
    }
    return true
}