package main

import "flag"

// name ...
func main() {
	countCommand := flag.NewFlagSet("count", flag.ExitOnError)
	listCommand := flag.NewFlagSet("list", flag.ExitOnError)

	countTextPtr := countCommand.String("text", "", "Text to parse(required")
	countMetricPtr := countCommand.String("metric", "chars", "Metric  {chars|text|words}")
	countSubstringPtr := countCommand.String("substring", "", "The Substring")
	countUniquePtr := countCommand.Bool("unique", false, "Measure unique values")

	// list subcommand flag pointers

}
