package controller

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/antonio-nirina/go-refs/server/model"
	"github.com/antonio-nirina/go-refs/server/viewModel"
)

const URL_V3 = "https://restcountries.com/v3.1/"

type HttpClient struct {
	client *http.Client
}

var tr = &http.Transport{
	TLSClientConfig: &tls.Config{
		Renegotiation:      tls.RenegotiateOnceAsClient,
		InsecureSkipVerify: true,
	},
}

func Handle(fn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func CountriesHandler(w http.ResponseWriter, r *http.Request) {
	apiClient := &HttpClient{}
	apiClient.client = http.DefaultClient
	apiClient.client.Transport = tr
	var region string = "region/"+r.URL.Query().Get("region") 

	if r.URL.Query().Get("region")  == "" {
		region = "region/europe"
	}
	
	url := fmt.Sprintf("%s%s",URL_V3,region)
	request, err := http.NewRequest("GET", url, nil)
	
	if err != nil {
		handleError(err,w)
		return
	}

	response, err := apiClient.client.Do(request)

	if err != nil {
		handleError(err,w)
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	countryModel := []model.Countries{}
	err = json.Unmarshal(body, &countryModel)

	if err != nil {
		handleError(err,w)
		return
	}

	var responseViewModel []viewModel.CountriesViewModel

	for _,val := range countryModel {
		respViewModel := viewModel.CountriesViewModel{
			Name:val.Name.Official,
			Cca2:val.Cca2,
			Cca3:val.Cca3,
			Region:val.Region,
			Subregion:val.Subregion,
			Currency:val.Currencies,
		}
		responseViewModel = append(responseViewModel, respViewModel)
	}

	err = json.NewEncoder(w).Encode(responseViewModel)

	if err != nil {
		handleError(err, w)
		return
	}
}

func handleError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(fmt.Sprintf(`{"err": "%s"}`, err.Error())))
}