package internal_test

import (
	"context"
	"log"
	"testing"

	"news/internal/storage"
	"news/internal/storage/custom"
	"news/internal/storage/migrate"

	"entgo.io/ent/dialect"

	_ "github.com/mattn/go-sqlite3"
)

var entClient *storage.Client

func init() {
	client, err := storage.Open(dialect.SQLite, "../yola.db?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	ctx := context.Background()
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	entClient = client
}

func TestGetNewsCategories(t *testing.T) {
	entClient.Categories.Delete().Exec(context.Background())
	// log.Println(entClient.Source.Query().AllX(context.Background()))
}

func TestCreateNewsCategories(t *testing.T) {
	entClient.Categories.Create().
		SetLanguage("fr").
		SetArticleCategories(map[string]string{
			custom.PoliticsArticleCategory:      "politique",
			custom.EconomyArticleCategory:       "économie",
			custom.SocietyArticleCategory:       "société",
			custom.SportArticleCategory:         "sport",
			custom.CultureArticleCategory:       "culture",
			custom.TechnologyArticleCategory:    "technologie",
			custom.HealthArticleCategory:        "santé",
			custom.InternationalArticleCategory: "international",
			custom.MusicArticleCategory:         "musique",
		}).
		SetMediaCategories(map[string]string{
			custom.BulletinMediaCategory: "journaux",
		}).
		SaveX(context.Background())
}

/// Source
///////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////

func TestGetSources(t *testing.T) {
	entClient.Source.Delete().Exec(context.Background())
	// log.Println(entClient.Source.Query().AllX(context.Background()))
}

// Cote d'ivoire

func TestCreateFratmatInfoSource(t *testing.T) {
	entClient.Source.Create().
		SetName("Fratmat Info").
		SetURL("https://www.fratmat.info").
		SetLogo("https://www.fratmat.info/theme_fratmat/images/favicon.ico").
		SetDescription("L'actualité en continu...").
		SetArticleCategories(map[string]string{
			custom.PoliticsArticleCategory: "politique",
			custom.EconomyArticleCategory:  "économie",
			custom.SocietyArticleCategory:  "société",
			custom.CultureArticleCategory:  "culture",
			custom.SportArticleCategory:    "sport",

			custom.FeaturedArticleCategory: "alaune",
		}).
		SetArticleFeaturedPostURL("/").
		SetArticleFeaturedPostSelector(&custom.SourcePostSelector{
			Title: []string{".article-title"},
			Image: []string{"img", "data-src"},
			Date:  []string{".publishTime"},
			Link:  []string{"a", "href"},
			List:  []string{".fratmat-more-articles .ajaxArticles .article-info"},
		}).
		SetArticleCategoryPostURL("/morearticles/%v?pgno=%v").
		SetArticleCategoryPostSelector(&custom.SourcePostSelector{
			Title: []string{".article-title"},
			Image: []string{"img", "data-src"},
			Date:  []string{".publishTime"},
			Link:  []string{"a", "href"},
			List:  []string{".fratmat-more-articles .ajaxArticles .article-info"},
		}).
		SetArticleContentSelector(&custom.SourcePostSelector{
			Content: []string{".body-desc", ".social-media"},
		}).
		SetMediaCategories(map[string]string{
			custom.BulletinMediaCategory:   "video",
			custom.FeaturedArticleCategory: "alaune",
		}).
		SetMediaCategoryPostURL("/morearticles/%v?pgno=%v").
		SetMediaCategoryPostSelector(&custom.SourcePostSelector{
			Title: []string{".article-title"},
			Image: []string{"img", "src"},
			Date:  []string{".publishTime"},
			Link:  []string{"a", "href"},
			List:  []string{".fratmat-more-articles .ajaxArticles .article-info"},
		}).
		SetMediaContentSelector(&custom.SourcePostSelector{
			Content: []string{".videoIframe", "src"},
		}).
		SaveX(context.Background())
}

func TestCreateRTIInfoSource(t *testing.T) {
	entClient.Source.Create().
		SetName("RTI Info").
		SetURL("https://vodadmin.rtireplays.com/api/v1/news/get").
		SetLogo("https://rti.info/icon.ico").
		SetDescription("Toute l'actualité de la Côte d'Ivoire chaque jour sur RTI Info.").
		SetArticleCategories(map[string]string{
			custom.EconomyArticleCategory:  "economie",
			custom.CultureArticleCategory:  "culture",
			custom.FeaturedArticleCategory: "alaune",
		}).
		SetArticleFeaturedPostURL("/home?page=1").
		SetArticleFeaturedPostSelector(&custom.SourcePostSelector{
			Title: []string{"title"},
			Image: []string{"image", "https://d2kpp1ajbxq1e1.cloudfront.net/news/image/%v"},
			Date:  []string{"created_at"},
			List:  []string{"ListTopArticles"},
			Link:  []string{"id", "/article/info/%v"},
		}).
		SetArticleCategoryPostURL("/articles/%v?page/%v").
		SetArticleCategoryPostSelector(&custom.SourcePostSelector{
			Title: []string{"title"},
			Image: []string{"image", "https://d2kpp1ajbxq1e1.cloudfront.net/news/image/%v"},
			Date:  []string{"created_at"},
			List:  []string{"ListArticles", "data"},
			Link:  []string{"id", "/article/info/%v"},
		}).
		SetArticleContentSelector(&custom.SourcePostSelector{
			Content: []string{
				"ArticleInfo",
				"description",
			},
		}).
		SetMediaCategories(map[string]string{
			custom.FeaturedArticleCategory: "alaune",
			custom.BulletinMediaCategory:   "journauxtv",
		}).
		SetMediaCategoryPostURL("/articles/%v?page/%v").
		SetMediaCategoryPostSelector(&custom.SourcePostSelector{
			Title: []string{"title"},
			Image: []string{"image", "https://d2kpp1ajbxq1e1.cloudfront.net/news/image/%v"},
			Date:  []string{"created_at"},
			List:  []string{"ListArticles", "data"},
			Link:  []string{"id", "/article/info/%v"},
		}).
		SetMediaContentSelector(&custom.SourcePostSelector{
			Content: []string{
				"ListVideos",
				"playlist",
				"sources",
				"file",
			},
		}).
		SaveX(context.Background())
}

func TestCreateYecloSource(t *testing.T) {
	entClient.Source.Create().
		SetName("Yeclo").
		SetURL("https://www.ivoiresoir.net").
		SetLogo("https://www.ivoiresoir.net/favicon.ico").
		SetDescription("L'actualité Africaine made in Côte d'Ivoire").
		SetArticleCategories(map[string]string{
			custom.EconomyArticleCategory:  "economie",
			custom.CultureArticleCategory:  "culture",
			custom.FeaturedArticleCategory: "alaune",
		}).
		SetArticleFeaturedPostURL("/").
		SetArticleFeaturedPostSelector(&custom.SourcePostSelector{
			Title: []string{".td-module-title"},
			Image: []string{".td-image-wrap span", "data-img-url"},
			Date:  []string{"time", "datetime"},
			Link:  []string{"a", "href"},
			List:  []string{"#tdi_74 .td_module_flex"},
		}).
		SetArticleCategoryPostURL("/%v/page/%v").
		SetArticleCategoryPostSelector(&custom.SourcePostSelector{
			Title: []string{".td-module-title"},
			Image: []string{".td-image-wrap span", "data-img-url"},
			Date:  []string{"time", "datetime"},
			Link:  []string{"a", "href"},
			List:  []string{"#tdi_64 .tdb_module_loop"},
		}).
		SetArticleContentSelector(&custom.SourcePostSelector{
			Content: []string{
				".td-post-content .tdb-block-inner > p, .td-post-content .tdb-block-inner > figure img, .td-post-content .tdb-block-inner > h2, .td-post-content .tdb-block-inner > ol",
			},
		}).
		SaveX(context.Background())
}

// International

func TestCreateFrance24Source(t *testing.T) {
	entClient.Source.Create().
		SetName("France 24").
		SetCountry("international").
		SetURL("https://www.france24.com").
		SetLogo("https://www.france24.com/favicon.ico").
		SetDescription("La chaîne française d'actualité internationale").
		SetArticleCategories(map[string]string{
			custom.EconomyArticleCategory:  "éco-tech",
			custom.SportArticleCategory:    "sports",
			custom.CultureArticleCategory:  "culture",
			custom.FeaturedArticleCategory: "alaune",
		}).
		SetArticleFeaturedPostURL("/fr/afrique").
		SetArticleFeaturedPostSelector(&custom.SourcePostSelector{
			Title: []string{".article__title"},
			Image: []string{"source", "srcset"},
			Link:  []string{"a", "href"},
			List:  []string{".t-content > .t-content__section-pb > div .o-layout-list__item"},
		}).
		SetArticleCategoryPostURL("/fr/%v").
		SetArticleCategoryPostSelector(&custom.SourcePostSelector{
			Title: []string{".article__title"},
			Image: []string{"source", "srcset"},
			Link:  []string{"a", "href"},
			List:  []string{".t-content > .t-content__section-pb > div .o-layout-list__item"},
		}).
		SetArticleContentSelector(&custom.SourcePostSelector{
			Content: []string{".t-content__chapo, .t-content__body > p, .t-content__body > h1, .t-content__body > h2, .t-content__body figure img"},
		}).
		SaveX(context.Background())
}

func TestCreateRFISource(t *testing.T) {
	entClient.Source.Create().
		SetName("RFI").
		SetCountry("international").
		SetURL("https://www.rfi.fr").
		SetLogo("https://www.rfi.fr/favicon.ico").
		SetDescription("La radio française d'actualité internationale").
		SetArticleCategories(map[string]string{
			custom.SportArticleCategory:    "afrique-foot",
			custom.CultureArticleCategory:  "culture-médias",
			custom.EconomyArticleCategory:  "économie",
			custom.FeaturedArticleCategory: "alaune",
		}).
		SetArticleFeaturedPostURL("/fr/afrique").
		SetArticleFeaturedPostSelector(&custom.SourcePostSelector{
			Title: []string{".article__title"},
			Image: []string{"source", "srcset"},
			Link:  []string{"a", "href"},
			List: []string{
				".t-content > .t-content__section-pb > div .o-layout-list__item",
			},
		}).
		SetArticleCategoryPostURL("/fr/%v").
		SetArticleCategoryPostSelector(&custom.SourcePostSelector{
			Title: []string{".article__title"},
			Image: []string{"source", "srcset"},
			Link:  []string{"a", "href"},
			List:  []string{".t-content > .t-content__section-pb > div .o-layout-list__item"},
		}).
		SetArticleContentSelector(&custom.SourcePostSelector{
			Content: []string{".t-content__chapo, .t-content__body > p, .t-content__body > h1, .t-content__body > h2, .t-content__body figure img"},
		}).
		SaveX(context.Background())
}

func TestCreateAfricaSource(t *testing.T) {
	entClient.Source.Create().
		SetName("Africa News").
		SetCountry("international").
		SetURL("https://fr.africanews.com").
		SetLogo("https://fr.africanews.com/favicon.ico").
		SetDescription("La chaîne d’information panafricaine").
		SetArticleCategories(map[string]string{
			custom.EconomyArticleCategory:    "economie",
			custom.TechnologyArticleCategory: "science-technologie",
			custom.SportArticleCategory:      "sport",
			custom.CultureArticleCategory:    "culture",
			custom.FeaturedArticleCategory:   "alaune",
		}).
		SetArticleFeaturedPostURL("/infos").
		SetArticleFeaturedPostSelector(&custom.SourcePostSelector{
			Title: []string{".teaser__title"},
			Image: []string{"img", "src"},
			Date:  []string{"time", "datetime"},
			Link:  []string{"a", "href"},
			List:  []string{".main-content article"},
		}).
		SetArticleCategoryPostURL("/%v").
		SetArticleCategoryPostSelector(&custom.SourcePostSelector{
			Title: []string{".teaser__title"},
			Image: []string{"img", "src"},
			Date:  []string{"time", "datetime"},
			Link:  []string{"a", "href"},
			List:  []string{".main-content article"},
		}).
		SetArticleContentSelector(&custom.SourcePostSelector{
			Content: []string{".article-content__text p"},
		}).
		SaveX(context.Background())
}

func TestCreateBBCSource(t *testing.T) {
	entClient.Source.Create().
		SetName("BBC").
		SetCountry("international").
		SetURL("https://www.bbc.com").
		SetLogo("https://www.bbc.com/favicon.ico").
		SetDescription("La Société de radiodiffusion britannique").
		SetArticleCategories(map[string]string{
			custom.EconomyArticleCategory:    "cnq687nr9v1t",
			custom.TechnologyArticleCategory: "cnq687nn703t",
			custom.HealthArticleCategory:     "c06gq9jxz3rt",
			custom.CultureArticleCategory:    "cnq687nrrw8t",
			custom.FeaturedArticleCategory:   "alaune",
		}).
		SetArticleFeaturedPostURL("/afrique").
		SetArticleFeaturedPostSelector(&custom.SourcePostSelector{
			Title: []string{"h3"},
			Image: []string{"picture source:nth-child(1)", "srcset"},
			Date:  []string{"time", "datetime"},
			Link:  []string{"a", "href"},
			List:  []string{"main section:nth-child(1) li"},
		}).
		SetArticleCategoryPostURL("/afrique/topics/%v?page=%v").
		SetArticleCategoryPostSelector(&custom.SourcePostSelector{
			Title: []string{"h2"},
			Image: []string{"picture source:nth-child(1)", "srcset"},
			Date:  []string{"time", "datetime"},
			Link:  []string{"a", "href"},
			List:  []string{"main li"},
		}).
		SetArticleContentSelector(&custom.SourcePostSelector{
			Content: []string{
				"main > div:not(:nth-child(-n+4)) > p, main > div:not(:nth-child(-n+4)) > figure img, main > div:not(:nth-child(-n+4)) > h2, main > div:not(:nth-child(-n+4)) > ul",
				"main noscript",
			},
		}).
		SaveX(context.Background())
}
