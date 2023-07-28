package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/stiorg/mightydragon/types"
)

func NewCatFactService(url string) types.Service {
	return &CatFactService{
		url: url,
	}
}

type CatFactService struct {
	url string
}

func (s *CatFactService) GetCatFact(ctx context.Context) (*types.CatFact, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fact := &types.CatFact{}

	if err := json.NewDecoder(resp.Body).Decode(fact); err != nil {
		return nil, err
	}

	return fact, nil
}
