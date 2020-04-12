// wheather Пакет погоды

package wheather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var WheatherNow Wheather

func init() {
	WheatherNow = New()
}

var postalCode string = "121552"
var country string = "RU"

// Wheather Погода
type Wheather struct {
	City                string  `json:"city_name"`
	Temperature         float64 `json:"temp"`
	TemperatureApparent float64 `json:"app_temp"`
	Pressure            float64 `json:"pres"`
	RelativeHumidity    float64 `json:"rh"`
}

// New wheather
func New() Wheather {
	return getWheather("https://api.weatherbit.io/v2.0/current?postal_code=121552&key=9eea2cdd9d39423bb8544131ae377e30")
}

// getWheather Возвращает объект погоды
func getWheather(url string) Wheather {
	body := getData(url)

	var wht Wheather
	json.Unmarshal(body, &wht)
	str := fmt.Sprint(string(body))
	fmt.Println(str)
	return wht
}

func getData(url string) []byte {

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body) // was "res", now "res.Body"
	if err != nil {
		panic(err)
	}
	// res.Close() // was absent
	return body

}