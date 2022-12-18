package response

type Dashboard struct {
	Booking int
}

type DashboardVaccine struct {
	Name string
	Dose int
}

type IsCloseFalse struct {
	Active int
}

type SessionFinished struct {
	Amount int
}

type VaccinatedUser struct {
	Vaccinated int
}

type RegisterStatistic struct {
	RegisteredStat []DashboardForm
}

type VaccineStatistic struct {
	VaccineStat []DashboardForm
}

type DashboardForm struct {
	Name      string
	DoseOne   int
	DoseTwo   int
	DoseThree int
}
