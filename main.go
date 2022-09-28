package main

import (
	"fmt"
	"net/http"
	"sort"
	"time"
)

func stops(w http.ResponseWriter, req *http.Request, g Global) {
	latlng, err := ParseLatLng(req)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	t1 := time.Now()
	// assume API must contain latlng and range
	stops := FilterStopByLatLng(g.Stops, latlng[0], latlng[1], ParseDistance(req))
	fmt.Println("lat,lon", time.Now().Sub(t1))

	t2 := time.Now()
	// get ETA in batch by Stops
	stops = BatchGetETA(stops)
	fmt.Println("eta", time.Now().Sub(t2))

	// filter by route + direction (if set)
	route := req.URL.Query().Get("route")
	direction := req.URL.Query().Get("dir")

	if route != "" && direction != "" {
		t3 := time.Now()
		stops = FilterStopByRoute(stops, route)
		fmt.Println("stop_route", time.Now().Sub(t3))

		t4 := time.Now()
		stops = FilterETAByRoute(stops, route)
		fmt.Println("eta_route", time.Now().Sub(t4))

		t5 := time.Now()
		stops = FilterStopByDirection(stops, direction)
		fmt.Println("eta_dir", time.Now().Sub(t5))
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

	t6 := time.Now()
	g.Stops = FetchAllStops()
	fmt.Println("all_stop", time.Now().Sub(t6))

	http.HandleFunc("/stops", func(w http.ResponseWriter, r *http.Request) {
		stops(w, r, g)
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, "pong")
	})

	fmt.Println("Server started...")
	http.ListenAndServe(":8090", nil)
}
