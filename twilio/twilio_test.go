package twilio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizePhoneNumber(t *testing.T) {
	tests := []struct {
		desc   string
		s      string
		exp    string
		expErr bool
	}{{
		desc:   `just numbers`,
		s:      `1234567890`,
		exp:    `+1234567890`,
		expErr: false,
	}, {
		desc:   `leading +`,
		s:      `+1234567890`,
		exp:    `+1234567890`,
		expErr: false,
	}, {
		desc:   `dashes`,
		s:      `123-456-7890`,
		exp:    `+1234567890`,
		expErr: false,
	}, {
		desc:   `parens and dashes`,
		s:      `(123)456-7890`,
		exp:    `+1234567890`,
		expErr: false,
	}, {
		desc:   `leading and trailing whitespace`,
		s:      "   1234567890 \n",
		exp:    `+1234567890`,
		expErr: false,
	}, {
		desc:   `empty string`,
		s:      ``,
		exp:    ``,
		expErr: true,
	}, {
		desc:   `only whitespace`,
		s:      "\t \n",
		exp:    ``,
		expErr: true,
	}}
	for _, tc := range tests {
		norm, err := normalizePhoneNumber(tc.s)
		assert.Equal(t, tc.exp, norm, tc.desc)
		if tc.expErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
