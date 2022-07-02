// Code generated by entc, DO NOT EDIT.

package store

import (
	"fmt"
	"news/internal/store/tvsource"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// TvSource is the model entity for the TvSource schema.
type TvSource struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Logo holds the value of the "logo" field.
	Logo string `json:"logo,omitempty"`
	// Video holds the value of the "video" field.
	Video string `json:"video,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Status holds the value of the "status" field.
	Status bool `json:"status,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Language holds the value of the "language" field.
	Language string `json:"language,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TvSource) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case tvsource.FieldStatus:
			values[i] = new(sql.NullBool)
		case tvsource.FieldID:
			values[i] = new(sql.NullInt64)
		case tvsource.FieldLogo, tvsource.FieldVideo, tvsource.FieldTitle, tvsource.FieldCountry, tvsource.FieldDescription, tvsource.FieldLanguage:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TvSource", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TvSource fields.
func (ts *TvSource) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tvsource.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ts.ID = int(value.Int64)
		case tvsource.FieldLogo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo", values[i])
			} else if value.Valid {
				ts.Logo = value.String
			}
		case tvsource.FieldVideo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video", values[i])
			} else if value.Valid {
				ts.Video = value.String
			}
		case tvsource.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				ts.Title = value.String
			}
		case tvsource.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ts.Status = value.Bool
			}
		case tvsource.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				ts.Country = value.String
			}
		case tvsource.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ts.Description = value.String
			}
		case tvsource.FieldLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				ts.Language = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TvSource.
// Note that you need to call TvSource.Unwrap() before calling this method if this TvSource
// was returned from a transaction, and the transaction was committed or rolled back.
func (ts *TvSource) Update() *TvSourceUpdateOne {
	return (&TvSourceClient{config: ts.config}).UpdateOne(ts)
}

// Unwrap unwraps the TvSource entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ts *TvSource) Unwrap() *TvSource {
	tx, ok := ts.config.driver.(*txDriver)
	if !ok {
		panic("store: TvSource is not a transactional entity")
	}
	ts.config.driver = tx.drv
	return ts
}

// String implements the fmt.Stringer.
func (ts *TvSource) String() string {
	var builder strings.Builder
	builder.WriteString("TvSource(")
	builder.WriteString(fmt.Sprintf("id=%v", ts.ID))
	builder.WriteString(", logo=")
	builder.WriteString(ts.Logo)
	builder.WriteString(", video=")
	builder.WriteString(ts.Video)
	builder.WriteString(", title=")
	builder.WriteString(ts.Title)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", ts.Status))
	builder.WriteString(", country=")
	builder.WriteString(ts.Country)
	builder.WriteString(", description=")
	builder.WriteString(ts.Description)
	builder.WriteString(", language=")
	builder.WriteString(ts.Language)
	builder.WriteByte(')')
	return builder.String()
}

// TvSources is a parsable slice of TvSource.
type TvSources []*TvSource

func (ts TvSources) config(cfg config) {
	for _i := range ts {
		ts[_i].config = cfg
	}
}