package chatbot

// Respond gives a response to a user input, yo
func Respond(input string) string {
	if input == "yo-yo" {
		return "yo-yo"
	}

	if input == "yo" {
		return "yo"
	} else {
		return "???"
	}
}
