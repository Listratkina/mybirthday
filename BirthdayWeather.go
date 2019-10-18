package main

import (
	"encoding/json" 
	"fmt"           
	"net/http"      
	"time"          
)

func check(err error) { 
	if err != nil {
		panic(err)
	}
}

type jsonObject = map[string]interface{}
type jsonArray = []interface{}

// Configuration
var (
	key               = "bbb8fa34a5d604ee97fc430999b4f172" // register with Dark Sky API to obtain your own
	vladivostokCoords = "43.10562,131.87"
)

func main() {

	d := time.Date(1988, 02, 12, 0, 0, 0, 0, time.UTC)
	weekday := d.Weekday()
	mt := float64(-100)
	md := d
	for i := 0; i < 32; i++ {
		d = d.AddDate(1, 0, 0)
		weekday = d.Weekday()
		fmt.Println(d, weekday)

		// See https://darksky.net/dev/docs#time-machine-request
		url := fmt.Sprintf("https://api.darksky.net/forecast/%v/%v,%v?exclude=currently,flags?units=si", key, vladivostokCoords, d.Format(time.RFC3339))
		//fmt.Println(url)
		resp, err := http.Get(url)
		check(err)

		// See https://darksky.net/dev/docs#response-format
		var decodedResponse jsonObject
		check(json.NewDecoder(resp.Body).Decode(&decodedResponse))
		hours := decodedResponse["hourly"].(jsonObject)["data"].(jsonArray)
		timed := float64(0)
		mtd := float64(-100)
		for _, h := range hours {
			hourlyData := h.(jsonObject)
			if hourlyData["temperature"] == nil {
				continue
			}
			t := (hourlyData["temperature"].(float64) - 32) / 1.8
			if t > mtd {
				mtd = t
				timed = (hourlyData["time"]).(float64)
			}

		}
		if mtd > mt {
			mt = mtd
			md = d
		}
		fmt.Printf("%v C\n", mtd)
		fmt.Println(time.Unix(int64(timed), 0))
	}
	fmt.Println(mt, md.Format("2006-01-02"))
}
