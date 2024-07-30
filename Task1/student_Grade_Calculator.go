package main


import (
	"fmt"
)

func calculateGrade(subject_grade_pair map[string]float32 , totalsubject float32) float32{

	var total float32

	// Calculate the totalsum
	for _, value := range subject_grade_pair { 
		total += value
	}

	// Calculate the average grade
	average := total/totalsubject

	return average
	

} 

func displayDetails(subject_grade_pair *map[string]float32, average float32) {
	border := "----------------------------"
	centered := "        "
	fmt.Println(centered + border)
	fmt.Println(centered + "Subject          Grades\n" , )

	for subject, grade := range *subject_grade_pair {
		fmt.Printf(centered+"%-15s: %.2f\n", subject, grade)
	}
	fmt.Println(centered + border)
	fmt.Printf(centered+"Average Grade:   %.2f\n", average)
	fmt.Println(centered + border)
}



func main(){

	
	var name string 
	var number_of_subject int
	subject_grade_pair := make(map[string]float32)
	
	
	fmt.Print("Enter your name : ") 
	fmt.Scanln(&name)
	fmt.Print("Enter the number of subjects : ")
	fmt.Scanln(&number_of_subject)

	for i := 0; i < number_of_subject; i++ { 
		var subject_name string
		var subject_grade float32
		fmt.Print("Enter the subject name : ")
		fmt.Scanln(&subject_name)
		fmt.Print("Enter the grade of the subject : ")
		fmt.Scanln(&subject_grade)

		// handling invalid grades
		if subject_grade < 0 || subject_grade > 100 { 

			// loop until the user enters a valid grade
			for {
				fmt.Print("Please enter a valid grade : ")
				fmt.Scanln(&subject_grade)

				// check if the grade is valid
				if subject_grade >= 0 && subject_grade <= 100 {
					break
				}
			}
			
		}
		// add the subject and grade to the map

		subject_grade_pair[subject_name] = subject_grade
	}


	// call the function to calculate the average grade
	avergae := calculateGrade(subject_grade_pair , float32(number_of_subject))

	// call by reference to display the details for efficiency
	displayDetails(&subject_grade_pair, avergae)
	




}