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
)

// ObliquityCalculator calculates the mean obliquity of the earth's axis
type ObliquityCalculator interface {
	CalcObl(jd float64) float64
}

type ObliquityCalculation struct {
}

func NewObliquityCalculation() ObliquityCalculator {
	return ObliquityCalculation{}
}

// CalcObl calculates the mean obliquity with the formul fro Jean Meeus,
// Astronomical Algorithms, p. 147
func (oc ObliquityCalculation) CalcObl(jd float64) float64 {

	u := (jd - 2451545) / 3652500.0 // factor t divided by 100
	baseValue := 23.0 + 26.0/60.0 + 21.448/3600.0
	corr1 := -1.30025833333 * u
	corr2 := -0.000430555555556 * math.Pow(u, 2)
	corr3 := 0.55534722222222 * math.Pow(u, 3)
	corr4 := 0.01427222222222 * math.Pow(u, 4)
	corr5 := -0.06935277777778 * math.Pow(u, 5)
	corr6 := -0.10847222222222 * math.Pow(u, 6)
	corr7 := 0.00197777777778 * math.Pow(u, 7)
	corr8 := 0.00774166666667 * math.Pow(u, 8)
	corr9 := 0.00160833333333 * math.Pow(u, 9)
	corr10 := 0.00068055555556 * math.Pow(u, 10)
	allCor := corr1 + corr2 + corr3 + corr4 + corr5 + corr6 + corr7 + corr8 + corr9 + corr10
	obliquity := baseValue + allCor
	return obliquity
}
