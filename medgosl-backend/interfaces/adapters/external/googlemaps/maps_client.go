package googlemaps

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	apikey     string
	httpClient *http.Client
}

func NewClient(apikey string) *Client {
	return &Client{
		apikey:     apikey,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

type geocodeResponse struct {
	Status  string `json:"status"`
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}

func (c *Client) Geocode(ctx context.Context, address string) (float64, float64, error) {
	endpoint := "https://maps.googleapis.com/maps/api/geocode/json"
	q := url.Values{}
	q.Set("address", address)
	q.Set("apikey", c.apikey)

	req, _ := http.NewRequestWithContext(ctx, "GET", endpoint+"?"+q.Encode(), nil)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var out geocodeResponse
	// Decodes from json
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return 0, 0, fmt.Errorf("error in decoding from json %s", err)
	}

	// check if the status was okay and and if the len of results is equal to zero
	if out.Status != "OK" || len(out.Results) == 0 {
		return 0, 0, fmt.Errorf("geocode failed: %s", out.Status)
	}

	return out.Results[0].Geometry.Location.Lat, out.Results[0].Geometry.Location.Lng, nil

}

type DistanceMatrixResponse struct {
	Rows []struct {
		Elements []struct {
			Status string `json:"status"`

			Distance struct {
				Value int    `json:"value"`
				Text  string `json:"text"`
			} `json:"Distance"`

			Duration struct {
				Value int    `json:"value"`
				Text  string `json:"text"`
			} `json:"duration"`
		} `json:"elements"`
	} `json:"rows"`

	Status string `json:"status"`
}

func (c *Client) DistanceMatrix(ctx context.Context, origins, destinations []string) (*DistanceMatrixResponse, error) {
	endpoint := "https://maps.googleapis.com/maps/api/distancematrix/json"
	q := url.Values{}
	q.Set("origins", joinPipe(origins))
	q.Set("destination", joinPipe(destinations))
	q.Set("apikey", c.apikey)

	req, _ := http.NewRequestWithContext(ctx, "GET", endpoint+"?"+q.Encode(), nil)
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}
	defer resp.Body.Close()

	var out DistanceMatrixResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	if out.Status != "OK" {
		return nil, fmt.Errorf("distance matrix failed: %s", out.Status)
	}
	return &out, nil

}

func joinPipe(items []string) string {
	return url.QueryEscape(textJoin(items, "|"))
}

func textJoin(items []string, sep string) string {
	var s string
	for i, v := range items{
		if i == 0 {
			s = v
		} else {
			s += sep + v
		}
	}
	return s
}
