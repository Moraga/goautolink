package main

import (
	"context"
	"log"
	"sort"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type tmpContent struct {
	Title string
	URL   string
	Stem  []string
}

type Content struct {
	Title string
	URL   string
	Stem  []string
	Ustem []string
}

var db []Content

var vectors map[string]float64

func loadDB() {
	ctx, err1 := context.WithTimeout(context.Background(), 30*time.Second)
	if err1 != nil {
		// log.Fatal(err1)
	}
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("graph").Collection("content")

	// options := options.FindOptions{}
	// limit := int64(10)
	// options.Limit = &limit

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var r tmpContent
		err := cur.Decode(&r)
		if err != nil {
			log.Fatal(err)
		}
		uniqueStems := removeSliceDupes(r.Stem)
		db = append(db, Content{r.Title, r.URL, r.Stem, uniqueStems})
	}
}

func loadVectors() {
	vecs := make(map[string]int)
	for _, row := range db {
		for _, s := range row.Stem {
			if _, exists := vecs[s]; exists {
				vecs[s]++
			} else {
				vecs[s] = 1
			}
		}
	}

	delete(vecs, "")

	var list []int

	for _, a := range vecs {
		found := false
		for _, b := range list {
			if a == b {
				found = true
				break
			}
		}
		if !found {
			list = append(list, a)
		}
	}

	sort.Ints(list)

	var sets [][]int

	for list != nil {
		groups := clusterizeScalar2(list...)
		size := len(groups)
		list = groups[0]
		if size > 1 {
			sets = append(groups[1:], sets...)
		} else {
			sets = append([][]int{list}, sets...)
			list = nil
		}
	}

	total := float64(len(sets))

	kv := make(map[int]float64)

	for n, set := range sets {
		weight := (total - float64(n)) / total
		for _, k := range set {
			kv[k] = weight
		}
	}

	vecs2 := make(map[string]float64)

	for vec, k := range vecs {
		vecs2[vec] = kv[k]
	}

	vectors = vecs2
}

func insert(s string) {
	stem, _ := stemText(s)
	uniqueStems := removeSliceDupes(stem)
	db = append(db, Content{s, "", stem, uniqueStems})
}
