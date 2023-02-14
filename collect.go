package main

import (
	"errors"
	"io"
	"net/http"
)

func Get_all_data() (io.ReadCloser, error) {
	if data_URL == "" {
		return nil, errors.New("CARBOFRA_DATA_URL is not set")
	}

	resp, err := http.Get(data_URL)
	if err != nil {
			return nil, errors.New("Error while getting data from " + data_URL)
	}

	if resp.StatusCode != http.StatusOK {
			return nil, errors.New("Error while getting data from " + data_URL + " : " + resp.Status)
	}

	return resp.Body, nil
}