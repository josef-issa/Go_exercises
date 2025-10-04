//The program calculates the grade of 5 students

package main

import "fmt"

func main() {

	student_name := []string{"Alix", "Benjamen", "David", "Noah", "Tom"}
	degrees := []int{60, 70, 90, 50, 40}

	for i := 0; i < len(student_name)&len(degrees); i++ {
		fmt.Println(student_name[i], degrees[i])
	}

	sum := 0
	for _, degree := range degrees {
		sum += degree

	}
	average := sum / len(degrees)
	fmt.Println("The grade level 5 students is:", average)

}
