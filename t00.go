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

	influx, err := influxdb.New("http://127.0.0.1:9999/api/v2", "7pOQF90FzHQPFQcu77gRV9SUDWXaB4wqU_8K_zA0lAQg7GxyE5JLHV9brZrAae3drj1T64dFjC3iqfL5eQ_tLA==")
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
			NameStr: "test",
			Tags: []*influxdb.Tag{
				{Key: "ktest1", Value: "k-test1"},
				{Key: "ktest2,k-test3", Value: "k-test2, k-test3"}},
			Fields: []*influxdb.Field{
				{Key: "ftest1", Value: 3},
				{Key: "ftest2", Value: "kftest2"}},
			TS: now.Add(-time.Minute)},
		&influxdb.RowMetric{
			NameStr: "test",
			Tags: []*influxdb.Tag{
				{Key: "ktest1", Value: "k-test1"},
				{Key: "ktest2", Value: "k-test2"},
				{Key: "ktest3", Value: "k-test3"}},
			Fields: []*influxdb.Field{
				{Key: "ftest1", Value: 3},
				{Key: "ftest2", Value: "kftest2"}},
			TS: now.Add(-time.Minute)},
		&influxdb.RowMetric{
			NameStr: "test",
			Tags: []*influxdb.Tag{
				{Key: "ktest1", Value: "k-test1"},
				{Key: "ktest2", Value: "k-test2"}},
			Fields: []*influxdb.Field{
				{Key: "ftest1", Value: 3},
				{Key: "ftest2", Value: "kftest2"}},
			TS: now.Add(-time.Minute)},
		&influxdb.RowMetric{
			NameStr: "test",
			Tags: []*influxdb.Tag{
				{Key: "ktest1", Value: "k-test1"},
				{Key: "ktest3", Value: "k-test3"}},
			Fields: []*influxdb.Field{
				{Key: "ftest1", Value: 4},
				{Key: "ftest2", Value: "kftest2"}},
			TS: now.Add(-time.Minute)},
		&influxdb.RowMetric{
			NameStr: "test",
			Tags: []*influxdb.Tag{
				{Key: "ktest1", Value: "k-test1"},
				{Key: "ktest3", Value: "k-test3"}},
			Fields: []*influxdb.Field{
				{Key: "ftest1", Value: 3},
				{Key: "ftest2", Value: "kftest2"}},
			TS: now.Add(-time.Second * 30)},
		&influxdb.RowMetric{
			NameStr: "test",
			Tags: []*influxdb.Tag{
				{Key: "ktest1", Value: "k-test1"},
				{Key: "ktest3", Value: "k-test3"}},
			Fields: []*influxdb.Field{
				{Key: "ftest1", Value: 6},
				{Key: "ftest2", Value: "kftest3"}},
			TS: now.Add(-time.Second * 25)},
		&influxdb.RowMetric{
			NameStr: "tes0",
			Tags: []*influxdb.Tag{
				{Key: "ktest1", Value: "k-test1"},
				{Key: "ktest3", Value: "k-test3"}},
			Fields: []*influxdb.Field{
				{Key: "ftest1", Value: 5},
				{Key: "ftest2", Value: "kftest3"}},
			TS: now.Add(-time.Second * 25).Truncate(time.Microsecond).Add(1)},
		&influxdb.RowMetric{
			NameStr: "tes0",
			Tags: []*influxdb.Tag{
				{Key: "ktest1", Value: "k-test1"},
				{Key: "ktest3", Value: "k-test3"}},
			Fields: []*influxdb.Field{
				{Key: "ftest1", Value: 5},
				{Key: "ftest2", Value: "kftest3"}},
			TS: now.Add(-time.Second * 25).Truncate(time.Microsecond).Add(time.Microsecond)},
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
