// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package sakila

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type filmTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *filmTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("film").
func (v *filmTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *filmTableType) Columns() []string {
	return []string{
		"film_id",
		"title",
		"description",
		"release_year",
		"language_id",
		"original_language_id",
		"rental_duration",
		"rental_rate",
		"length",
		"replacement_cost",
		"rating",
		"special_features",
		"last_update",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *filmTableType) NewStruct() reform.Struct {
	return new(Film)
}

// NewRecord makes a new record for that table.
func (v *filmTableType) NewRecord() reform.Record {
	return new(Film)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *filmTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// FilmTable represents film view or table in SQL database.
var FilmTable = &filmTableType{
	s: parse.StructInfo{
		Type:    "Film",
		SQLName: "film",
		Fields: []parse.FieldInfo{
			{Name: "FilmID", Type: "int16", Column: "film_id"},
			{Name: "Title", Type: "string", Column: "title"},
			{Name: "Description", Type: "*string", Column: "description"},
			{Name: "ReleaseYear", Type: "*uint16", Column: "release_year"},
			{Name: "LanguageID", Type: "int8", Column: "language_id"},
			{Name: "OriginalLanguageID", Type: "*int8", Column: "original_language_id"},
			{Name: "RentalDuration", Type: "int8", Column: "rental_duration"},
			{Name: "RentalRate", Type: "string", Column: "rental_rate"},
			{Name: "Length", Type: "*int16", Column: "length"},
			{Name: "ReplacementCost", Type: "string", Column: "replacement_cost"},
			{Name: "Rating", Type: "[]uint8", Column: "rating"},
			{Name: "SpecialFeatures", Type: "[]uint8", Column: "special_features"},
			{Name: "LastUpdate", Type: "time.Time", Column: "last_update"},
		},
		PKFieldIndex: 0,
	},
	z: new(Film).Values(),
}

// String returns a string representation of this struct or record.
func (s Film) String() string {
	res := make([]string, 13)
	res[0] = "FilmID: " + reform.Inspect(s.FilmID, true)
	res[1] = "Title: " + reform.Inspect(s.Title, true)
	res[2] = "Description: " + reform.Inspect(s.Description, true)
	res[3] = "ReleaseYear: " + reform.Inspect(s.ReleaseYear, true)
	res[4] = "LanguageID: " + reform.Inspect(s.LanguageID, true)
	res[5] = "OriginalLanguageID: " + reform.Inspect(s.OriginalLanguageID, true)
	res[6] = "RentalDuration: " + reform.Inspect(s.RentalDuration, true)
	res[7] = "RentalRate: " + reform.Inspect(s.RentalRate, true)
	res[8] = "Length: " + reform.Inspect(s.Length, true)
	res[9] = "ReplacementCost: " + reform.Inspect(s.ReplacementCost, true)
	res[10] = "Rating: " + reform.Inspect(s.Rating, true)
	res[11] = "SpecialFeatures: " + reform.Inspect(s.SpecialFeatures, true)
	res[12] = "LastUpdate: " + reform.Inspect(s.LastUpdate, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Film) Values() []interface{} {
	return []interface{}{
		s.FilmID,
		s.Title,
		s.Description,
		s.ReleaseYear,
		s.LanguageID,
		s.OriginalLanguageID,
		s.RentalDuration,
		s.RentalRate,
		s.Length,
		s.ReplacementCost,
		s.Rating,
		s.SpecialFeatures,
		s.LastUpdate,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Film) Pointers() []interface{} {
	return []interface{}{
		&s.FilmID,
		&s.Title,
		&s.Description,
		&s.ReleaseYear,
		&s.LanguageID,
		&s.OriginalLanguageID,
		&s.RentalDuration,
		&s.RentalRate,
		&s.Length,
		&s.ReplacementCost,
		&s.Rating,
		&s.SpecialFeatures,
		&s.LastUpdate,
	}
}

// View returns View object for that struct.
func (s *Film) View() reform.View {
	return FilmTable
}

// Table returns Table object for that record.
func (s *Film) Table() reform.Table {
	return FilmTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Film) PKValue() interface{} {
	return s.FilmID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Film) PKPointer() interface{} {
	return &s.FilmID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Film) HasPK() bool {
	return s.FilmID != FilmTable.z[FilmTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.FilmID = pk.
func (s *Film) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = FilmTable
	_ reform.Struct = (*Film)(nil)
	_ reform.Table  = FilmTable
	_ reform.Record = (*Film)(nil)
	_ fmt.Stringer  = (*Film)(nil)
)

func init() {
	parse.AssertUpToDate(&FilmTable.s, new(Film))
}
