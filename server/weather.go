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

type EsAggResponse struct {
	Aggregations Aggregations `json:"aggregations"`
}

type Source struct {
	TimeStampMeasured string  `json:"timestampmeasured"`
	ValueFloat        float32 `json:"value_float"`
	ValueString       string  `json:"value_string"`
}

type SubHit struct {
	Source Source `json:"_source"`
}

type Hit struct {
	SubHits []SubHit `json:"hits"`
}

type EsStdResponse struct {
	Hit Hit `json:"hits"`
}

type EsResponse interface {
	toDatums() []Datum
}
type Datum struct {
	TimeStamp   string
	Value       float32
	ValueString string
}

func (esStdResponse *EsStdResponse) toDatums() []Datum {
	var datums []Datum
	for _, subhit := range esStdResponse.Hit.SubHits {
		datums = append(datums, Datum{subhit.Source.TimeStampMeasured, subhit.Source.ValueFloat, subhit.Source.ValueString})
	}
	return datums
}

func (esAggResponse *EsAggResponse) toDatums() []Datum {
	var datums []Datum
	for _, bucket := range esAggResponse.Aggregations.Aggregation.Buckets {
		datums = append(datums, Datum{bucket.KeyAsString, bucket.Value.Val, ""})
	}
	return datums
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Please see the documentation for the weather API")
}

func ParamsToEsQuery(site string, datumName string, start string, end string, agg bool) ([]byte, error) {
	buf := new(bytes.Buffer)
	query := DatumQuery{site, datumName, start, end}

	template_str := ""
	if agg {
		template_str = "aggsearchtemplate.json"
	} else {
		template_str = "searchtemplate.json"
	}
	tmpl, err := template.ParseFiles(template_str)

	if err != nil {
		return nil, err
	}
	err = tmpl.Execute(buf, query)

	return buf.Bytes(), nil
}

func EsSearch(SearchString []byte, agg bool) (EsResponse, error) {
	esUrl := "http://elasticsearch.lco.gtn:9200/_search?pretty"
	res, err := http.Post(esUrl, "application/json", bytes.NewBuffer(SearchString))
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	var esr EsResponse
	// Golang doesn't like commas in keys
	newbody := string(body)
	replaced := strings.Replace(newbody, ",30m", "30m", -1)
	if agg {
		esr = new(EsAggResponse)
	} else {
		esr = new(EsStdResponse)
	}
	err = json.Unmarshal([]byte(replaced), esr)
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
		start_time := time.Now().UTC()
		start_time = start_time.Add(-time.Duration(3600*24) * time.Second)
		start = start_time.Format("2006-01-02T15:04:05Z")
	}
	log.Println(start)

	end := queryValues.Get("end")
	if end == "" {
		end = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	}
	log.Println(end)

	agg := queryValues.Get("agg")
	aggBool := false
	if agg == "" || agg == "true" {
		aggBool = true
	}

	parsed, err := ParamsToEsQuery(site, datumName, start, end, aggBool)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	esr, err := EsSearch(parsed, aggBool)

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
	log.Println("Starting server on http://127.0.0.1:80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
