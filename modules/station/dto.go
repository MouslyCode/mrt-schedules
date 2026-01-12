package station

type Station struct {
	Id   string `json:"uid"`
	Name string `json:"title"`
}

type StationResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
