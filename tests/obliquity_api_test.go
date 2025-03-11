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
	baseUrlObl  = "http://localhost:8080"
	oblEndpoint = "/api/obliquity"
)

type OblResponse struct {
	Message   string  `json:"message"`
	Obliquity float64 `json:"obliquity"`
}

func TestGetEndpointObl(t *testing.T) {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	jdForTest := "?jd=2457695.387152778"

	// Make GET request
	resp, err := client.Get(baseUrlObl + oblEndpoint + jdForTest)
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
	var response OblResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatalf("Failed to parse Obliquity response: %v", err)
	}
	expected := 23.437101628
	if math.Abs(response.Obliquity-expected) > 1e-5 {
		t.Errorf("Obliquity = %f, expected %f", response.Obliquity, expected)
	}

}
