package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/connorszcz/Interview-Challenge/csv"

	"github.com/connorszcz/Interview-Challenge/config"
	"github.com/connorszcz/Interview-Challenge/twilio"
)

func main() {
	cfg, err := config.ReadFile(`config.json`)
	if err != nil {
		log.Fatal(err)
	}

	contacts, err := csv.ParseFile(`addressBook.csv`)
	if err != nil {
		log.Fatal(err)
	}

	thisMonth := time.Now().Month()

	for _, c := range contacts {
		if c.BirthMonth != thisMonth {
			continue
		}
		msg := createMessage(c.FirstName, cfg.SenderFirstName, cfg.SenderLastName, cfg.SenderMobileNumber)
		twilioClient, err := twilio.NewClient(cfg.TwilioAccountSID, cfg.TwilioAuthToken)
		if err != nil {
			log.Fatal(err)
		}
		sendTo := c.MobilePhone
		if sendTo == `` {
			if c.HomePhone == `` {
				log.Fatal(errors.New(`No valid number for contact: ` + c.FirstName + ` ` + c.LastName))
			}
			// TODO: verify that we want to use home phone as a fallback
			sendTo = c.HomePhone
		}
		err = twilioClient.SendSMS(cfg.TwilioFromNumber, msg, sendTo)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createMessage(recipientFirst, senderFirst, senderLast, senderPhoneNumber string) string {
	return fmt.Sprintf(`Happy Birthday %s from %s %s! Call me at %s to plan a lunch sometime.`, recipientFirst, senderFirst, senderLast, senderPhoneNumber)
}
