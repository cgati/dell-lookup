package main

import (
	"time"
)

func InWarranty(end string) bool {
	e, err := time.Parse("2006-01-02T15:04:05", end)
	if err != nil {
		return false
	}
	now := time.Now()
	if e.After(now) {
		return true
	}
	return false
}

func AddWarrantyStatus(result DellResult) DellResult {
	for i, a := range result.Assets {
		for j, w := range a.Warranties.Warranty {
			result.Assets[i].Warranties.Warranty[j].InWarranty = InWarranty(w.EndDate)
		}
	}
	return result
}
