/*
 *  Celestial Math.
 *  Copyright (c) Jan Kampherbeek.
 *  Celestial Math is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package main

import (
	"celestialmath/internal"
	_ "embed"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type JulianDayRequest struct {
	Year      int  `form:"year" binding:"required"`
	Month     int  `form:"month" binding:"required,min=1,max=12"`
	Day       int  `form:"day" binding:"required,min=1,max=31"`
	Hours     int  `form:"hours" binding:"min=0,max=23"`
	Minutes   int  `form:"minutes" binding:"min=0,max=59"`
	Seconds   int  `form:"seconds" binding:"min=0,max=59"`
	Gregorian bool `form:"gregorian"`
}

type JulianDayResponse struct {
	JulianDay float64          `json:"julianDay"`
	Input     JulianDayRequest `json:"input"`
}

func main() {
	router := gin.Default()

	// Endpoint registreren
	router.GET("/api/julian-day", GetJulianDay)

	// Start de server
	router.Run(":8080")
}

func GetJulianDay(c *gin.Context) {
	var req JulianDayRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input parameters",
			"details": err.Error(),
		})
		return
	}
	if !isValidDate(req.Year, req.Month, req.Day, req.Gregorian) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid date",
			"details": "The specified date does not exist in the calendar",
		})
		return
	}

	calc := internal.NewJulDayCalculation()
	dt := time.Date(req.Year, time.Month(req.Month), req.Day, req.Hours, req.Minutes, req.Seconds, 0, time.Local)
	greg := req.Gregorian
	jdn := calc.CalcJd(dt, greg)
	c.JSON(http.StatusOK, JulianDayResponse{
		JulianDay: jdn,
		Input:     req,
	})
}

func isValidDate(year, month, day int, gregorian bool) bool {
	daysInMonth := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// Handle leap year
	if month == 2 {
		if gregorian {
			if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
				daysInMonth[2] = 29
			}
		} else {
			if year%4 == 0 {
				daysInMonth[2] = 29
			}
		}
	}
	return day <= daysInMonth[month]
}
