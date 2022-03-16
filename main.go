package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/h2non/bimg"
)

const envPrefix = "IMGIN_"
const envPort = envPrefix + "PORT"
const envOrigin = envPrefix + "ORIGIN"
const envSecret = envPrefix + "SECRET"

func main() {
	http.HandleFunc("/", handle)

	p := os.Getenv(envPort)
	if p == "" {
		p = "8080"
	}
	addr := ":" + p
	fmt.Println("listening on", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	o := os.Getenv(envOrigin)
	if o == "" {
		o = "*"
	}
	w.Header().Add("Access-Control-Allow-Origin", o)
	w.Header().Add("Access-Control-Allow-Methods", "POST")

	if r.Method == "OPTIONS" {
		return
	}

	s := os.Getenv(envSecret)
	if s != r.Header.Get("authorization") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	width := r.URL.Query().Get("width")
	height := r.URL.Query().Get("height")
	format := r.URL.Query().Get("format")

	var tW, tH int
	var errW, errH error

	if width != "" {
		tW, errW = strconv.Atoi(width)
	}
	if height != "" {
		tH, errH = strconv.Atoi(height)
	}
	if errW != nil || errH != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	i := bimg.NewImage(buf.Bytes())

	var t bimg.ImageType
	switch format {
	case "webp":
		t = bimg.WEBP
	case "avif":
		t = bimg.AVIF
	default:
		t = bimg.JPEG
	}

	res, err := i.Process(bimg.Options{
		Width:  tW,
		Height: tH,
		Type:   t,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
}
