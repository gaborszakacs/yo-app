package chatbot

// Respond gives a response to a user input
func Respond(input string) string {
	if input == "yo" {
		return "yo"
	} else {
		return "???"
	}
}
