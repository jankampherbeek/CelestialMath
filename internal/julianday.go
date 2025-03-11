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
	"time"
)

// JulianDayCalculator calculates the Julian Day.
type JulianDayCalculator interface {
	CalcJd(dateTime time.Time, gregorian bool) float64
}

type JulDayCalculation struct {
}

func NewJulDayCalculation() JulianDayCalculator {
	return JulDayCalculation{}
}

// CalcJd handles the calculation of a Julian day number.
func (jc JulDayCalculation) CalcJd(dateTime time.Time, gregorian bool) float64 {
	y := dateTime.Year()
	m := dateTime.Month()
	d := dateTime.Day()
	var b float64
	if m < 2 {
		y--
		m += 12
	}
	if gregorian {
		a := math.Floor(float64(y) / 100.0)
		b = 2 - a + math.Floor(a/4)
	} else { // Julian calendar
		b = 0.0
	}
	c := math.Floor(365.25 * float64(y+4716))
	e := math.Floor(30.6001 * float64(m+1))

	jd := c + e + float64(d) + b - 1524.5
	decHours := float64(dateTime.Hour()) + float64(dateTime.Minute())/60.0 + float64(dateTime.Second())/3600.0
	jd += decHours / 24.0
	return jd
}
