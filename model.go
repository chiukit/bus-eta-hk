package main

import "time"

type ETAResponse struct {
	Type               string    `json:type`
	Version            string    `json:version`
	GeneratedTimestamp time.Time `json:generated_timestamp,string`
}

type StopETAResponse struct {
	ETAResponse
	Data []ETA `json:data`
}

type StopListResponse struct {
	ETAResponse
	Data []StopInput `json:data`
}

type StopInput struct {
	StopID string  `json:"stop"`
	NameEn string  `json:"name_en"`
	NameTC string  `json:"name_tc"`
	NameSC string  `json:"name_sc"`
	Lat    float64 `json:"lat,string"`
	Long   float64 `json:"long,string"`
}

type Stop struct {
	StopID   string           `json:"stop"`
	NameEn   string           `json:"name_en"`
	NameTC   string           `json:"name_tc"`
	NameSC   string           `json:"name_sc"`
	Lat      float64          `json:"lat"`
	Long     float64          `json:"long"`
	Distance float64          `json:"distance"`
	Geohash  string           `json:"geohash"`
	ETA      map[string][]ETA `json:"eta"`
}

type ETA struct {
	Co             string    `json:"co"`
	DataTimestamp  time.Time `json:"data_timestamp,string"`
	DestEn         string    `json:"dest_en"`
	DestSc         string    `json:"dest_sc"`
	DestTc         string    `json:"dest_tc"`
	Dir            string    `json:"dir"`
	Eta            time.Time `json:"eta,string"`
	EtaSeq         int64     `json:"eta_seq"`
	RmkEn          string    `json:"rmk_en"`
	RmkSc          string    `json:"rmk_sc"`
	RmkTc          string    `json:"rmk_tc"`
	Route          string    `json:"route"`
	Seq            int64     `json:"seq"`
	ServiceType    int64     `json:"service_type"`
	EtaMinutesDiff float64   `json:"eta_minutes_diff"`
}

type Global struct {
	Stops []Stop
}
