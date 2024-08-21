package mock_client

import (
	"StudyZinx/business"
	"fmt"
)

type UI struct {
}

func NewUI() *UI {
	return &UI{}
}

func (ui *UI) ShowLaunchUI() {
	fmt.Println("ShowLaunchUI")
	fmt.Printf("%d. Register Account \n", business.Register)
	fmt.Printf("%d. Login \n", business.Login)
	fmt.Printf("%d. Exit \n", business.Exit)
}

func (ui *UI) ReadInput() string {
	var input string
	fmt.Println("Please input")
	fmt.Scanln(&input)
	return input
}

func (ui *UI) ShowRegister() {
	fmt.Println("input phone number")
}

func (ui *UI) ShowVerificationCode() {
	fmt.Println("input verification code")
}
