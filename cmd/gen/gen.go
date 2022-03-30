package main

import (
	"context"
	"flag"
	"log"
	"math"
	"os"
	"path/filepath"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/AlekSi/test_db/cmd/internal/mongodb"
)

const (
	mongoURI = "mongodb://localhost/values"
)

func main() {
	flag.Parse()

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("values").Collection("values")
	collection.Drop(ctx)

	// in order of element types from https://bsonspec.org/spec.html
	data := []bson.D{
		{{"_id", 0x0101}, {"name", "double"}, {"value", 42.13}},
		{{"_id", 0x0102}, {"name", "double-zero"}, {"value", 0.0}},
		{{"_id", 0x0103}, {"name", "double-max"}, {"value", math.MaxFloat64}},
		{{"_id", 0x0104}, {"name", "double-smallest"}, {"value", math.SmallestNonzeroFloat64}},
		{{"_id", 0x0105}, {"name", "double-positive-infinity"}, {"value", math.Inf(+1)}},
		{{"_id", 0x0106}, {"name", "double-negative-infinity"}, {"value", math.Inf(-1)}},
		{{"_id", 0x0107}, {"name", "double-nan"}, {"value", math.NaN()}},

		{{"_id", 0x0201}, {"name", "string"}, {"value", "foo"}},
		{{"_id", 0x0202}, {"name", "string-empty"}, {"value", ""}},
		{{"_id", 0x0203}, {"name", "string-shorter"}, {"value", "z"}},
		{{"_id", 0x0204}, {"name", "string-longer"}, {"value", "abcdefghijklmnopqrstuvwxyz"}},
		//{{"_id", mongodb.NewObjectID(0x0205)}, {"name", "string-nul"},{"value", "\x00"}},

		{{"_id", 0x0301}, {"name", "document"}, {"value", map[string]any{"document": 42}}},
		{{"_id", 0x0302}, {"name", "document-empty"}, {"value", map[string]any{}}},
		{{"_id", 0x0303}, {"name", "document-two"}, {"value", map[string]any{"document": 42.13, "foo": "bar"}}},
		{{"_id", 0x0304}, {"name", "document-three"}, {"value", map[string]any{"document": int32(0), "baz": nil}}},

		{{"_id", 0x0401}, {"name", "array"}, {"value", []any{"array", 42}}},
		{{"_id", 0x0402}, {"name", "array-empty"}, {"value", []any{}}},
		{{"_id", 0x0403}, {"name", "array-one"}, {"value", []any{42.13}}},
		{{"_id", 0x0404}, {"name", "array-three"}, {"value", []any{42, "foo", nil}}},
		{{"_id", 0x0405}, {"name", "array-embedded"}, {"value", []any{
			map[string]any{"document": "abc", "score": 42.13, "age": 1000},
			map[string]any{"document": "def", "score": 42.13, "age": 1000},
			map[string]any{"document": "jkl", "score": 24, "age": 1002},
		}}},

		{{"_id", 0x0501}, {"name", "binary"}, {"value", primitive.Binary{Subtype: 0x80, Data: []byte{42, 0, 13}}}},
		{{"_id", 0x0502}, {"name", "binary-empty"}, {"value", primitive.Binary{}}},
		{{"_id", 0x0503}, {"name", "binary-big"}, {"value", primitive.Binary{Subtype: 0, Data: []byte{0, 0, 128}}}},

		// no Undefined
		{{"_id", 0x0801}, {"name", "bool-false"}, {"value", false}},
		{{"_id", 0x0802}, {"name", "bool-true"}, {"value", true}},

		{{"_id", 0x901}, {"name", "datetime"}, {"value", time.Date(2021, 11, 1, 10, 18, 42, 123000000, time.UTC)}},
		{{"_id", 0x902}, {"name", "datetime-epoch"}, {"value", time.Unix(0, 0)}},
		{{"_id", 0x903}, {"name", "datetime-year-min"}, {"value", time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)}},
		{{"_id", 0x904}, {"name", "datetime-year-max"}, {"value", time.Date(9999, 12, 31, 23, 59, 59, 999000000, time.UTC)}},

		{{"_id", 0x0a01}, {"name", "null"}, {"value", nil}},

		{{"_id", 0x0b01}, {"name", "regex"}, {"value", primitive.Regex{Pattern: "foo", Options: "i"}}},
		{{"_id", 0x0b02}, {"name", "regex-empty"}, {"value", primitive.Regex{}}},

		// no DBPointer
		// no JavaScript code
		// no Symbol
		// no JavaScript code w/ scope

		{{"_id", 0x1001}, {"name", "int32"}, {"value", int32(42)}},
		{{"_id", 0x1002}, {"name", "int32-zero"}, {"value", int32(0)}},
		{{"_id", 0x1003}, {"name", "int32-max"}, {"value", int32(math.MaxInt32)}},
		{{"_id", 0x1004}, {"name", "int32-min"}, {"value", int32(math.MinInt32)}},

		{{"_id", 0x1101}, {"name", "timestamp"}, {"value", primitive.Timestamp{T: 42, I: 13}}},
		{{"_id", 0x1102}, {"name", "timestamp-i"}, {"value", primitive.Timestamp{I: 1}}},

		{{"_id", 0x1201}, {"name", "int64"}, {"value", int64(42)}},
		{{"_id", 0x1202}, {"name", "int64-zero"}, {"value", int64(0)}},
		{{"_id", 0x1203}, {"name", "int64-max"}, {"value", int64(math.MaxInt64)}},
		{{"_id", 0x1204}, {"name", "int64-min"}, {"value", int64(math.MinInt64)}},

		// no 128-bit decimal floating point (yet)

		// no Min key
		// no Max key
	}

	names := make(map[string]struct{}, len(data))
	for _, record := range data {
		name := (record[1].Value).(string)
		if _, ok := names[name]; ok {
			log.Fatalf("duplicate name: %s", record[1].Key)
		}
		names[name] = struct{}{}

		key := uint32(record[0].Value.(int))

		mongodb.SetObjectIDCounter(key - 1)

		record[0].Value = mongodb.NewObjectID(key)

		if _, err = collection.InsertOne(ctx, record); err != nil {
			log.Fatal(err)
		}
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(filepath.Join(pwd, "..", "mongodb", "values", "values.json"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := mongodb.Export(mongoURI, "values", f, 0); err != nil {
		log.Fatal(err)
	}
}
