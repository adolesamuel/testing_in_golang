package dob

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func calculateAge() {
	// Get the current date

	// Get the date of birth
	fmt.Print("Enter your date of birth (YYYY-MM-DD): ")
	reader := bufio.NewReader(os.Stdin)
	dateOfBirthInput, _ := reader.ReadString('\n')

	//convert dateOfBirthInput to time
	dateOfBirth, err := time.Parse("2006-01-02", dateOfBirthInput[:len(dateOfBirthInput)-1])

	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	// Calculate the difference between the two dates
	age := time.Since(dateOfBirth).Hours() / 24 / 365

	// Return the age

	fmt.Println("You are", age, "years old.")
}
