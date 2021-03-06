package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps Review - Histogram")
	data := map[int]int{2: 3, 3: 6, 4: 9, 5: 13, 6: 15, 7: 18,
		8: 16, 9: 14, 10: 10, 12: 2}

	hist(data, 1)
	hhist(data)
}

/*
 Histogram
10    *
9     *
8     *  *
7     *  *     *
6     *  *     *              *
5     *  *  *  *        *     *
4     *  *  *  *        *     *
3     *  *  *  *  *     *     *
2     *  *  *  *  *     *     *  *
1     *  *  *  *  *     *  *  *  *
0     *  *  *  *  *  *  *  *  *  *
      1  2  3  4  5  6  7  8  9  10
*/
func hist(d map[int]int, sf int) {
	fmt.Println("\nHistogram")
	if 0 == len(d) {
		return
	}

	max := 0
	keys := []int{}
	for k, v := range d {
		keys = append(keys, k)
		if max < v {
			max = v
		}
	}

	sort.Ints(keys)

	for i := max / sf; i > -1; i-- {
		// fmt.Printf("%2v  ", i)  // simplest solution
		if i > 9 {
			fmt.Print(i, " ")
		} else {
			fmt.Print(" ", i, " ")
		}

		for _, k := range keys {
			if i <= d[k]/sf {
				fmt.Print("*  ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}

	// print column number
	fmt.Print("   ")
	for _, k := range keys {
		if k < 10 {
			fmt.Print(k, "  ")
		} else {
			fmt.Print(k, " ")
		}
	}
	fmt.Println()
	fmt.Println("------------------------------------------------")
}

/*
 Histogram (horizontal)
 ************************
 *****************
 ********************
 *
 ************
 ******
 *****************
 ***********************
*/
func hhist(d map[int]int) {
	fmt.Println("\nHistogram (horizontal)")
	keys := []int{}
	for k := range d {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		// fmt.Printf("%2v ", k) // simplest solution or
		if k < 10 {
			fmt.Print(" ", k, " ")
		} else {
			fmt.Print("", k, " ")
		}

		for j := 0; j < d[k]; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println("------------------------------------------------")
}
