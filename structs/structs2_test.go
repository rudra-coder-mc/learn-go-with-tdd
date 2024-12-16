package structs

import (
	"fmt"
	"testing"
)

func TestCanSendMessage(t *testing.T) {
	type testCase struct {
		name     string
		mToSend  messageToSend
		expected bool
	}

	tests := []testCase{
		{
			name: "Valid message with complete sender and recipient details",
			mToSend: messageToSend{
				message:   "You have an appointment tomorrow",
				sender:    user{name: "Brenda Halafax", number: 16545550987},
				recipient: user{name: "Sally Sue", number: 19035558973},
			},
			expected: true,
		},
		{
			name: "Missing sender name",
			mToSend: messageToSend{
				message:   "You have an event tomorrow",
				sender:    user{number: 16545550987},
				recipient: user{name: "Suzie Sall", number: 19035558973},
			},
			expected: false,
		},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{
				name: "Missing recipient number",
				mToSend: messageToSend{
					message:   "You have a birthday tomorrow",
					sender:    user{name: "Jason Bjorn", number: 16545550987},
					recipient: user{name: "Jim Bond"},
				},
				expected: false,
			},
			{
				name: "Missing sender number",
				mToSend: messageToSend{
					message:   "You have a party tomorrow",
					sender:    user{name: "Njorn Halafax"},
					recipient: user{name: "Becky Sue", number: 19035558973},
				},
				expected: false,
			},
			{
				name: "Valid recipient but missing sender name",
				mToSend: messageToSend{
					message:   "You have a celebration tomorrow",
					sender:    user{number: 16545550987},
					recipient: user{number: 19035558973},
				},
				expected: false,
			},
		}...)
	}

	passCount := 0
	failCount := 0

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test #%d: %s", i+1, test.name), func(t *testing.T) {
			output := canSendMessage(test.mToSend)
			if output != test.expected {
				failCount++
				t.Errorf(`Test Failed:
Inputs:
* message:          %s
* sender.name:      %s
* sender.number:    %d
* recipient.name:   %s
* recipient.number: %d
Expected:           %v
Actual:             %v`,
					test.mToSend.message,
					test.mToSend.sender.name,
					test.mToSend.sender.number,
					test.mToSend.recipient.name,
					test.mToSend.recipient.number,
					test.expected,
					output)
			} else {
				passCount++
				t.Logf(`Test Passed:
Inputs:
* message:          %s
* sender.name:      %s
* sender.number:    %d
* recipient.name:   %s
* recipient.number: %d
Expected:           %v
Actual:             %v`,
					test.mToSend.message,
					test.mToSend.sender.name,
					test.mToSend.sender.number,
					test.mToSend.recipient.name,
					test.mToSend.recipient.number,
					test.expected,
					output)
			}
		})
	}

	t.Logf("---------------------------------")
	t.Logf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is a compile-time variable.
var withSubmit = true
