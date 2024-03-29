// Code generated by entc, DO NOT EDIT.

package source

const (
	// Label holds the string label denoting the source type in the database.
	Label = "source"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldArticleFeaturedPostSelector holds the string denoting the article_featured_post_selector field in the database.
	FieldArticleFeaturedPostSelector = "article_featured_post_selector"
	// FieldArticleFeaturedPostURL holds the string denoting the article_featured_post_url field in the database.
	FieldArticleFeaturedPostURL = "article_featured_post_url"
	// FieldArticleCategoryPostSelector holds the string denoting the article_category_post_selector field in the database.
	FieldArticleCategoryPostSelector = "article_category_post_selector"
	// FieldArticleCategoryPostURL holds the string denoting the article_category_post_url field in the database.
	FieldArticleCategoryPostURL = "article_category_post_url"
	// FieldArticleContentSelector holds the string denoting the article_content_selector field in the database.
	FieldArticleContentSelector = "article_content_selector"
	// FieldMediaFeaturedPostSelector holds the string denoting the media_featured_post_selector field in the database.
	FieldMediaFeaturedPostSelector = "media_featured_post_selector"
	// FieldMediaFeaturedPostURL holds the string denoting the media_featured_post_url field in the database.
	FieldMediaFeaturedPostURL = "media_featured_post_url"
	// FieldMediaCategoryPostSelector holds the string denoting the media_category_post_selector field in the database.
	FieldMediaCategoryPostSelector = "media_category_post_selector"
	// FieldMediaCategoryPostURL holds the string denoting the media_category_post_url field in the database.
	FieldMediaCategoryPostURL = "media_category_post_url"
	// FieldMediaContentSelector holds the string denoting the media_content_selector field in the database.
	FieldMediaContentSelector = "media_content_selector"
	// FieldArticleCategories holds the string denoting the article_categories field in the database.
	FieldArticleCategories = "article_categories"
	// FieldMediaCategories holds the string denoting the media_categories field in the database.
	FieldMediaCategories = "media_categories"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldLogo holds the string denoting the logo field in the database.
	FieldLogo = "logo"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// Table holds the table name of the source in the database.
	Table = "sources"
)

// Columns holds all SQL columns for source fields.
var Columns = []string{
	FieldID,
	FieldArticleFeaturedPostSelector,
	FieldArticleFeaturedPostURL,
	FieldArticleCategoryPostSelector,
	FieldArticleCategoryPostURL,
	FieldArticleContentSelector,
	FieldMediaFeaturedPostSelector,
	FieldMediaFeaturedPostURL,
	FieldMediaCategoryPostSelector,
	FieldMediaCategoryPostURL,
	FieldMediaContentSelector,
	FieldArticleCategories,
	FieldMediaCategories,
	FieldDescription,
	FieldLanguage,
	FieldCountry,
	FieldStatus,
	FieldLogo,
	FieldName,
	FieldURL,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultLanguage holds the default value on creation for the "language" field.
	DefaultLanguage string
	// DefaultCountry holds the default value on creation for the "country" field.
	DefaultCountry string
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus bool
	// LogoValidator is a validator for the "logo" field. It is called by the builders before save.
	LogoValidator func(string) error
)
