package data

import "fmt"

const (
	ByLatitude  float64 = 0.0555560
	ByLongitude float64 = 0.0666672
)

// 経度、緯度の範囲データ型
type PointerScope struct {
	Current    Coordinates
	PointRange Range
}

// 次の走査座標の値で上書きする
func (ps *PointerScope) Next() error {
	newLon := ps.Current.Longitude + ByLongitude
	if isInLogitudeRange(newLon, ps.PointRange) {
		ps.Current = Coordinates{
			Latitude:  ps.Current.Latitude,
			Longitude: ps.Current.Longitude + ByLongitude,
		}
	} else {
		newLat := ps.Current.Latitude
		if isInLatitudeRange(newLat, ps.PointRange) {
			ps.Current = Coordinates{
				Latitude:  ps.Current.Latitude + ByLatitude,
				Longitude: ps.Current.Longitude,
			}
		} else {
			return fmt.Errorf("Error: %s", "search range over!")
		}
	}

	return nil

}

// 緯度が範囲内であるかをテストする
func isInLatitudeRange(lat float64, latRange Range) bool {
	return latRange.Lat[0] >= lat || latRange.Lat[1] <= lat
}

// 経度が範囲内であるかをテストする
func isInLogitudeRange(lon float64, lonRange Range) bool {
	return lonRange.Lon[0] >= lon || lonRange.Lon[1] <= lon
}
