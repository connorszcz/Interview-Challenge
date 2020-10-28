# eFuneral Interview Challenge

This is my solution to the eFuneral challenge. I'm using go 1.14.

## Configuration

Before running locally, some things need to be configured. This is done via a JSON file stored at the root of the project called `config.json`. The JSON schema I'm using looks like this (see `config.example.json`):

```json
{
  "TwilioAccountSID": "Twilio Account SID goes here",
  "TwilioAuthToken": "Twilio Auth Token goes here",
  "TwilioFromNumber": "Twilio From Number goes here",
  "SenderFirstName": "First name of the birthday wisher",
  "SenderLastName": "Last name of the birthday wisher",
  "SenderMobileNumber": "Mobile number of the birthday wisher"
}
```

The real config file contains things like API keys, so the one I'm using has been ignored in `.gitignore`.

## Running

To run the program, simply run `go run main.go`. That's it!

## Testing

This project comes with _some_ unit tests. I didn't test exhaustively, but wanted to make sure my code was working. To run the tests, run `go test ./...`.
