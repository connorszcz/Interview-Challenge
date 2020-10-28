package main

import (
	"fmt"
	"log"
	"time"

	"github.com/connorszcz/Interview-Challenge/contacts"

	"github.com/connorszcz/Interview-Challenge/config"
	"github.com/connorszcz/Interview-Challenge/twilio"
)

func main() {
	cfg, err := config.ReadFile(`config.json`)
	if err != nil {
		log.Fatal(err)
	}

	contacts, err := contacts.ParseFile(`addressBook.csv`)
	if err != nil {
		log.Fatal(err)
	}

	thisMonth := time.Now().Month()

	twilioClient, err := twilio.NewClient(cfg.TwilioAccountSID, cfg.TwilioAuthToken)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range contacts {
		if c.BirthMonth != thisMonth {
			continue
		}

		sendTo, err := c.GetValidPhoneNumber()
		if err != nil {
			log.Fatal(err)
		}

		msg := createMessage(c.FirstName, cfg.SenderFirstName, cfg.SenderLastName, cfg.SenderMobileNumber)
		err = twilioClient.SendSMS(cfg.TwilioFromNumber, msg, sendTo)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createMessage(recipientFirst, senderFirst, senderLast, senderPhoneNumber string) string {
	return fmt.Sprintf(`Happy Birthday %s from %s %s! Call me at %s to plan a lunch sometime.`, recipientFirst, senderFirst, senderLast, senderPhoneNumber)
}
