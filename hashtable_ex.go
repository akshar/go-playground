package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func HashBucket(word string) int {
	return int(word[0])
}

func main() {
	res, err := http.Get("https://www.gutenberg.org/files/2701/old/moby10b.txt")

	if err != nil {
		log.Fatalln(err)
	}
	//scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)

	//create slice to hold counts
	buckets := make([]int, 200)
	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}

	fmt.Println(buckets[65:123]) // A,Z  a,z
}
