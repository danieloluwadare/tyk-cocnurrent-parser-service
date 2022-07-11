package service

import (
	"github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/model"
)

type Parser interface {
	Parse() ([]model.TykTaskConfig, error)
}
