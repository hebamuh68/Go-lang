/*
*function called BubbleSort() which
*takes a slice of integers as an argument and returns nothing.
*The BubbleSort() function should modify the slice so that the
*elements are in sorted order.

===========================================================================================

*You should write a Swap() function which performs this operation.
*Your Swap() function should take two arguments, a slice of integers and an index
*value i which indicates a position in the slice.
*The Swap() function should return nothing, but it should swap
*the contents of the slice in position i with the contents in position i+1.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func Swap(Nums []int, ndx int){
	lenNums := len(Nums)

	for j := 0; j < lenNums - ndx - 1; j++ {
			
		if (Nums[j] > Nums[j+1]) {
			tmp := Nums[j]
			Nums[j] = Nums[j+1]
			Nums[j+1] = tmp
		}
	}
}

func BubbleSort(Nums []int){
	lenNums := len(Nums)

	for i := 0; i < lenNums; i++ {
		Swap(Nums, i)
	}
}

func main() {
    var Nums []int

	//Take the input
	fmt.Println("Please input numbers(separate with space):")
    br := bufio.NewReader(os.Stdin)
    a, _, _ := br.ReadLine()
    ns := strings.Split(string(a), " ")

    for _, s := range(ns) {
      n, _ := strconv.Atoi(s)
      Nums = append(Nums, n)
    }

	//Main RUN
	BubbleSort(Nums)
	fmt.Println(Nums)
}