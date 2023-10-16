package internal

func Authentication(input string) {
	switch input {
	case "Create":
		CreateAccount()
	default:
		CheckPassword(ValidateAndPad(input))
	}

}
