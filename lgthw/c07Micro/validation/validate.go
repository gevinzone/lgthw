package validation

import "errors"

type Verror struct {
	error
}

func ValidatePayload(p *Payload) error {
	if p.Name == "" {
		return Verror{error: errors.New("name is required")}
	}
	if p.Age <= 0 || p.Age >= 120 {
		return Verror{errors.New("age is required and must be a value greater than 0 and less than 120")}
	}
	return nil
}
