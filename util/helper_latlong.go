package util

import (
	"math"
)

type coordinate struct {
	lat float64
	lng float64
}

func FindRange(userLat, userLong, healthLat, healthLong float64) float64 {

	rs := coordinate{userLat, userLong}
	user := coordinate{healthLat, healthLong}

	kilo := distance(rs.lat, rs.lng, user.lat, user.lng, "K")
	kilo /= 1000
	result := roundFloat(kilo, 2)

	return result
}

func distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	radlat1 := float64(math.Pi * lat1 / 180)
	radlat2 := float64(math.Pi * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1609.344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
