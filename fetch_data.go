package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Define your structs
type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	FirstAlbum   string   `json:"firstAlbum"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
}

type Location struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}

type Date struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type Relation struct {
	Index []struct {
		ID        int                 `json:"id"`
		Relations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

// Fetch data from URL and unmarshal into target struct
func fetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, target)
}

// Function to fetch all data
func getData() (artists []Artist, locations Location, dates Date, relations Relation, err error) {
	artistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	locationsURL := "https://groupietrackers.herokuapp.com/api/locations"
	datesURL := "https://groupietrackers.herokuapp.com/api/dates"
	relationsURL := "https://groupietrackers.herokuapp.com/api/relation"

	err = fetchData(artistsURL, &artists)
	if err != nil {
		return
	}

	err = fetchData(locationsURL, &locations)
	if err != nil {
		return
	}

	err = fetchData(datesURL, &dates)
	if err != nil {
		return
	}

	err = fetchData(relationsURL, &relations)
	if err != nil {
		return
	}

	return artists, locations, dates, relations, nil
}
