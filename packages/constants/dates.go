package dates

//this const can be used in other packages
//for example weekly appointment count
const DaysInWeek = 7

func WeekToDays(week int) int {
	return week * DaysInWeek
}

func DaysToWeeks(days int) float64 {
	return float64(days) / float64(DaysInWeek)
}