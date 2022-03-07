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
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

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
	data := map[uint32]struct {
		name string
		v    any
	}{
		0x0101: {name: "double", v: 42.13},
		0x0102: {name: "double-zero", v: 0.0},
		0x0103: {name: "double-max", v: math.MaxFloat64},
		0x0104: {name: "double-smallest", v: math.SmallestNonzeroFloat64},
		0x0105: {name: "double-positive-infinity", v: math.Inf(+1)},
		0x0106: {name: "double-negative-infinity", v: math.Inf(-1)},
		0x0107: {name: "double-nan", v: math.NaN()},

		0x0201: {name: "string", v: "foo"},
		0x0202: {name: "string-empty", v: ""},
		0x0203: {name: "string-shorter", v: "z"},
		0x0204: {name: "string-longer", v: "abcdefghijklmnopqrstuvwxyz"},
		// 0x0205: {name: "string-nul", v: "\x00"},

		0x0301: {name: "document", v: map[string]any{"document": 42}},
		0x0302: {name: "document-empty", v: map[string]any{}},
		0x0303: {name: "document-two", v: map[string]any{"document": 42.13, "foo": "bar"}},
		0x0304: {name: "document-three", v: map[string]any{"document": int32(0), "baz": nil}},

		0x0401: {name: "array", v: []any{"array", 42}},
		0x0402: {name: "array-empty", v: []any{}},
		0x0403: {name: "array-one", v: []any{42.13}},
		0x0404: {name: "array-three", v: []any{42, "foo", nil}},
		0x0405: {name: "array-embedded", v: []any{
			map[string]any{"document": "abc", "score": 42.13, "age": 999},
			map[string]any{"document": "def", "score": 42.13, "age": 1000},
			map[string]any{"document": "def", "score": 24, "age": 1002},
		}},

		0x0501: {name: "binary", v: primitive.Binary{Subtype: 0x80, Data: []byte{42, 0, 13}}},
		0x0502: {name: "binary-empty", v: primitive.Binary{}},

		// no Undefined

		0x0801: {name: "bool-false", v: false},
		0x0802: {name: "bool-true", v: true},

		0x901: {name: "datetime", v: time.Date(2021, 11, 1, 10, 18, 42, 123000000, time.UTC)},
		0x902: {name: "datetime-epoch", v: time.Unix(0, 0)},
		0x903: {name: "datetime-year-min", v: time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)},
		0x904: {name: "datetime-year-max", v: time.Date(9999, 12, 31, 23, 59, 59, 999000000, time.UTC)},

		0x0a01: {name: "null", v: nil},

		0x0b01: {name: "regex", v: primitive.Regex{Pattern: "foo", Options: "i"}},
		0x0b02: {name: "regex-empty", v: primitive.Regex{}},

		// no DBPointer
		// no JavaScript code
		// no Symbol
		// no JavaScript code w/ scope

		0x1001: {name: "int32", v: int32(42)},
		0x1002: {name: "int32-zero", v: int32(0)},
		0x1003: {name: "int32-max", v: int32(math.MaxInt32)},
		0x1004: {name: "int32-min", v: int32(math.MinInt32)},

		0x1101: {name: "timestamp", v: primitive.Timestamp{T: 42, I: 13}},
		0x1102: {name: "timestamp-i", v: primitive.Timestamp{I: 1}},

		0x1201: {name: "int64", v: int64(42)},
		0x1202: {name: "int64-zero", v: int64(0)},
		0x1203: {name: "int64-max", v: int64(math.MaxInt64)},
		0x1204: {name: "int64-min", v: int64(math.MinInt64)},

		// no 128-bit decimal floating point (yet)

		// no Min key
		// no Max key
	}

	keys := maps.Keys(data)
	slices.Sort(keys)

	names := make(map[string]struct{}, len(keys))
	for _, key := range keys {
		value := data[key]
		if _, ok := names[value.name]; ok {
			log.Fatalf("duplicate name: %s", value.name)
		}
		names[value.name] = struct{}{}

		// to keep IDs stable if we insert new documents in the middle
		mongodb.SetObjectIDCounter(key - 1)

		doc := bson.D{{"_id", mongodb.NewObjectID(key)}, {"name", value.name}, {"value", value.v}}
		if _, err = collection.InsertOne(ctx, doc); err != nil {
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
