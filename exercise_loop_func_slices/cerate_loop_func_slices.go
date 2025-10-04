//Loop and functions and arrays and slices
//Function: Print numbers from 1 to 10

package main

import "fmt"

// cerate function
func myfunction(x int, y int) int {
	return x + y
}

func main() {
	//cerate an arryas with loop
	name := []string{"Josef", "Lomay", "Amjad"}
	age := []int{34, 27, 34}
	fmt.Println(name)
	fmt.Println(age)

	for j := 0; j < len(name)&len(age); j++ {
		//for j := 0; j < len(age); j++ {
		fmt.Println(name[j], age[j])
	}
	//print func as cerate befor
	fmt.Println("sum is", (myfunction(3, 5)))

	//cerate a new loop to print 1 to 10
	for i := 0; i <= 10; i++ {

		fmt.Println(i)
	}
}
