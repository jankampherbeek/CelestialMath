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
)

func TestObliquityCalculation_CalcObl(t *testing.T) {
	jd := 2446895.5
	oblCalc := NewObliquityCalculation()
	expected := 23.0 + 26.0/60.0 + 27.407/3600.0
	result := oblCalc.CalcObl(jd)
	if math.Abs(result-expected) > 1e-5 {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
