package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser User

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@test.com", "123")
	admin.OutputUserDetails()

	// ... do something awesome with that gathered data!

	// fmt.Println(userFirstName, userLastName, userBirthdate)

	// outputUserDetails(appUser)

	// outputUserDetails(&appUser)

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}

// func outputUserDetails(u User) {
// 	fmt.Println(u.firstName, u.lastName, u.birthDate)
// }

// func outputUserDetails(u *User) {
// 	fmt.Println(u.firstName, u.lastName, u.birthDate)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
