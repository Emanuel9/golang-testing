package locations_provider

import (
	"api/domain/locations"
	"api/errors"
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/mercadolibre/golang-restclient/rest"
)

const (
	urlGetCountry = "https://api.mercadolibre.com/countries/%s"
)

func GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	response := rest.Get(fmt.Sprintf(urlGetCountry, countryId))
	if response == nil || response.Response == nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient response when trying to get country %s", countryId),
		}
	}

	if response.StatusCode > 299 {
		var apiError errors.ApiError
		if err := json.Unmarshal(response.Bytes(), &apiError); err != nil {
			return nil, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid error interface when getting country %s", countryId),
			}
		}
		return nil, &apiError
	}

	var result locations.Country
	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal country data fro %s", countryId),
		}
	}
	return &result, nil
}
