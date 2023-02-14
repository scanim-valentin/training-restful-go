package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
)

func get_best_places(left, bottom, right, top float32) []string {
	/*
	* left is the longitude of the left (westernmost) side of the bounding box.
	* bottom is the latitude of the bottom (southernmost) side of the bounding box.
	* right is the longitude of the right (easternmost) side of the bounding box.
	* top is the latitude of the top (northernmost) side of the bounding box.
	 */
	r := make([]string, 1)
	return r
}

func osm_get(left, bottom, right, top float64) {
	resp, err := http.Get(fmt.Sprintf("https://api.openstreetmap.org/api/0.6/map.json?bbox=%v,%v,%v,%v", left, bottom, right, top))
	if err != nil {
		fmt.Println("Unable to retrieve map data: \n", err)
	} else {

		var body []byte
		var err error

		defer resp.Body.Close()
		body, err = io.ReadAll(resp.Body)

		fmt.Println(reflect.TypeOf(body))
		if err != nil {
			fmt.Println("Unable to read body: \n", err)
		}
		err = os.WriteFile(fmt.Sprintf("map_%v_%v_%v_%v.json", left, bottom, right, top), body, 0666)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}
	}
}

func main() {
	// resp, err := http.Get("https://master.apis.dev.openstreetmap.org/api/0.6/map?bbox=7.0191821,49.2785426,7.0197485,49.2793101")
	// resp, err := http.Get("https://master.apis.dev.openstreetmap.org/api/0.6/map.json?bbox=9.08839,48.82086,9.09949,48.82469")
	coordinate_x, coordinate_y := 43.57041728011695, 1.46819945229431
	delta := float64(0.05)
	osm_get(coordinate_x-delta, coordinate_y-delta, coordinate_x+delta, coordinate_y+delta)
}
