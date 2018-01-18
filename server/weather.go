package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"
)

// Start, End format: 2017-04-01T00:00:00Z
type DatumQuery struct {
	Site      string
	DatumName string
	Start     string
	End       string
}

type Value struct {
	Val float32 `json:"value"`
}

type Bucket struct {
	KeyAsString string  `json:"key_as_string"`
	Key         float32 `json:"key"`
	DocCount    float32 `json:"doc_count"`
	Value       Value   `json:"avg(value_float)"`
}

type Aggregation struct {
	Buckets []Bucket `json:"buckets"`
}

type Aggregations struct {
	Aggregation Aggregation `json:"date(timestampmeasured30m)"`
}

type EsResponse struct {
	Aggregations Aggregations `json:"aggregations"`
}

type Shit struct {
	Took float32 `json:"took"`
}

type Datum struct {
	TimeStamp string
	Value     float32
}

func (esResponse *EsResponse) toDatums() []Datum {
	var datums []Datum
	for _, bucket := range esResponse.Aggregations.Aggregation.Buckets {
		datums = append(datums, Datum{bucket.KeyAsString, bucket.Value.Val})
	}
	return datums
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Please see the documentation for the weather API")
}

func ParamsToEsQuery(site string, datumName string, start string, end string) ([]byte, error) {
	buf := new(bytes.Buffer)
	query := DatumQuery{site, datumName, start, end}
	tmpl, err := template.ParseFiles("searchtemplate.json")
	if err != nil {
		return nil, err
	}
	err = tmpl.Execute(buf, query)

	return buf.Bytes(), nil
}

func EsSearch(SearchString []byte) (*EsResponse, error) {
	esUrl := "http://elasticsearch.lco.gtn:9200/_search?pretty"
	res, err := http.Post(esUrl, "application/json", bytes.NewBuffer(SearchString))
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	var esr = new(EsResponse)
	// Golang doesn't like commas in keys
	newbody := string(body)
	replaced := strings.Replace(newbody, ",30m", "30m", -1)
	err = json.Unmarshal([]byte(replaced), &esr)
	if err != nil {
		return esr, err
	}
	return esr, nil

}

func Query(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	site := queryValues.Get("site")
	datumName := queryValues.Get("datumname")

	start := queryValues.Get("start")
	if start == "" {
		start_time := time.Now()
		start_time = start_time.Add(-time.Duration(36000) * time.Second)
		start = start_time.Format("2006-01-02T15:04:05Z")
	}
	log.Println(start)
	end := queryValues.Get("end")
	if end == "" {
		end = time.Now().Format("2006-01-02T15:04:05Z")
	}
	log.Println(end)
	parsed, err := ParamsToEsQuery(site, datumName, start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	esr, err := EsSearch(parsed)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	datums := esr.toDatums()

	b, err := json.Marshal(datums)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, string(b))
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/query", Query)
	log.Println("Starting server on http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
