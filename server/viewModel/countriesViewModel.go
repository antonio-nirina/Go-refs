package viewModel

import "github.com/antonio-nirina/go-refs/server/model"

type CountriesViewModel struct {
	Name string  `json:"name"`
  	Cca2 string  `json:"cca2"`
  	Cca3 string  `json:"cca3"`
  	Region string`json:"region"`
  	Subregion string `json:"subregion"`
  	Currency model.CurrenciesModel `json:"currency"`
}

