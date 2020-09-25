package gcToWg

func To_gps (lat, lng float64) (float64, float64) {
	d_lat, d_lng := delta(lat, lng)
	return lat - d_lat, lng -d_lng
}
