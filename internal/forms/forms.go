package forms

import (
	"fmt"
	"net/url"

	"github.com/asaskevich/govalidator"
)

// Form creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//New Initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors{},
	}
}

// Required checks all the required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if value == "" {
			f.Errors.Add(field, "This field cannot be empty")
		}
	}
}

// checks if the form has that field or that field is empty
func (f *Form) Has(field string) bool {
	x := f.Get(field)

	return x != ""
}

// MinLength checks if the provided field has required length
func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be atleast %d long", length))
		return false
	}
	return true
}

func (f *Form) IsEmail(field string) bool {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add("email", "Invalid Email Address")
		return false
	}
	return true
}
