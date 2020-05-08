package utils

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// GetLastMonthStart return the first day of last month in format YYYY-DD-MM
func GetLastMonthStart() string {
	t := time.Now()
	log.WithFields(log.Fields{
		"time": t,
	}).Debug("current time:")

	// this year, last month, first day of the month, hour, min, sec, nanosec, location
	d := time.Date(t.Year(), t.Month()-1, 1, 12, 30, 0, 0, t.Location())

	log.WithFields(log.Fields{
		"time": d,
	}).Debug("start of last month:")

	// this is the **magical reference date**
	return d.Format("2006-01-02")
}

// GetLastMonthEnd return the first day of this month (considered the end of billing period) in format YYYY-DD-MM
func GetLastMonthEnd() string {
	t := time.Now()
	log.WithFields(log.Fields{
		"time": t,
	}).Debug("current time:")

	// this year, last month, first day of the month, hour, min, sec, nanosec, location
	d := time.Date(t.Year(), t.Month(), 1, 12, 30, 0, 0, t.Location())

	log.WithFields(log.Fields{
		"time": d,
	}).Debug("end of last month:")

	// this is the **magical reference date**
	return d.Format("2006-01-02")
}
