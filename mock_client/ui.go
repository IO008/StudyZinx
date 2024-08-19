package mock_client

import "fmt"

const (
	register         = "0"
	login            = "1"
	exit             = "2"
	verificationCode = "3"
)

type UI struct {
}

func NewUI() *UI {
	return &UI{}
}

func (ui *UI) ShowLaunchUI() {
	fmt.Println("ShowLaunchUI")
	fmt.Printf("%s. Register Account \n", register)
	fmt.Printf("%s. Login \n", login)
	fmt.Printf("%s. Exit \n", exit)
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
