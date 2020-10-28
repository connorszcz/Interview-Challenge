package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

const (
	firstNameIdx int = iota
	lastNameIdx
	homePhoneIdx
	mobilePhoneIdx
	streetAddressIdx
	cityIdx
	stateIdx
	zipIdx
	birthDateIdx
)

func ParseFile(filename string) ([]Contact, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return Parse(f)
}

func Parse(r io.Reader) ([]Contact, error) {
	csvReader := csv.NewReader(r)
	// Don't parse the header
	_, err := csvReader.Read()
	if err != nil {
		if err == io.EOF {
			return nil, errors.New(`CSV file is empty!`)
		}
		return nil, err
	}

	contacts := make([]Contact, 0)
	for {

		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		birthMonth, err := parseBirthMonth(record[birthDateIdx])
		if err != nil {
			return nil, fmt.Errorf(`Error parsing birth month in line %d`, len(contacts)+1)
		}
		contacts = append(contacts, Contact{
			FirstName:     record[firstNameIdx],
			LastName:      record[lastNameIdx],
			HomePhone:     record[homePhoneIdx],
			MobilePhone:   record[mobilePhoneIdx],
			StreetAddress: record[streetAddressIdx],
			City:          record[cityIdx],
			State:         record[stateIdx],
			Zip:           record[zipIdx],
			BirthMonth:    birthMonth,
		})
	}
	if len(contacts) == 0 {
		return nil, errors.New(`No contacts in the file!`)
	}
	return contacts, nil
}

func parseBirthMonth(s string) (time.Month, error) {
	parts := strings.Split(s, `/`)
	if len(parts) != 3 {
		return 0, errors.New(`Invalid date format!`)
	}
	val, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}
	if val < 1 || val > 12 {
		return 0, errors.New(`Month must be between 1 and 12!`)
	}
	return time.Month(val), nil
}
