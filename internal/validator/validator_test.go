package validator

import "testing"

func TestCreateUserInputValidation(t *testing.T) {
    tests := []struct {
        name  string
        input CreateUserInput
        valid bool
    }{
        {"valid input", CreateUserInput{"John Doe", "john@example.com", 25}, true},
        {"empty name", CreateUserInput{"", "john@example.com", 25}, false},
        {"long name", CreateUserInput{string(make([]byte, 300)), "john@example.com", 25}, false},
        {"invalid age", CreateUserInput{"John", "john@example.com", 17}, false},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            err := tc.input.Validate()
            if tc.valid && err != nil {
                t.Errorf("expected valid but got error: %v", err)
            }
            if !tc.valid && err == nil {
                t.Errorf("expected error but got none")
            }
        })
    }
}
