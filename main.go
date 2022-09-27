package main

import (
	"fmt"
	"net/http"
	"sort"
)

func stops(w http.ResponseWriter, req *http.Request, g Global) {
	latlng, err := ParseLatLng(req)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	// assume API must contain latlng and range
	stops := FilterStopByLatLng(g.Stops, latlng[0], latlng[1], ParseDistance(req))

	// get ETA in batch by Stops
	stops = BatchGetETA(stops)

	// filter by route + direction (if set)
	route := req.URL.Query().Get("route")
	direction := req.URL.Query().Get("dir")

	if route != "" && direction != "" {
		stops = FilterStopByRoute(stops, route)
		stops = FilterETAByRoute(stops, route)
		stops = FilterStopByDirection(stops, direction)
	}

	// add distance to each stop
	stops = Map(stops, func(v Stop) Stop {
		v.Distance = Distance(latlng[0], latlng[1], v.Lat, v.Long, "K")
		return v
	})

	// order by distance (asc order)
	sort.Slice(stops, func(i, j int) bool {
		return stops[i].Distance < stops[j].Distance
	})

	WriteJSON(w, http.StatusOK, stops)
}

func main() {
	g := Global{}
	g.Stops = FetchAllStops()

	http.HandleFunc("/stops", func(w http.ResponseWriter, r *http.Request) {
		stops(w, r, g)
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, "pong1")
	})

	fmt.Println("Server started...")
	http.ListenAndServe(":8090", nil)
}
