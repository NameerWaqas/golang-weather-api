package weather

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/mux"
)

func GetWeatherByCityName(res http.ResponseWriter, req *http.Request){
	// Path variable
	name := mux.Vars(req)["name"]

	baseUrl,_ := url.Parse("http://api.weatherapi.com")
	
	// Path params
	baseUrl.Path += "/v1/current.json"

	//Query params
	params := url.Values{}
	params.Add("key",os.Getenv("API_KEY"))
	params.Add("q",name)
	params.Add("aqi","no")
	baseUrl.RawQuery = params.Encode()

	// Fetching Data from weather API
	resp,err := http.Get(baseUrl.String())
	if err!=nil{
		panic(err)
	}
	
	bytes,_:=ioutil.ReadAll(resp.Body)

	json.NewEncoder(res).Encode(string(bytes))
}