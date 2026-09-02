package excelize

import (
	"time"
)

const (
	nanosInADay    = float64((24 * time.Hour) / time.Nanosecond)
	dayNanoseconds = 24 * time.Hour
	maxDuration    = 290 * 364 * dayNanoseconds
	roundEpsilon   = 1e-9
)

var (
	daysInMonth           = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	excel1900Epoc         = time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	excel1904Epoc         = time.Date(1904, time.January, 1, 0, 0, 0, 0, time.UTC)
	excelMinTime1900      = time.Date(1899, time.December, 31, 0, 0, 0, 0, time.UTC)
	excelBuggyPeriodStart = time.Date(1900, time.March, 1, 0, 0, 0, 0, time.UTC).Add(-time.Nanosecond)
)

func timeToExcelTime(t time.Time, date1904 bool) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func shiftJulianToNoon(julianDays, julianFraction float64) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func fractionOfADay(fraction float64) (hours, minutes, seconds, nanoseconds int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

func julianDateToGregorianTime(part1, part2 float64) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func doTheFliegelAndVanFlandernAlgorithm(jd int) (day, month, year int) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

func timeFromExcelTime(excelTime float64, date1904 bool) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func ExcelDateToTime(excelDate float64, use1904Format bool) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func isLeapYear(y int) bool { _ = "STUB: not implemented"; return false }

func getDaysInMonth(y, m int) int { _ = "STUB: not implemented"; return 0 }

func validateDate(y, m, d int) bool { _ = "STUB: not implemented"; return false }

func formatYear(y int) int { _ = "STUB: not implemented"; return 0 }

func getDurationNumFmt(d time.Duration) int { _ = "STUB: not implemented"; return 0 }

func getTimeNumFmt(t time.Time) int { _ = "STUB: not implemented"; return 0 }
