package authenticate

import (
	"errors"

	"github.com/vk-rv/pvx"
)

type NetSepioClaims struct {
	pvx.RegisteredClaims
	UserId string
}

func (c *NetSepioClaims) Valid() error {

	validationErr := &pvx.ValidationError{}
	if err := c.RegisteredClaims.Valid(); err != nil {
		errors.As(err, &validationErr)
		return err
	}

	return nil

}
