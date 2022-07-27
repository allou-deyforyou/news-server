// Code generated by entc, DO NOT EDIT.

package storage

import (
	"news/internal/storage/articlepost"
	"news/internal/storage/categories"
	"news/internal/storage/mediapost"
	"news/internal/storage/schema"
	"news/internal/storage/source"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	articlepostFields := schema.ArticlePost{}.Fields()
	_ = articlepostFields
	// articlepostDescStatus is the schema descriptor for status field.
	articlepostDescStatus := articlepostFields[0].Descriptor()
	// articlepost.DefaultStatus holds the default value on creation for the status field.
	articlepost.DefaultStatus = articlepostDescStatus.Default.(bool)
	categoriesFields := schema.Categories{}.Fields()
	_ = categoriesFields
	// categoriesDescLanguage is the schema descriptor for language field.
	categoriesDescLanguage := categoriesFields[2].Descriptor()
	// categories.DefaultLanguage holds the default value on creation for the language field.
	categories.DefaultLanguage = categoriesDescLanguage.Default.(string)
	// categoriesDescStatus is the schema descriptor for status field.
	categoriesDescStatus := categoriesFields[3].Descriptor()
	// categories.DefaultStatus holds the default value on creation for the status field.
	categories.DefaultStatus = categoriesDescStatus.Default.(bool)
	mediapostFields := schema.MediaPost{}.Fields()
	_ = mediapostFields
	// mediapostDescStatus is the schema descriptor for status field.
	mediapostDescStatus := mediapostFields[0].Descriptor()
	// mediapost.DefaultStatus holds the default value on creation for the status field.
	mediapost.DefaultStatus = mediapostDescStatus.Default.(bool)
	// mediapostDescLive is the schema descriptor for live field.
	mediapostDescLive := mediapostFields[1].Descriptor()
	// mediapost.DefaultLive holds the default value on creation for the live field.
	mediapost.DefaultLive = mediapostDescLive.Default.(bool)
	sourceFields := schema.Source{}.Fields()
	_ = sourceFields
	// sourceDescLanguage is the schema descriptor for language field.
	sourceDescLanguage := sourceFields[12].Descriptor()
	// source.DefaultLanguage holds the default value on creation for the language field.
	source.DefaultLanguage = sourceDescLanguage.Default.(string)
	// sourceDescCountry is the schema descriptor for country field.
	sourceDescCountry := sourceFields[13].Descriptor()
	// source.DefaultCountry holds the default value on creation for the country field.
	source.DefaultCountry = sourceDescCountry.Default.(string)
	// sourceDescStatus is the schema descriptor for status field.
	sourceDescStatus := sourceFields[14].Descriptor()
	// source.DefaultStatus holds the default value on creation for the status field.
	source.DefaultStatus = sourceDescStatus.Default.(bool)
	// sourceDescLogo is the schema descriptor for logo field.
	sourceDescLogo := sourceFields[15].Descriptor()
	// source.LogoValidator is a validator for the "logo" field. It is called by the builders before save.
	source.LogoValidator = sourceDescLogo.Validators[0].(func(string) error)
}
