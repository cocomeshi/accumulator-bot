package data

type Restaurant struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Url        string  `json:"url"`
	Address    string  `json:"address"`
	Opentime   string  `json:"opentime"`
	Categories string  `json:"category"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

type RestaurantResponse struct {
	HitNum      int8         `json:"hit_per_page"`
	Restaurants []Restaurant `json:"rest"`
}
