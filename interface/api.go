package interface

import (
	"github.com/cocomeshi/accumulator-bot/data"
)

func NearbySearch() error {
	latitude := 34.726799
	longitude := 135.401687
	strLatitude := strconv.FormatFloat(latitude, 'f', -1, 64)
	strLongitude := strconv.FormatFloat(longitude, 'f', -1, 64)
	// TODO queryをstruct化する
	url := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?" + "location=" + strLatitude + "," + strLongitude + "&radius=3000&type=restaurant&key=" + key
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var datas data.RestaurantResponse
	if err := json.Unmarshal(body, &datas); err != nil {
		fmt.Println(err)
	}
	return datas, err
}