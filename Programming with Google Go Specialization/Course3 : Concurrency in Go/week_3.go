package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"

)


//=================================Main
func main() {

	//Get the user input
	fmt.Print("Enter a series of integers separated by spaces: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	integers := strings.Split(input, " ")

	var nums []int
	for _, str := range integers {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}


	//Partition the list into 4
	Array_len := len(nums)
	q := Array_len/4
	list_1 := nums[0: q]
	list_2 := nums[q: q*2]
	list_3 := nums[q*2: q*3]
	list_4 := nums[q*3:]
	fmt.Println(
	"List 1: ", list_1,"\n",
	"List 2 ", list_2, "\n",
	"List 3: ",list_3,"\n",
	"List 4: ", list_4)


	//Sort small lists
	go sort.Ints(list_1)
	go sort.Ints(list_2)
	go sort.Ints(list_3)
	go sort.Ints(list_4)

	//Merge and sort the listed
	merged_list := make([]int, 0, Array_len)
	merged_list = append(merged_list, list_1...)
    merged_list = append(merged_list, list_2...)
    merged_list = append(merged_list, list_3...)
    merged_list = append(merged_list, list_4...)

	sort.Ints(merged_list)

	fmt.Println("\nMerged array:", merged_list)
	
}