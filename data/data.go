package data

type Restaurant struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	Address    string `json:"address"`
	Opentime   string `json:"opentime"`
	Categories string `json:"category"`
}

type RestaurantResponse struct {
	HitNum int8         `json:"hit_per_page"`
	Rests  []Restaurant `json:"rest"`
}
