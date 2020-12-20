package amka_test

import (
	"fmt"
	"testing"

	"github.com/zoispag/amka-go"
)

func ExampleValidate() {
	isValid, err := amka.Validate("09095986680")
	if !isValid {
		fmt.Printf("Amka is not valid: %s", err)
	} else {
		fmt.Println("Amka is valid")
	}
	// Output: Amka is not valid: Provided AMKA does not have a valid checkdigit
}

func Test_Validate(t *testing.T) {
	tests := []struct {
		name string
		amka string
		want bool
	}{
		{
			"Valid AMKA",
			"09095986684",
			true,
		},
		{
			"Not a number",
			"asvs",
			false,
		},
		{
			"Short number",
			"09095986",
			false,
		},
		{
			"Long number",
			"090959866845",
			false,
		},
		{
			"Not a date",
			"39095986681",
			false,
		},
		{
			"Bad checkdigit",
			"09095986680",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := amka.Validate(tt.amka)
			if got != tt.want {
				t.Errorf("validate(%q) got = %v, want %v", tt.amka, got, tt.want)
			}
		})
	}
}
