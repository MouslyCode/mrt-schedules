package station

type Station struct {
	Id                string     `json:"nid"`
	Name              string     `json:"title"`
	StationEstimate   []Estimate `json:"estimasi"`
	StationFacility   []Facility `json:"fasilitas"`
	StationScheduleHI string     `json:"jadwal_hi_biasa"`
	StationScheduleLB string     `json:"jadwal_lb_biasa"`
}

type StationResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Schedule struct {
	ScheduleBundaranHI string `json:"jadwal_hi_biasa"`
	ScheduleLebakBulus string `json:"jadwal_lb_biasa"`
}

type ScheduleResponse struct {
	BundaranHIRegular []string `json:"bundaran_hi_regular"`
	LebakBulusRegular []string `json:"lebak_bulus_regular"`
}

func (s ScheduleResponse) Format(param1 string) {
	panic("unimplemented")
}

type StationScheduleResponse struct {
	StationName string             `json:"station"`
	Schedules   []ScheduleResponse `json:"schedules"`
}

type StationEstimateResponse struct {
	StationName string             `json:"station"`
	Estimates   []EstimateResponse `json:"estimates"`
}

type Estimate struct {
	StationId string `json:"stasiun_nid"`
	Fare      string `json:"tarif"`
	Time      string `json:"waktu"`
}

type EstimateResponse struct {
	StationName string `json:"to_station"`
	Fare        string `json:"fare"`
	Time        string `json:"time"`
}

type Facility struct {
	Id    string `json:"nid"`
	Title string `json:"title"`
	Type  string `json:"jenis_fasilitas"`
	Img   string `json:"cover"`
}

type FacilityResponse struct {
	Title string `json:"title"`
	Type  string `json:"jenis_fasilitas"`
	Img   string `json:"cover"`
}

type StationFacilityResponse struct {
	StationName string             `json:"station"`
	Facilities  []FacilityResponse `json:"facilities"`
}
