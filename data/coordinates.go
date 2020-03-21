package data

type Coordinates struct {
	Longitude float64
	Latitude  float64
}

func NewCoordinate(longitude float64, latitude float64) *Coordinates {
	return &Coordinates{
		Longitude: longitude,
		Latitude:  latitude,
	}
}
