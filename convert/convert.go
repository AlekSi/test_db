package main

import (
	"database/sql"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"

	_ "github.com/go-sql-driver/mysql" // register database/sql driver
	_ "github.com/jackc/pgx/v4/stdlib" // register database/sql driver
	_ "github.com/lib/pq"              // register database/sql driver
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/reform.v1"
	mysqldialect "gopkg.in/reform.v1/dialects/mysql"
	postgresqldialect "gopkg.in/reform.v1/dialects/postgresql"

	"github.com/AlekSi/test_db/convert/mongodb"
	"github.com/AlekSi/test_db/convert/mysql/sakila"
	mysqlWorld "github.com/AlekSi/test_db/convert/mysql/world"
	"github.com/AlekSi/test_db/convert/postgresql/pagila"
	postgresqlWorld "github.com/AlekSi/test_db/convert/postgresql/world"
)

const (
	mysqlURI    = "root@/%s?parseTime=true&clientFoundRows=true&time_zone='UTC'"
	postgresURI = "postgres://postgres@127.0.0.1:5432/%s?sslmode=disable"
	pgxURI      = "postgres://postgres@127.0.0.1:5432/%s"
)

func main() {
	verboseF := flag.Bool("verbose", false, "be verbose")
	fromF := flag.String("from", "postgres", "import from: postgres (lib/pq), pgx, mysql")
	databaseF := flag.String("database", "sakila", "database name: sakila (for both Sakila and Pagila), world")
	flag.Usage = func() {
		u := "Convert tool reads Sakila/Pagila or World dataset from MySQL/PostgreSQL,\n" +
			"adds ObjectID _id fields to rows, then converts them to ExtJSON and\n" +
			"imports them into MongoDB via mongoimport.\n" +
			"It then exports the resulting Monila or World database into ExtJSON files.\n\n" +
			"Please note that MySQL's Sakila and PostgreSQL's Pagila datasets\n" +
			"have small differences, most notably in dates.\n"
		fmt.Fprintln(flag.CommandLine.Output(), u)

		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	var reformLogger reform.Logger
	if *verboseF {
		reformLogger = reform.NewPrintfLogger(log.New(os.Stderr, "reform: ", 0).Printf)
	}

	var views []reform.View
	switch *databaseF {
	case "sakila":
		if *fromF == "mysql" {
			views = []reform.View{
				// sakila.ActorInfoView, // view
				sakila.ActorTable,
				sakila.AddressTable,
				sakila.CategoryTable,
				sakila.CityTable,
				sakila.CountryTable,
				// sakila.CustomerListView, // view
				sakila.CustomerTable,
				sakila.FilmActorView,
				sakila.FilmCategoryView,
				// sakila.FilmListView, // view
				sakila.FilmTable,
				sakila.FilmTextTable,
				sakila.InventoryTable,
				sakila.LanguageTable,
				// sakila.NicerButSlowerFilmListView, // view
				sakila.PaymentTable,
				sakila.RentalTable,
				// sakila.SalesByFilmCategoryView, // view
				// sakila.SalesByStoreView,        // view
				// sakila.StaffListView,           // view
				sakila.StaffTable,
				sakila.StoreTable,
			}
		} else {
			views = []reform.View{
				// pagila.ActorInfoView, // view
				pagila.ActorTable,
				pagila.AddressTable,
				pagila.CategoryTable,
				pagila.CityTable,
				pagila.CountryTable,
				// pagila.CustomerListView, // view
				pagila.CustomerTable,
				pagila.FilmActorView,
				pagila.FilmCategoryView,
				// pagila.FilmListView, // view
				pagila.FilmTable,
				// pagila.FilmTextTable, // missing
				pagila.InventoryTable,
				pagila.LanguageTable,
				// pagila.NicerButSlowerFilmListView, // view
				// pagila.PaymentTable,               // missing
				pagila.RentalTable,
				// pagila.SalesByFilmCategoryView, // view
				// pagila.SalesByStoreView,        // view
				// pagila.StaffListView,           // view
				pagila.StaffTable,
				pagila.StoreTable,
			}
		}
	case "world":
		if *fromF == "mysql" {
			views = []reform.View{
				mysqlWorld.CityTable,
				mysqlWorld.CountryTable,
			}
		} else {
			views = []reform.View{
				postgresqlWorld.CityTable,
				postgresqlWorld.CountryTable,
			}
		}
	default:
		log.Fatalf("Unhandled -database value: %q.", *databaseF)
	}

	var reformDB *reform.DB
	sqlDatabaseName := "world"
	mongoDatabaseName := "world"

	switch *fromF {
	case "mysql":
		if *databaseF == "sakila" {
			sqlDatabaseName = "sakila"
			mongoDatabaseName = "monila"
		}

		db, err := sql.Open("mysql", fmt.Sprintf(mysqlURI, sqlDatabaseName))
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		reformDB = reform.NewDB(db, mysqldialect.Dialect, reformLogger)

	case "postgres":
		if *databaseF == "sakila" {
			sqlDatabaseName = "pagila"
			mongoDatabaseName = "monila"
		}

		db, err := sql.Open("postgres", fmt.Sprintf(postgresURI, sqlDatabaseName))
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		reformDB = reform.NewDB(db, postgresqldialect.Dialect, reformLogger)

	case "pgx":
		if *databaseF == "sakila" {
			sqlDatabaseName = "pagila"
			mongoDatabaseName = "monila"
		}

		db, err := sql.Open("pgx", fmt.Sprintf(pgxURI, sqlDatabaseName))
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		reformDB = reform.NewDB(db, postgresqldialect.Dialect, reformLogger)

	default:
		log.Fatalf("Unhandled -from value: %q.", *fromF)
	}

	mongoURI := "mongodb://localhost/" + mongoDatabaseName

	script := []string{
		"#!/bin/bash\n",
		"# Generated by convert.go. DO NOT EDIT.\n",

		// docker-entrypoint.sh sources the script with the current directory set to /
		"DIR=$(dirname $BASH_SOURCE)\n",
	}

	for _, view := range views {
		viewName := view.Name()
		log.Printf("%s ...", viewName)

		importView(reformDB, view, mongoURI, *verboseF)
		exportView(mongoURI, mongoDatabaseName, viewName, *verboseF)

		cmd := fmt.Sprintf(
			"mongoimport --uri=%s --collection=%s --drop --maintainInsertionOrder $DIR/%s",
			mongoURI, viewName, viewName+".json",
		)
		script = append(script, cmd)
	}

	fn := filepath.Join("..", "mongodb", mongoDatabaseName, "import.sh")
	log.Printf("Writing %s ...", fn)
	if err := os.WriteFile(fn, []byte(strings.Join(script, "\n")+"\n"), 0o755); err != nil {
		log.Fatal(err)
	}
}

// importView imports SQL view into MongoDB.
func importView(db *reform.DB, view reform.View, mongoURI string, verbose bool) {
	rows, err := db.SelectRows(view, "")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	r, w := io.Pipe()
	done := make(chan error)
	go func() {
		verbosity := -1
		if verbose {
			verbosity = 1
		}

		done <- mongodb.Import(mongoURI, view.Name(), r, verbosity)
	}()

	for {
		str := view.NewStruct()
		if err = db.NextRow(str, rows); err != nil {
			break
		}

		var id interface{}
		if r, ok := str.(reform.Record); ok {
			switch v := r.PKValue(); v := v.(type) {
			case int8:
				id = newObjectID(uint32(v))
			case int16:
				id = newObjectID(uint32(v))
			case int32:
				id = newObjectID(uint32(v))
			case string:
				id = v
			default:
				log.Fatalf("Unhandled PK type %T for %s.", v, str)
			}
		}

		cols := view.Columns()
		d := make(bson.D, 0, len(cols)+1)
		d = append(d, primitive.E{Key: "_id", Value: id})
		for i, val := range str.Values() {
			d = append(d, primitive.E{Key: cols[i], Value: val})
		}

		b, err := bson.MarshalExtJSON(d, true, true)
		if err != nil {
			log.Fatal(err)
		}

		if verbose {
			log.Printf("%s", b)
		}

		w.Write(b)
	}
	if err != reform.ErrNoRows {
		log.Fatal(err)
	}

	if err := w.Close(); err != nil {
		log.Fatal(err)
	}

	err = <-done
	if err != nil {
		log.Fatal(err)
	}
}

// exportView exports MongoDB collection into JSON file.
func exportView(mongoURI, database, collection string, verbose bool) {
	f, err := os.Create(filepath.Join("..", "mongodb", database, collection+".json"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	verbosity := -1
	if verbose {
		verbosity = 1
	}

	if err := mongodb.Export(mongoURI, collection, f, verbosity); err != nil {
		log.Fatal(err)
	}
}

var (
	ts              = uint32(time.Date(2021, 9, 1, 0, 0, 0, 0, time.UTC).Unix())
	objectIDCounter uint32
)

// newObjectID generates stable BSON ObjectID to make conversion results stable.
func newObjectID(id uint32) primitive.ObjectID {
	processUnique := make([]byte, 5)
	binary.BigEndian.PutUint32(processUnique, id)

	var b [12]byte

	binary.BigEndian.PutUint32(b[0:4], ts)
	copy(b[4:9], processUnique)

	v := atomic.AddUint32(&objectIDCounter, 1)
	b[9] = byte(v >> 16)
	b[10] = byte(v >> 8)
	b[11] = byte(v)

	return b
}
