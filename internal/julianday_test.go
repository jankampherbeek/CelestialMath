/*
 * Celestial Math.
 * Copyright (c) Jan Kampherbeek.
 * Celestial Math is open source.
 * Please check the file copyright.txt in the root of the source for further details.
 *
 */

package internal

import (
	"math"
	"testing"
	"time"
)

func TestJulDayCalculation_CalcJd_Gregorian(t *testing.T) {
	calc := NewJulDayCalculation()
	dateTime := time.Date(1957, 10, 4, 19, 26, 24, 0, time.UTC)
	expected := 2436116.31
	result := calc.CalcJd(dateTime, true)
	if math.Abs(result-expected) > 1e-8 {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestJulDayCalculation_CalcJd_Julian(t *testing.T) {
	calc := NewJulDayCalculation()
	dateTime := time.Date(333, 1, 27, 12, 0, 0, 0, time.UTC)
	expected := 1842713.0
	result := calc.CalcJd(dateTime, false)
	if math.Abs(result-expected) > 1e-8 {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
