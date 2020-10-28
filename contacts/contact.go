package contacts

import (
	"fmt"
	"time"
)

type Contact struct {
	FirstName     string
	LastName      string
	HomePhone     string
	MobilePhone   string
	StreetAddress string
	City          string
	State         string
	Zip           string
	BirthMonth    time.Month
}

func (c Contact) GetPhoneNumber() (string, error) {
	if c.MobilePhone != `` {
		return c.MobilePhone, nil
	}
	if c.HomePhone == `` {
		return ``, fmt.Errorf(`No valid phone number for contact: %s %s`, c.FirstName, c.LastName)
	}
	// TODO: check the assumption that we want to fall back to HomePhone!
	return c.HomePhone, nil
}
