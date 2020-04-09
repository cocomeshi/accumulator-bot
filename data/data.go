package data

type Restaurant struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Rating   float64  `json:"rating"`
	Geometry Geometry `json:"geometry"`
	Types    []string `json:"types"`
}

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type Geometry struct {
	Location Location `json:"location"`
}

type RestaurantResponse struct {
	HitNum        int8         `json:"hit_per_page"`
	NextPageToken string       `json:"next_page_token"`
	Restaurants   []Restaurant `json:"results"`
}
