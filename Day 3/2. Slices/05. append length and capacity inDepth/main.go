package main

import "fmt"

func main() {
	var nums []int
	fmt.Printf("%#v\n", nums)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))
	// Output: Length: 0, Capacity: 0

	nums = append(nums, 1, 2)
	fmt.Printf("%#v\n", nums)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))
	// Output: Length: 2, Capacity: 2

	nums = append(nums, 3)
	fmt.Printf("%#v\n", nums)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))
	// Output: Length: 3, Capacity: 4

	nums = append(nums, 4)
	fmt.Printf("%#v\n", nums)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))
	// Output: Length: 4, Capacity: 4

	nums = append(nums, 5)
	fmt.Printf("%#v\n", nums)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))
	// Output: Length: 5, Capacity: 8

	nums = append(nums, 100)
	fmt.Printf("%#v\n", nums)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))
	// Output: Length: 6, Capacity: 8

	nums = append(nums, 5)
	fmt.Printf("%#v\n", nums)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))
	// Output: Length: 7, Capacity: 8

	nums = append(nums, 5)
	fmt.Printf("%#v\n", nums)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))
	// Output: Length: 8, Capacity: 8

	nums = append(nums, 5)
	fmt.Printf("%#v\n", nums)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))
	// Output: Length: 9, Capacity: 16

	letters := []string{"A", "B", "C", "D", "E", "F"}
	letters = append(letters[0:1], "X", "Y")

	fmt.Println(letters, len(letters), cap(letters))

	// Error: Index Out of Range
	// fmt.Println(letters[4])

	fmt.Println(letters[3:6])
}
