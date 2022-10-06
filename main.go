package main

import (
	"github.com/ramialkaro/phone-verification-with-twilio/Api"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
)

func main() {

	err := godotenv.Load("development.env")

	if err != nil {
		log.Fatalf("Some error happened when load ENV file. Err: %s", err)
	}

	// Environment variables
	TwilioAccountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	TwilioAuthToken := os.Getenv("TWILIO_AUTH_TOKEN")

	// Phone number that you want to send message / verification code
	to := "+358401972828"

	// Setup credentials for the Twilio
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TwilioAccountSid,
		Password: TwilioAuthToken,
	})
	Api.SendMessageToPhone(to, client)
}
