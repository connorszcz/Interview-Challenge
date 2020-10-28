package csv

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		desc        string
		csv         string
		expContacts []Contact
		expErr      bool
	}{{
		desc: `test normal csv`,
		csv: `First Name,Last Name,Home Phone,Mobile Phone,Street Address,City,State,Zip,Date of Birth
Alex,Holmes,7077323644,5153164007,5928 E Third St.,Clive,IA,50123,1/1/1977
Brad,Reed,7077323644,5153164007,5929 E Third St.,Clive,IA,50124,2/2/1977`,
		expContacts: []Contact{{
			FirstName:     `Alex`,
			LastName:      `Holmes`,
			HomePhone:     `7077323644`,
			MobilePhone:   `5153164007`,
			StreetAddress: `5928 E Third St.`,
			City:          `Clive`,
			State:         `IA`,
			Zip:           `50123`,
			BirthMonth:    time.January,
		}, {
			FirstName:     `Brad`,
			LastName:      `Reed`,
			HomePhone:     `7077323644`,
			MobilePhone:   `5153164007`,
			StreetAddress: `5929 E Third St.`,
			City:          `Clive`,
			State:         `IA`,
			Zip:           `50124`,
			BirthMonth:    time.February,
		}},
		expErr: false,
	}, {
		desc:        `test empty csv`,
		csv:         ``,
		expContacts: nil,
		expErr:      true,
	}, {
		desc:        `test only headers`,
		csv:         `First Name,Last Name,Home Phone,Mobile Phone,Street Address,City,State,Zip,Date of Birth`,
		expContacts: nil,
		expErr:      true,
	}, {
		desc: `test invalid birth date format`,
		csv: `First Name,Last Name,Home Phone,Mobile Phone,Street Address,City,State,Zip,Date of Birth
Alex,Holmes,7077323644,5153164007,5928 E Third St.,Clive,IA,50123,1/1`,
		expContacts: nil,
		expErr:      true,
	}, {
		desc: `test nonnumeric birth month`,
		csv: `First Name,Last Name,Home Phone,Mobile Phone,Street Address,City,State,Zip,Date of Birth
Alex,Holmes,7077323644,5153164007,5928 E Third St.,Clive,IA,50123,a/1/2020`,
		expContacts: nil,
		expErr:      true,
	}, {
		desc: `test out of range birth month`,
		csv: `First Name,Last Name,Home Phone,Mobile Phone,Street Address,City,State,Zip,Date of Birth
Alex,Holmes,7077323644,5153164007,5928 E Third St.,Clive,IA,50123,13/1/2020`,
		expContacts: nil,
		expErr:      true,
	}}
	for _, tc := range tests {
		r := strings.NewReader(tc.csv)
		contacts, err := Parse(r)
		assert.Equal(t, tc.expContacts, contacts, tc.desc)
		if tc.expErr {
			assert.Error(t, err, tc.desc)
		} else {
			assert.NoError(t, err, tc.desc)
		}
	}
}
func TestParseFile(t *testing.T) {
	tests := []struct {
		desc        string
		file        string
		expContacts []Contact
		expErr      bool
	}{{
		desc: `test part of the given address book`,
		file: `testdata/testAddressBook.csv`,
		expContacts: []Contact{{
			FirstName:     `Alex`,
			LastName:      `Holmes`,
			HomePhone:     `7077323644`,
			MobilePhone:   `5153164007`,
			StreetAddress: `5928 E Third St.`,
			City:          `Clive`,
			State:         `IA`,
			Zip:           `50123`,
			BirthMonth:    time.January,
		}, {
			FirstName:     `Brad`,
			LastName:      `Reed`,
			HomePhone:     `7077323644`,
			MobilePhone:   `5153164007`,
			StreetAddress: `5929 E Third St.`,
			City:          `Clive`,
			State:         `IA`,
			Zip:           `50124`,
			BirthMonth:    time.February,
		}},
		expErr: false,
	}}
	for _, tc := range tests {
		contacts, err := ParseFile(tc.file)
		assert.Equal(t, tc.expContacts, contacts, tc.desc)
		if tc.expErr {
			assert.Error(t, err, tc.desc)
		} else {
			assert.NoError(t, err, tc.desc)
		}
	}
}
