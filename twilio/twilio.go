package twilio

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const baseURL = `api.twilio.com/2010-04-01/`

type TwilioClient struct {
	sid       string
	authToken string
}

func NewClient(sid, authToken string) (*TwilioClient, error) {
	if sid == `` || authToken == `` {
		return nil, errors.New(`SID and AuthToken must not be emprty!`)
	}
	return &TwilioClient{
		sid:       sid,
		authToken: authToken,
	}, nil
}

// SendSMS interacts with the twilio API to send a message with body `body` from mobile number `from` to mobile number `to`
func (tc *TwilioClient) SendSMS(from, body, to string) error {
	from, err := normalizePhoneNumber(from)
	if err != nil {
		return err
	}
	to, err = normalizePhoneNumber(to)
	if err != nil {
		return err
	}

	u := url.URL{
		Scheme: `https`,
		Path:   baseURL + `Accounts/` + tc.sid + `/Messages.json`,
	}

	q := url.Values{}
	q.Set(`To`, to)
	q.Set(`From`, from)
	q.Set(`Body`, body)
	queryReader := strings.NewReader(q.Encode())

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, u.String(), queryReader)
	if err != nil {
		return err
	}
	req.SetBasicAuth(tc.sid, tc.authToken)
	req.Header.Add(`Content-Type`, `application/x-www-form-urlencoded`)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	// accept any status code in the 200s
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		fmt.Println(res.Status)
		return errors.New(`API returned bad error code!`)
	}
	return nil
}

func normalizePhoneNumber(s string) (string, error) {
	// I won't do an exhaustive check on the format of the incoming phone number for this exercise, but I want to support _some_ basic normalization to get the Twilio's required format
	// Normalized format: +1234567890
	s = strings.TrimSpace(s)
	if s == `` {
		return ``, errors.New(`Phone number is empty!`)
	}
	s = strings.ReplaceAll(s, `-`, ``)
	s = strings.ReplaceAll(s, `(`, ``)
	s = strings.ReplaceAll(s, `)`, ``)
	s = strings.ReplaceAll(s, `+`, ``)
	// add a +1 for the country code
	return `+1` + s, nil
}
