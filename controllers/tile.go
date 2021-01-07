package controllers

import (
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

type Coord struct {
	x int
	y int
	z int
}

func GetTile(w http.ResponseWriter, r *http.Request){
	c := quadToCoord(chi.URLParam(r, "coord"))

	url := "https://api.mapbox.com/v4/mapbox.satellite/" + strconv.Itoa(c.z) + "/" + strconv.Itoa(c.x) + "/" +strconv.Itoa(c.y)+"@2x.png256?access_token="+os.Getenv("MAPBOX_TOKEN")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(respBody)
	return
}

func quadToCoord (key string) Coord {
	coord := Coord{
		x: 0,
		y: 0,
		z: len(key),
	}
	keyInt, _ := strconv.Atoi(key)
	for i := len(key); i > 0; i--{
		mask := 1 << (i - 1)
		q := digitAt(keyInt,i)

		if q == 1 {
			coord.x |= mask
		}
		if q == 2 {
			coord.y |= mask
		}
		if q == 3 {
			coord.x |= mask
			coord.y |= mask
		}
	}

	return coord
}

func digitAt(num, pos int) int {
	r := num % int(math.Pow(10, float64(pos)))
	return r / int(math.Pow(10,float64(pos-1)))
}