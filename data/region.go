package data

type Area struct {
	Kinki []Region `json:"kinki"`
}

type Region struct {
	Label      string `json:"label"`
	PointRange Range  `json:"point_range"`
}

type Range struct {
	Lat []float64 `json:"lat"`
	Lon []float64 `json:"lon"`
}
