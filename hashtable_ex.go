package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func HashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
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
		n := HashBucket(scanner.Text(), 12)
		buckets[n]++
	}

	fmt.Println(buckets[0:13]) // A,Z  a,z
}
