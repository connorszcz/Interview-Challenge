package config

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	tests := []struct {
		desc      string
		json      string
		expConfig AppConfig
		expErr    bool
	}{{
		desc: `test normal config`,
		json: `{
	"twilio_account_sid": "test1",
	"twilio_auth_token": "test2",
	"twilio_from_number": "test3",
	"sender_first_name": "test4",
	"sender_last_name": "test5",
	"sender_mobile_number": "test6"
}`,
		expConfig: AppConfig{
			TwilioAccountSID:   `test1`,
			TwilioAuthToken:    `test2`,
			TwilioFromNumber:   `test3`,
			SenderFirstName:    `test4`,
			SenderLastName:     `test5`,
			SenderMobileNumber: `test6`,
		},
		expErr: false,
	}, {
		desc: `test incomplete json`,
		json: `{
	"twilio_account_sid": "test1",
	"sender_mobile_number": "test6"
}`,
		expConfig: AppConfig{},
		expErr:    true,
	}}
	for _, tc := range tests {
		r := strings.NewReader(tc.json)
		cfg, err := Read(r)
		assert.Equal(t, tc.expConfig, cfg, tc.desc)
		if tc.expErr {
			assert.Error(t, err, tc.desc)
		} else {
			assert.NoError(t, err, tc.desc)
		}
	}
}
func TestReadFile(t *testing.T) {
	tests := []struct {
		desc      string
		file      string
		expConfig AppConfig
		expErr    bool
	}{{
		desc: `test json file`,
		file: `testdata/testConfig.json`,
		expConfig: AppConfig{
			TwilioAccountSID:   `test1`,
			TwilioAuthToken:    `test2`,
			TwilioFromNumber:   `test3`,
			SenderFirstName:    `test4`,
			SenderLastName:     `test5`,
			SenderMobileNumber: `test6`,
		},
		expErr: false,
	}}
	for _, tc := range tests {
		cfg, err := ReadFile(tc.file)
		assert.Equal(t, tc.expConfig, cfg, tc.desc)
		if tc.expErr {
			assert.Error(t, err, tc.desc)
		} else {
			assert.NoError(t, err, tc.desc)
		}
	}
}
