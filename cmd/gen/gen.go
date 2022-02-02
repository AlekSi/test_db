package main

import (
	"context"
	"flag"
	"log"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/AlekSi/test_db/cmd/internal/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/maps"
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
	data := map[uint32]any{
		0x0101: 42.13,
		0x0102: 0.0,
		0x0103: math.MaxFloat64,
		0x0104: math.SmallestNonzeroFloat64,
		0x0105: math.Inf(1),
		0x0106: math.Inf(-1),
		0x0107: math.NaN(),

		0x0201: "foo",
		0x0202: "",
		0x0203: "\x00",

		0x0301: map[string]any{"document": 42},
		0x0302: map[string]any{},

		0x0401: []any{"array", 42},
		0x0402: []any{},

		0x0501: primitive.Binary{Subtype: 0x80, Data: []byte{42, 0, 13}},
		0x0502: primitive.Binary{},

		// no Undefined

		0x0801: false,
		0x0802: true,

		0x901: time.Date(2021, 11, 1, 10, 18, 42, 123000000, time.UTC),
		0x902: time.Unix(0, 0),
		0x903: time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
		0x904: time.Date(9999, 12, 31, 23, 59, 59, 999000000, time.UTC),

		0x0a01: nil,

		0x0b01: primitive.Regex{Pattern: "foo", Options: "i"},
		0x0b02: primitive.Regex{},

		// no DBPointer
		// no JavaScript code
		// no Symbol
		// no JavaScript code w/ scope

		0x1001: int32(42),
		0x1002: int32(0),
		0x1003: int32(math.MaxInt32),
		0x1004: int32(math.MinInt32),

		0x1101: primitive.Timestamp{T: 42, I: 13},
		0x1102: primitive.Timestamp{},

		0x1201: int64(42),
		0x1202: int64(0),
		0x1203: int64(math.MaxInt64),
		0x1204: int64(math.MinInt64),

		// no 128-bit decimal floating point (yet)

		// no Min key
		// no Max key
	}

	keys := maps.Keys(data)
	for _, key := range keys {
		doc := bson.D{{"_id", mongodb.NewObjectID(key)}, {"value", data[key]}}
		if _, err = collection.InsertOne(ctx, doc); err != nil {
			log.Fatal(err)
		}
	}

	f, err := os.Create(filepath.Join("..", "..", "mongodb", "values", "values.json"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := mongodb.Export(mongoURI, "values", f, 2); err != nil {
		log.Fatal(err)
	}
}
