// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package pagila

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type countryTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *countryTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("country").
func (v *countryTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *countryTableType) Columns() []string {
	return []string{
		"country_id",
		"country",
		"last_update",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *countryTableType) NewStruct() reform.Struct {
	return new(Country)
}

// NewRecord makes a new record for that table.
func (v *countryTableType) NewRecord() reform.Record {
	return new(Country)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *countryTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// CountryTable represents country view or table in SQL database.
var CountryTable = &countryTableType{
	s: parse.StructInfo{
		Type:    "Country",
		SQLName: "country",
		Fields: []parse.FieldInfo{
			{Name: "CountryID", Type: "int32", Column: "country_id"},
			{Name: "Country", Type: "string", Column: "country"},
			{Name: "LastUpdate", Type: "time.Time", Column: "last_update"},
		},
		PKFieldIndex: 0,
	},
	z: new(Country).Values(),
}

// String returns a string representation of this struct or record.
func (s Country) String() string {
	res := make([]string, 3)
	res[0] = "CountryID: " + reform.Inspect(s.CountryID, true)
	res[1] = "Country: " + reform.Inspect(s.Country, true)
	res[2] = "LastUpdate: " + reform.Inspect(s.LastUpdate, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Country) Values() []interface{} {
	return []interface{}{
		s.CountryID,
		s.Country,
		s.LastUpdate,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Country) Pointers() []interface{} {
	return []interface{}{
		&s.CountryID,
		&s.Country,
		&s.LastUpdate,
	}
}

// View returns View object for that struct.
func (s *Country) View() reform.View {
	return CountryTable
}

// Table returns Table object for that record.
func (s *Country) Table() reform.Table {
	return CountryTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Country) PKValue() interface{} {
	return s.CountryID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Country) PKPointer() interface{} {
	return &s.CountryID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Country) HasPK() bool {
	return s.CountryID != CountryTable.z[CountryTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.CountryID = pk.
func (s *Country) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = CountryTable
	_ reform.Struct = (*Country)(nil)
	_ reform.Table  = CountryTable
	_ reform.Record = (*Country)(nil)
	_ fmt.Stringer  = (*Country)(nil)
)

func init() {
	parse.AssertUpToDate(&CountryTable.s, new(Country))
}
