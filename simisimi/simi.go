package simisimi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	API_HOST string = "https://api.akuari.my.id/simi/simi?query="
)

func GetSimi(query string) string {
	uri := API_HOST + encodeURL(query)

	resp, err := http.Get(uri)
	if err != nil {
		return err.Error()
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	// Read body
	body, err := io.ReadAll(resp.Body)

	var responseapi Response
	err = json.Unmarshal(body, &responseapi)
	if err != nil {
		return "error unmarshal: " + err.Error()
	}

	return responseapi.Respon
}

func encodeURL(uri string) string {
	s := url.QueryEscape(uri)
	return s
}
