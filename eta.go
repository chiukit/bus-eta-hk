package main

import (
	"encoding/json"
	"log"
	"sync"
	"time"
)

func GetETAByStopID(id string) []ETA {
	body := Fetch("https://data.etabus.gov.hk/v1/transport/kmb/stop-eta/" + id)

	r := StopETAResponse{}
	err := json.Unmarshal(body, &r)
	if err != nil {
		log.Fatalln(err)
	}

	return Map(r.Data, CalETADiff)
}

func CalETADiff(v ETA) ETA {
	diffMinutes := v.Eta.Sub(v.DataTimestamp).Round(time.Minute).Minutes()
	if diffMinutes >= 0 {
		v.EtaMinutesDiff = diffMinutes
	} else {
		v.EtaMinutesDiff = -1
	}
	return v
}

func GroupByRoute(etas []ETA) map[string][]ETA {
	group := make(map[string][]ETA)
	for _, v := range etas {
		if _, exists := group[v.Route]; !exists {
			group[v.Route] = []ETA{}
		}
		group[v.Route] = append(group[v.Route], v)
	}

	return group
}

func BatchGetETA(stops []Stop) []Stop {
	s := []Stop{}

	wg := sync.WaitGroup{}

	for _, v := range stops {
		wg.Add(1)
		go func(v Stop) {
			defer wg.Done()

			etas := GetETAByStopID(v.StopID)
			v.ETA = GroupByRoute(etas)

			s = append(s, v)
		}(v)
	}
	wg.Wait()

	return s
}

func FetchAllStops() []Stop {
	body := Fetch("https://data.etabus.gov.hk/v1/transport/kmb/stop")

	f := StopListResponse{}
	err := json.Unmarshal(body, &f)
	if err != nil {
		log.Fatalln(err)
	}
	return Map(f.Data, CalGeohash)
}
