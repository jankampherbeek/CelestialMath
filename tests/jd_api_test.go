/*
 * Celestial Math.
 * Copyright (c) Jan Kampherbeek.
 * Celestial Math is open source.
 * Please check the file copyright.txt in the root of the source for further details.
 *
 */

package tests

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math"
	"net/http"
	"testing"
	"time"
)

const (
	baseUrl    = "http://localhost:8080"
	jdEndpoint = "/api/julian-day"
)

type Response struct {
	Message   string  `json:"message"`
	JulianDay float64 `json:"julianDay"`
}

func TestGetEndpoint(t *testing.T) {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	dateTimeForTest := "?year=1977&month=4&day=26&hours=9&minutes=36&seconds=0&gregorian=true"

	// Make GET request
	resp, err := client.Get(baseUrl + jdEndpoint + dateTimeForTest)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Check status code
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// Print raw response for debugging
	fmt.Printf("Raw response: %s\n", string(body))

	// Parse JSON response
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatalf("Failed to parse JSON response: %v", err)
	}
	expected := 2443259.9
	if math.Abs(response.JulianDay-expected) > 1e-8 {
		t.Errorf("JulianDay = %f, expected %f", response.JulianDay, expected)
	}

}
