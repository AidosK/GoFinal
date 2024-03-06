package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
)

// A struct that holds all the information assciated with the form and embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// Valid, returns true if there are errors, otherwise return false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New, initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks for required fields using a Variadic function (...string). that is our function can have many types of string, where some maybe required while some won't
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field is required")
		}
	}
}

// HasField, checks if form field is in post and is not empty
func (f *Form) HasField(field string) bool {
	x := f.Get(field)
	return x != ""
}
func CheckEmail(email string) bool {
	reg := regexp.MustCompile("^((\"[\\w-\\s]+\")|([\\w-]+(?:\\.[\\w-]+)*)|(\"[\\w-\\s]+\")([\\w-]+(?:\\.[\\w-]+)*))(@((?:[\\w-]+\\.)*\\w[\\w-]{0,66})\\.([a-z]{2,6}(?:\\.[a-z]{2})?)$)|(@\\[?((25[0-5]\\.|2[0-4][0-9]\\.|1[0-9]{2}\\.|[0-9]{1,2}\\.))((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\\.){2}(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\\]?$)")

	return reg.Match([]byte(email))
}

// MinLength checks for field minimum length
func (f *Form) MinLength(field string, minimum int, maximum int) bool {
	x := f.Get(field)
	if len(x) < minimum {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", minimum))
		return false
	}
	if len(x) > maximum {
		f.Errors.Add(field, fmt.Sprintf("This field must not be longer than %d characters", maximum))
		return false
	}

	return true
}

// IsEmail checks for valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
