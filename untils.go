package gcToWg

import {
	"math"
}

const (
	PI float64 = 3.14159265358979324
	x_pi float64 = 0
	a float64 = 6378245.0
	ee float64 = 0.00669342162296594323
)

func delta (lat, lng float64) (float64, float64) {
	dlat := transform_lat(lng - 105.0, lat - 35.0)
	dlng := transform_lng(lng - 105.0, lat - 35.0)
	radLat := lat / 180.0 * PI
	magic := math.Sin(radLat)
	magic = 1 - ee * magic *magic
	sqrtMagic := math.Sqrt(magic)
	d_lat := (dlat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * PI)
	d_lng := (dlng * 180.0) / (a / sqrtMagic * math.Cos(radLat) * PI)
	return d_lat, d_lng
}

func transform_lat (x, y float64) (float64) {
	ret := -100.0 + 2.0 * x + 3.0 * y + 0.2 * y * y + 0.1 * x * y + 0.2 * math.Sqrt(math.Abs(x));
	ret += (20.0 * math.Sin(6.0 * x * PI) + 20.0 * math.Sin(2.0 * x * PI)) * 2.0 / 3.0;
	ret += (20.0 * math.Sin(y * PI) + 40.0 * math.Sin(y / 3.0 * PI)) * 2.0 / 3.0;
	ret += (160.0 * math.Sin(y / 12.0 * PI) + 320 * math.Sin(y * PI / 30.0)) * 2.0 / 3.0;
	return ret;
}

func transform_lng (x, y float64) (float64) {
	ret := 300.0 + x + 2.0 * y + 0.1 * x * x + 0.1 * x * y + 0.1 * math.Sqrt(math.Sbs(x));
	ret += (20.0 * math.Sin(6.0 * x * PI) + 20.0 * math.Sin(2.0 * x * PI)) * 2.0 / 3.0;
	ret += (20.0 * math.Sin(x * PI) + 40.0 * math.Sin(x / 3.0 * PI)) * 2.0 / 3.0;
	ret += (150.0 * math.Sin(x / 12.0 * PI) + 300.0 * math.Sin(x / 30.0 * PI)) * 2.0 / 3.0;
	return ret;
}