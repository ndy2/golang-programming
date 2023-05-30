package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type StringSorter []string

func (x StringSorter) Len() int { return len(x) }
func (x StringSorter) Less(i, j int) bool {
	if len(x[i]) != len(x[j]) {
		return len(x[i]) < len(x[j])
	} else {
		return x[i] < x[j]
	}
}
func (x StringSorter) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// https://www.acmicpc.net/problem/1181
func main() {
	var n int

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanln(r, &n)
	var s = make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(r, &s[i])
	}

	sort.Sort(StringSorter(s))

	for i := range s {
		if i > 0 && s[i-1] == s[i] {
			continue
		}
		fmt.Fprintln(w, s[i])
	}
}
