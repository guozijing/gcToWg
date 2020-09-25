package gcToWg

func ToGps (lat, lng float64) (float64, float64) {
	dLat, dLng := delta(lat, lng)
	return lat - dLat, lng -dLng
}
