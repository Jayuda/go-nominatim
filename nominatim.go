package gonominatim

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ReverseData struct {
	PlaceID     int   `json:"place_id"`
	Licence     string   `json:"licence"`
	OsmType     string   `json:"osm_type"`
	OsmID       int   `json:"osm_id"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	PlaceRank   int   `json:"place_rank"`
	Category    string   `json:"category"`
	Type        string   `json:"type"`
	Importance  float32   `json:"importance"`
	Addresstype string   `json:"addresstype"`
	Name        string   `json:"name"`
	DisplayName string   `json:"display_name"`
	Address     Addresss `json:"address"`
	BoundingBox []string `json:"boundingbox"`
}

type Addresss struct {
	Road         string `json:"road"`
	CityBlock    string `json:"city_block"`
	Village      string `json:"village"`
	Suburb       string `json:"suburb"`
	CityDistrict string `json:"city_district"`
	City         string `json:"city"`
	Postcode     string `json:"postcode"`
	Country      string `json:"country"`
	CountryCode  string `json:"country_code"`
}

func GPSToAddress(sURL, sLat, sLong string) (ReverseData, error) {

	url := sURL + "/reverse?format=jsonv2&lat=" + sLat + "&lon=" + sLong + "&zoom=16"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ReverseData{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ReverseData{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ReverseData{}, err
	}

	// convert body to struct ReverseData
	var data ReverseData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return ReverseData{}, err
	}
	return data, nil
}
