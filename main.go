package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/connorszcz/Interview-Challenge/config"
	"github.com/connorszcz/Interview-Challenge/contacts"
	"github.com/connorszcz/Interview-Challenge/twilio"
)

// Provide a CLI to specify the address book and config files. Could just as easily be a .env file or environment vars
var (
	addressBookFile = kingpin.Flag(`address-book`, `path to address book CSV file`).Default(`addressBook.csv`).String()
	configFile      = kingpin.Flag(`config-file`, `path to config file`).Default(`config.json`).String()
)

func main() {
	kingpin.Parse()
	cfg, err := config.ReadFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	contacts, err := contacts.ParseFile(*addressBookFile)
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

		msg := createMessage(c.FirstName, cfg.SenderFirstName, cfg.SenderLastName, cfg.SenderMobileNumber)
		// `c.MobilePhone` is guaranteed not to be empty by `contacts.Parse`
		err = twilioClient.SendSMS(cfg.TwilioFromNumber, msg, c.MobilePhone)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createMessage(recipientFirst, senderFirst, senderLast, senderPhoneNumber string) string {
	return fmt.Sprintf(`Happy Birthday %s from %s %s! Call me at %s to plan a lunch sometime.`, recipientFirst, senderFirst, senderLast, senderPhoneNumber)
}
