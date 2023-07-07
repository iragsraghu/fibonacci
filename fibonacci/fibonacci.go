package fibonacci

import (
	"log"
	"strconv"
	"strings"
)

func GetFibonacci(n int) string {
	log.Printf("Fibonacci series of %v\n", n)
	var result []string
	firstTerm, secondTerm, nextTerm, counterTerm := 0, 1, 0, 1
	result = append(result, strconv.Itoa(firstTerm))
	for counterTerm <= n+1 {
		if counterTerm <= 1 {
			nextTerm = counterTerm
		} else {
			result = append(result, strconv.Itoa(nextTerm))
			nextTerm = firstTerm + secondTerm
			firstTerm = secondTerm
			secondTerm = nextTerm
		}
		counterTerm++
	}
	return strings.Join(result, ",")
}
