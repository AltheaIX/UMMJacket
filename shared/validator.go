package shared

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"sync"
)

var once sync.Once
var v *validator.Validate

func GetValidator() *validator.Validate {
	once.Do(func() {
		log.Info().Msg("Validator Initialized")
		v = validator.New()
	})

	return v
}
