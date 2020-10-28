package config

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

type AppConfig struct {
	TwilioAccountSID   string `json:"twilio_account_sid,omitempty"`
	TwilioAuthToken    string `json:"twilio_auth_token,omitempty"`
	TwilioFromNumber   string `json:"twilio_from_number,omitempty"`
	SenderFirstName    string `json:"sender_first_name,omitempty"`
	SenderLastName     string `json:"sender_last_name,omitempty"`
	SenderMobileNumber string `json:"sender_mobile_number,omitempty"`
}

func ReadFile(filename string) (AppConfig, error) {
	f, err := os.Open(filename)
	if err != nil {
		return AppConfig{}, nil
	}
	return Read(f)
}

func Read(r io.Reader) (AppConfig, error) {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return AppConfig{}, nil
	}
	var cfg AppConfig
	err = json.Unmarshal(bs, &cfg)
	if err != nil {
		return AppConfig{}, nil
	}
	if cfg.SenderFirstName == `` || cfg.SenderLastName == `` || cfg.SenderMobileNumber == `` || cfg.TwilioAccountSID == `` || cfg.TwilioAuthToken == `` || cfg.TwilioFromNumber == `` {
		return AppConfig{}, errors.New(`All fields in the config are required!`)
	}
	return cfg, nil
}
