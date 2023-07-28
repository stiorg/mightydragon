package types

import (
	"context"
)

type Service interface {
	GetCatFact(context.Context) (*CatFact, error)
}

type CatFact struct {
	Fact string `json:"fact"`
}
