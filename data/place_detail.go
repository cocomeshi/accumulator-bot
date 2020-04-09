package data

type Detail struct {
	ResultData Result `json:"result"`
}

type Result struct {
	Address   string       `json:"formatted_address"`
	OpenClose OpeningHours `json:"opening_hours"`
}

type OpeningHours struct {
	Periods ADayOpening `json:"periods"`
}

type ADayOpening struct {
	Open  DayTime `json:"open"`
	Close DayTime `json:"close"`
}

type DayTime struct {
	Day  int64  `json:"day"`
	Time string `json:"time"`
}
