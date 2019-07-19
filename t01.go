package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	influxdb "github.com/influxdata/influxdb-client-go"
)

func main() {

	dat, err := ioutil.ReadFile("/Users/ma/.influxdbv2/credentials")
	check(err)

	token := string(dat)

	influx, err := influxdb.New("http://127.0.0.1:9999/api/v2", token)
	if err != nil {
		log.Fatal(err)
	}

	err = influx.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	err = influx.Write(context.Background(), "rick", "ag",
		&influxdb.RowMetric{
			NameStr: "michele",
			Tags: []*influxdb.Tag{
				{Key: "ktest1", Value: "k-test1"},
				{Key: "ktest2,k-test3", Value: "k-test2, k-test3"}},
			Fields: []*influxdb.Field{
				{Key: "ftest1", Value: 3},
				{Key: "ftest2", Value: "kftest2"}},
			TS: now.Add(-time.Minute)},
	)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)
	r, err := influx.QueryCSV(
		context.Background(),
		`from(bucket:bucket)|>range(start:-1000h)|>group()`,
		`ag`,
		struct {
			Bucket string `flux:"bucket"`
		}{Bucket: "rick"})
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
