// Code generated by entc, DO NOT EDIT.

package store

import (
	"encoding/json"
	"fmt"
	"news/internal/store/newscategories"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// NewsCategories is the model entity for the NewsCategories schema.
type NewsCategories struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Status holds the value of the "status" field.
	Status bool `json:"status,omitempty"`
	// Language holds the value of the "language" field.
	Language string `json:"language,omitempty"`
	// TvCategories holds the value of the "tv_categories" field.
	TvCategories map[string]string `json:"tv_categories,omitempty"`
	// ArticleCategories holds the value of the "article_categories" field.
	ArticleCategories map[string]string `json:"article_categories,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NewsCategories) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case newscategories.FieldTvCategories, newscategories.FieldArticleCategories:
			values[i] = new([]byte)
		case newscategories.FieldStatus:
			values[i] = new(sql.NullBool)
		case newscategories.FieldID:
			values[i] = new(sql.NullInt64)
		case newscategories.FieldLanguage:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type NewsCategories", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NewsCategories fields.
func (nc *NewsCategories) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case newscategories.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			nc.ID = int(value.Int64)
		case newscategories.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				nc.Status = value.Bool
			}
		case newscategories.FieldLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				nc.Language = value.String
			}
		case newscategories.FieldTvCategories:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tv_categories", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &nc.TvCategories); err != nil {
					return fmt.Errorf("unmarshal field tv_categories: %w", err)
				}
			}
		case newscategories.FieldArticleCategories:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field article_categories", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &nc.ArticleCategories); err != nil {
					return fmt.Errorf("unmarshal field article_categories: %w", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this NewsCategories.
// Note that you need to call NewsCategories.Unwrap() before calling this method if this NewsCategories
// was returned from a transaction, and the transaction was committed or rolled back.
func (nc *NewsCategories) Update() *NewsCategoriesUpdateOne {
	return (&NewsCategoriesClient{config: nc.config}).UpdateOne(nc)
}

// Unwrap unwraps the NewsCategories entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (nc *NewsCategories) Unwrap() *NewsCategories {
	tx, ok := nc.config.driver.(*txDriver)
	if !ok {
		panic("store: NewsCategories is not a transactional entity")
	}
	nc.config.driver = tx.drv
	return nc
}

// String implements the fmt.Stringer.
func (nc *NewsCategories) String() string {
	var builder strings.Builder
	builder.WriteString("NewsCategories(")
	builder.WriteString(fmt.Sprintf("id=%v", nc.ID))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", nc.Status))
	builder.WriteString(", language=")
	builder.WriteString(nc.Language)
	builder.WriteString(", tv_categories=")
	builder.WriteString(fmt.Sprintf("%v", nc.TvCategories))
	builder.WriteString(", article_categories=")
	builder.WriteString(fmt.Sprintf("%v", nc.ArticleCategories))
	builder.WriteByte(')')
	return builder.String()
}

// NewsCategoriesSlice is a parsable slice of NewsCategories.
type NewsCategoriesSlice []*NewsCategories

func (nc NewsCategoriesSlice) config(cfg config) {
	for _i := range nc {
		nc[_i].config = cfg
	}
}
