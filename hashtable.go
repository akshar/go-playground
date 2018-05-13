package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// bs, _ := ioutil.ReadAll(res.Body) // bs type []byte
	// str := string(bs)                 // string = slice of bytes and slice 0f bytes = rune, so convert to string
	// fmt.Println(str)
	// OR scanner takes a  Reader

	words := make(map[string]string)
	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
	}
	if err := sc.Err(); err != nil {
		fmt.Println(os.Stderr, "read err")
	}
	i := 0

	for k, _ := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++

	}
}
