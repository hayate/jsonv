package jsonv

import (
	"fmt"
)

/*
Logical OR. Ensures that one of the validator validate the data. Value in `V1` and `V2` must be present.

Note: `V2` is tried only if `V1` doesn't validate.
*/
type Or struct {
	V1 Validator
	V2 Validator
}

func (self *Or) Validate(data *interface{}) (outdesc string, err error) {

	desc, err1 := self.V1.Validate(data)
	if err1 == nil {
		return fmt.Sprintf("Or(V1)->%s", desc), nil
	}
	desc, err2 := self.V2.Validate(data)
	if err2 == nil {
		return fmt.Sprintf("Or(V2)->%s", desc), nil
	}
	return "Or", fmt.Errorf("V1:%s. V2:%s", err1, err2)
}
