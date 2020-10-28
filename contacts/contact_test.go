package contacts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContactGetPhoneNumber(t *testing.T) {
	tests := []struct {
		desc     string
		contact  Contact
		expPhone string
		expErr   bool
	}{{
		desc: `both mobile and home number`,
		contact: Contact{
			MobilePhone: `1234567890`,
			HomePhone:   `0987654321`,
		},
		expPhone: `1234567890`,
		expErr:   false,
	}, {
		desc: `only home number`,
		contact: Contact{
			MobilePhone: ``,
			HomePhone:   `0987654321`,
		},
		expPhone: `0987654321`,
		expErr:   false,
	}, {
		desc: `only mobile number`,
		contact: Contact{
			MobilePhone: `1234567890`,
			HomePhone:   ``,
		},
		expPhone: `1234567890`,
		expErr:   false,
	}, {
		desc: `neither number`,
		contact: Contact{
			MobilePhone: ``,
			HomePhone:   ``,
		},
		expPhone: ``,
		expErr:   true,
	}}
	for _, tc := range tests {
		p, err := tc.contact.GetPhoneNumber()
		assert.Equal(t, tc.expPhone, p)
		if tc.expErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
