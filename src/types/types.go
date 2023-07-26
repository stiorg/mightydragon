package types

import (
	"context"
)

type Service interface {
	GetCatFact(context.Context) (*CatFact, error)
}

type CatFactService struct {
	url string
}

type CatFact struct {
	Fact string `json:"fact"`
}
