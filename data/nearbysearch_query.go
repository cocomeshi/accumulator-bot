package data

import (
	"strconv"
)

type NearbysearchQuery struct {
	baseUrl   string
	key       string
	radius    string
	placetype string
	location  string
}

func NewNearbysearchQuery(key string, baseUrl string, radius string, placetype string) *NearbysearchQuery {
	return &NearbysearchQuery{
		baseUrl,
		key,
		radius,
		placetype,
		"",
	}
}

func (q *NearbysearchQuery) SetLocation(c Coordinates) {
	strLatitude := strconv.FormatFloat(c.Latitude, 'f', -1, 64)
	strLongitude := strconv.FormatFloat(c.Longitude, 'f', -1, 64)
	q.location = strLatitude + "," + strLongitude
}

func (q *NearbysearchQuery) QueryGen() string {
	return q.baseUrl + "location=" + q.location + "&radius=" + q.radius + "&type=" + q.placetype + "&key=" + q.key
}
