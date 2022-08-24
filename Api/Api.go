package Api

import (
	"fmt"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/twilio/twilio-go"
)

// SendMessageToPhone is function to send message
// From Phone number
// To Phone number
// that you choose
func SendMessageToPhone(to string, client *twilio.RestClient) {

	from := "+1234****9"
	params := &openapi.CreateMessageParams{}
	params.SetFrom(from)
	params.SetTo(to)
	params.SetBody("Hello Rami")

	resp, err := client.Api.CreateMessage(params)

	if err != nil {
		fmt.Printf("Unable to send message to this phone number %s\n", from)
		panic(err)
	} else {
		fmt.Printf("Sent verification '%s'\n", *resp.Sid)
	}
}
