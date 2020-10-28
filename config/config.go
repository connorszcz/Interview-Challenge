package config

type AppConfig struct {
	TwilioAccountSID   string `json:"twilio_account_sid,omitempty"`
	TwilioAuthToken    string `json:"twilio_auth_token,omitempty"`
	TwilioFromNumber   string `json:"twilio_from_number,omitempty"`
	SenderFirstName    string `json:"sender_first_name,omitempty"`
	SenderLastName     string `json:"sender_last_name,omitempty"`
	SenderMobileNumber string `json:"sender_mobile_number,omitempty"`
}
