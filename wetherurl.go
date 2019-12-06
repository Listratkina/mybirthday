package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
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

type pageData struct {
	DayData []dayData
}
type dayData struct {
	Day       time.Time
	Temperature float64
	Hours float64
}

// Configuration
var (
	key             = "bbb8fa34a5d604ee97fc430999b4f172" // register with Dark Sky API to obtain your own
	vladivostokCoords = "43.10562,131.87"
)

func main() {
	http.HandleFunc("/", webHandler)
	fmt.Println("Open http://localhost:8080/ in your browser")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	d := time.Date(1988, 02, 12, 0, 0, 0, 0, time.UTC)
	weekday := d.Weekday()
	mt := float64(-100)
	md := d
	pageData := pageData{}

	for i := 0; i < 32; i++ {
		d = d.AddDate(1, 0, 0)
		weekday = d.Weekday()

		// See https://darksky.net/dev/docs#time-machine-request

	url := fmt.Sprintf("https://api.darksky.net/forecast/%v/%v,%v?exclude=currently,flags?units=si", key, vladivostokCoords, d.Format(time.RFC3339))
		fmt.Println(url)
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
				mtd = t                                //макс темп за день
				timed = (hourlyData["time"]).(float64) //время макс темп за день
			}
			if mtd > mt { //макс темп за день > МАКС темп всего?
				mt = mtd //запоминает новую макс темп
				md = d   //запоминает день макс темп
			}
		pageData.DayData = append(pageData.DayData, dayData{Day: d, Temperature: mtd, Hours: timed} )
	}

	tplStr, err := ioutil.ReadFile("template.html")
	check(err)
	tpl, err := template.New("page").Parse(string(tplStr))
	check(err)
	tpl.Execute(w, pageData)
}
fmt.Println(mt, md.Format("2006-01-02"),weekday) //печать макс темп всего и когда была
}
