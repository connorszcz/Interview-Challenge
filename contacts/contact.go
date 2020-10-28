package contacts

import "time"

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
