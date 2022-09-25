package main

import (
	"errors"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/mmcloughlin/geohash"
)

func Distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	radlat1 := float64(math.Pi * lat1 / 180)
	radlat2 := float64(math.Pi * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}

func ParseLatLng(req *http.Request) ([]float64, error) {
	latlngStr := req.URL.Query().Get("latlng")
	if latlngStr == "" {
		return nil, errors.New("empty latlng")
	}
	latlng := Map(strings.Split(latlngStr, ","), func(v string) float64 {
		newValue, _ := strconv.ParseFloat(v, 64)
		return newValue
	})

	if len(latlng) != 2 {
		return nil, errors.New("invalid latlng")
	}

	return latlng, nil
}

func ParseDistance(req *http.Request) float64 {
	defaultValue := 0.5
	r := req.URL.Query().Get("r")
	if r != "" {
		distRangeInt, _ := strconv.ParseFloat(r, 64)

		if distRangeInt > 1 {
			return 1
		} else {
			return distRangeInt
		}
	}
	return defaultValue
}

func FilterStopByLatLng(stops []Stop, lat float64, lng float64, rangeKM float64) []Stop {
	return Filter(stops, func(v Stop) bool {
		return Distance(lat, lng, v.Lat, v.Long, "K") <= rangeKM
	})
}

func FilterStopByRoute(stops []Stop, route string) []Stop {
	return Filter(stops, func(v Stop) bool {
		_, exists := v.ETA[route]
		return exists
	})
}

func FilterETAByRoute(stops []Stop, route string) []Stop {
	return Map(stops, func(s Stop) Stop {
		for mK := range s.ETA {
			if mK != route {
				delete(s.ETA, mK)
			}
		}
		return s
	})
}

func FilterStopByDirection(stops []Stop, direction string) []Stop {
	return Filter(stops, func(s Stop) bool {
		for _, mV := range s.ETA {
			return mV[0].Dir == direction
		}
		return false
	})
}

func CalGeohash(v Stop) Stop {
	v.Geohash = geohash.EncodeWithPrecision(v.Lat, v.Long, 8)
	return v
}
