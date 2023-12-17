package main

import (
	"flag"
	"fmt"
)

func main() {
	flagPattern := flag.String("p", "", "filter by pattern")
	flagAll := flag.Bool("a", false, "show all files")
	flagNumberRecords := flag.Int("n", 0, "number of records")
	hasOrderByTime := flag.Bool("t", false, "order by time, oldest first")
	hasOrderBySize := flag.Bool("s", false, "order by size, smallest first")
	hasOrderReverse := flag.Bool("r", false, "reverse order")

	flag.Parse()

	fmt.Println("pattern:", *flagPattern)
	fmt.Println("all:", *flagAll)
	fmt.Println("number records:", *flagNumberRecords)
	fmt.Println("order by time:", *hasOrderByTime)
	fmt.Println("order by size:", *hasOrderBySize)
	fmt.Println("reverse order:", *hasOrderReverse)
}
