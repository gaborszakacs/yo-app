package chatbot_test

import (
	"testing"

	"github.com/gaborszakacs/yo-app/chatbot"
)

func TestRespond(t *testing.T) {
	testCases := []struct {
		input            string
		expectedResponse string
	}{
		{
			input:            "yo",
			expectedResponse: "yo",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			actualResponse := chatbot.Respond(tC.input)
			if actualResponse != tC.expectedResponse {
				t.Errorf("Expected the response to '%s' to be '%s' but instead got '%s'!", tC.input, tC.expectedResponse, actualResponse)
			}
		})
	}
}
