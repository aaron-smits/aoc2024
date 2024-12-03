package main

import (
	"bufio"
	"strconv"
	"strings"
	"fmt"
	"math"
	"os"
	"sort"
)



func main() {
	var total float64
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var iter int
	var list1 []int
	var list2 []int

	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(txt)
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		list1 = append(list1, num1)
		list2 = append(list2, num2)
		fmt.Println(iter, ": ", total)
	}
	total = findDistance(list1, list2)
	step2 := findCount(list1, list2)
 	if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
    fmt.Printf("%.0f\n", total)
    fmt.Printf("%v\n", step2)

}


func findDistance(list1, list2 []int)(total float64) {
	sorted1 := sortIntSlice(list1)
	sorted2 := sortIntSlice(list2)
	fmt.Println(sorted1)
	fmt.Println(sorted2)
	for i := range sorted1 {
		distance := sorted1[i] - sorted2[i]
		fmt.Println("Distance is ", sorted1[i], "-", sorted2[i], "=", distance)
		total += math.Abs(float64(distance))
		fmt.Println("Added ", math.Abs(float64(distance)), " to the total: ", total)
	}
	return total
}
// adding up each number in the left list after multiplying it by the number of times that number appears in the right list.
func findCount(list1, list2 []int)(total int) {
	sorted1 := sortIntSlice(list1)
	sorted2 := sortIntSlice(list2)
	var count int
	for i := range sorted1 {
		count = 0
		for j := range sorted2 {
			if sorted1[i] == sorted2[j] {
				count++
				fmt.Println(sorted1[i], "==", sorted2[j], ", count updated: ", count)
			}
		}
		total = total + (count * sorted1[i])
		fmt.Println("Total updated: ", total, "=", count, "*", sorted1[i])
	}
	return total
}

func sortIntSlice(intslice []int) []int {
	sort.Slice(intslice, func(i, j int) bool {
		return intslice[i] < intslice[j]
	})
	return intslice
}
