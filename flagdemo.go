package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	textPointer := flag.String("text", "", "Text to parse. (Req)")
	metricPointer := flag.String("metric", "chars", "Metrics (chars|words|lines) (Req)")
	uniquePointer := flag.Bool("unique", false, "Get unique value of words")

	flag.Parse()
	// flag.PrintDefaults()

	if *textPointer == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Printf("text: %s, metric: %s, uniquepointer: %t\n", *textPointer, *metricPointer, *uniquePointer)
	fmt.Println(flag.Arg(1))
}
