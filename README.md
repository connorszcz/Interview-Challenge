# eFuneral Interview Challenge

This is my solution to the eFuneral challenge. I'm using go 1.14.

## Configuration

Before running locally, some things need to be configured. This is done via a JSON file stored at the root of the project called `config.json`. The JSON schema I'm using looks like this (see `config.example.json`):

```json
{
  "twilio_account_sid": "Twilio Account SID goes here",
  "twilio_auth_token": "Twilio Auth Token goes here",
  "twilio_from_number": "Twilio From Number goes here",
  "sender_first_name": "First name of the birthday wisher",
  "sender_last_name": "Last name of the birthday wisher",
  "sender_mobile_number": "Mobile number of the birthday wisher"
}
```

The real config file contains things like API keys, so the one I'm using has been ignored in `.gitignore`.

## Running

To run the program, simply run `go run main.go`. That's it!

The tool provides some command-line options to configure it. You can get help at any time with

```
$ go run main.go --help
```

By default, the tool sources the address book and config files from "addressBook.csv" and "config.json", respectively. You can specify the paths to the address book and config files to override the defaults

```
$ go run main.go --address-book="addressBook.csv" --config-file="config.json"
```

## Testing

This project comes with _some_ unit tests. I didn't test exhaustively, but wanted to make sure my code was working. To run the tests, run `go test ./...`.
