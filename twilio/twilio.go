package twilio

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

const baseURL = `api.twilio.com/2010-04-01/`

type TwilioClient struct {
	sid       string
	authToken string
}

func NewClient(sid, authToken string) (TwilioClient, error) {
	if sid == `` || authToken == `` {
		return TwilioClient{}, errors.New(`SID and AuthToken must not be emprty!`)
	}
	return TwilioClient{
		sid:       sid,
		authToken: authToken,
	}, nil
}

func (tc *TwilioClient) SendSMS(from, body, to string) error {
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
	if res.StatusCode != http.StatusOK {
		return errors.New(`API returned bad error code!`)
	}
	return nil
}
