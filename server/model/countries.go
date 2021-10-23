package model

// C'est trop comme json ouff!!!!!

type Countries struct {
  Name countrynameElements `json:"name"`
  Cca2 string  `json:"cca2"`
  Cca3 string  `json:"cca3"`
  Currencies CurrenciesModel  `json:"currencies"`
  Region string  `json:"region"`
  Subregion string  `json:"subregion"`
  
}

type countrynameElements struct {
  Common string  `json:"common"`
  Official string  `json:"official"`
  NativeName nameCountries  `json:"NativeName"`
}

type nameElements struct {
  Common string  `json:"common"`
  Official string  `json:"official"`
}
type nameCountries map[string]nameElements

type currenciesModelInit struct {
  Name string  `json:"name"`
  Symbol string  `json:"symbol"`
}

type CurrenciesModel map[string]currenciesModelInit





