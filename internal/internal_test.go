package internal_test

import (
	"context"
	"log"
	"testing"

	"news/internal/store"
	"news/internal/store/migrate"
	"news/internal/store/schema"

	"entgo.io/ent/dialect"

	_ "github.com/mattn/go-sqlite3"
)

var entClient *store.Client

func init() {
	client, err := store.Open(dialect.SQLite, "../yola.db?mode=memory&cache=shared&_fk=1")
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

func TestCreateNewsCategories(t *testing.T) {
	entClient.NewsCategories.Create().
		SetArticleCategories([]string{
			schema.PoliticsArticleCategory,
			schema.EconomyArticleCategory,
			schema.SocietyArticleCategory,
			schema.SportArticleCategory,
			schema.CultureArticleCategory,
			schema.TechnologyArticleCategory,
			schema.HealthArticleCategory,
			schema.InternationalArticleCategory,
			schema.MusicArticleCategory,
		}).
		SetTvCategories([]string{}).
		SaveX(context.Background())
}

func TestGetNewsCategories(t *testing.T) {
	entClient.NewsCategories.Delete().Exec(context.Background())
	// log.Println(entClient.NewsArticleSource.Query().AllX(context.Background()))
}

/// NewsArticleSource
///////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////

// Cote d'ivoire

func TestCreateYecloSource(t *testing.T) {
	entClient.NewsArticleSource.Create().
		SetStatus(true).
		SetName("Yeclo").
		SetURL("https://www.ivoiresoir.net").
		SetLogo("https://www.ivoiresoir.net/favicon.ico").
		SetCategories(map[string]string{
			schema.EconomyArticleCategory: "economie",
			schema.CultureArticleCategory: "culture",
		}).
		SetLatestPostURL("/").
		SetLatestPostSelector(&schema.NewsPostSelector{
			Title: []string{".td-module-title"},
			Image: []string{".td-image-wrap span", "data-img-url"},
			Date:  []string{"time", "datetime"},
			Link:  []string{"a", "href"},
			List:  []string{"#tdi_74 .td_module_flex"},
		}).
		SetCategoryPostURL("/%v/page/%v").
		SetCategoryPostSelector(&schema.NewsPostSelector{
			Title: []string{".td-module-title"},
			Image: []string{".td-image-wrap span", "data-img-url"},
			Date:  []string{"time", "datetime"},
			Link:  []string{"a", "href"},
			List:  []string{"#tdi_64 .tdb_module_loop"},
		}).
		SetArticleSelector(&schema.NewsArticleSelector{
			Description: []string{
				".td-post-content .tdb-block-inner > p, .td-post-content .tdb-block-inner > figure img, .td-post-content .tdb-block-inner > h2, .td-post-content .tdb-block-inner > ol",
			},
		}).
		Save(context.Background())
}

func TestCreateRTISource(t *testing.T) {
	entClient.NewsArticleSource.Create().
		SetStatus(true).
		SetName("RTI Info").
		SetURL("https://vodadmin.rtireplays.com/api/v1/news/get").
		SetLogo("https://rti.info/icon.ico").
		SetCategories(map[string]string{
			schema.EconomyArticleCategory: "economie",
			schema.CultureArticleCategory: "culture",
		}).
		SetLatestPostURL("/home?page=1").
		SetLatestPostSelector(&schema.NewsPostSelector{
			Title: []string{"title"},
			Image: []string{"image", "https://d2kpp1ajbxq1e1.cloudfront.net/news/image/%v"},
			Date:  []string{"created_at"},
			List:  []string{"ListTopArticles"},
			Link:  []string{"id", "/article/info/%v"},
		}).
		SetCategoryPostURL("/articles/%v?page/%v").
		SetCategoryPostSelector(&schema.NewsPostSelector{
			Title: []string{"title"},
			Image: []string{"image", "https://d2kpp1ajbxq1e1.cloudfront.net/news/image/%v"},
			Date:  []string{"created_at"},
			List:  []string{"ListArticles", "data"},
			Link:  []string{"id", "/article/info/%v"},
		}).
		SetArticleSelector(&schema.NewsArticleSelector{
			Description: []string{
				"ArticleInfo",
				"description",
			},
		}).
		Save(context.Background())
}

// International

func TestCreateFrance24Source(t *testing.T) {
	entClient.NewsArticleSource.Create().
		SetStatus(true).
		SetName("France 24").
		SetCountry("world").
		SetURL("https://www.france24.com").
		SetLogo("https://www.france24.com/favicon.ico").
		SetCategories(map[string]string{
			schema.EconomyArticleCategory: "éco-tech",
			schema.SportArticleCategory:   "sports",
			schema.CultureArticleCategory: "culture",
		}).
		SetLatestPostURL("/fr/afrique").
		SetLatestPostSelector(&schema.NewsPostSelector{
			Title: []string{".article__title"},
			Image: []string{"source", "srcset"},
			Link:  []string{"a", "href"},
			List: []string{
				".t-content > .t-content__section-pb > div .o-layout-list__item",
			},
		}).
		SetCategoryPostURL("/fr/%v").
		SetCategoryPostSelector(&schema.NewsPostSelector{
			Title: []string{".article__title"},
			Image: []string{"source", "srcset"},
			Link:  []string{"a", "href"},
			List: []string{
				".t-content > .t-content__section-pb > div .o-layout-list__item",
			},
		}).
		SetArticleSelector(&schema.NewsArticleSelector{
			Description: []string{".t-content__chapo, .t-content__body > p, .t-content__body > h1, .t-content__body > h2, .t-content__body figure img"},
		}).
		Save(context.Background())
}

func TestCreateRFISource(t *testing.T) {
	entClient.NewsArticleSource.Create().
		SetStatus(true).
		SetName("RFI").
		SetCountry("world").
		SetURL("https://www.rfi.fr").
		SetLogo("https://www.rfi.fr/favicon.ico").
		SetCategories(map[string]string{
			schema.SportArticleCategory:   "afrique-foot",
			schema.CultureArticleCategory: "culture-médias",
			schema.EconomyArticleCategory: "économie",
		}).
		SetLatestPostURL("/fr/afrique").
		SetLatestPostSelector(&schema.NewsPostSelector{
			Title: []string{".article__title"},
			Image: []string{"source", "srcset"},
			Link:  []string{"a", "href"},
			List: []string{
				".t-content > .t-content__section-pb > div .o-layout-list__item",
			},
		}).
		SetCategoryPostURL("/fr/%v").
		SetCategoryPostSelector(&schema.NewsPostSelector{
			Title: []string{".article__title"},
			Image: []string{"source", "srcset"},
			Link:  []string{"a", "href"},
			List: []string{
				".t-content > .t-content__section-pb > div .o-layout-list__item",
			},
		}).
		SetArticleSelector(&schema.NewsArticleSelector{
			Description: []string{".t-content__chapo, .t-content__body > p, .t-content__body > h1, .t-content__body > h2, .t-content__body figure img"},
		}).
		Save(context.Background())
}

func TestCreateAfricaNewsArticleSource(t *testing.T) {
	entClient.NewsArticleSource.Create().
		SetStatus(true).
		SetName("Africa News").
		SetCountry("world").
		SetURL("https://fr.africanews.com").
		SetLogo("https://fr.africanews.com/favicon.ico").
		SetCategories(map[string]string{
			schema.EconomyArticleCategory:    "economie",
			schema.TechnologyArticleCategory: "science-technologie",
			schema.SportArticleCategory:      "sport",
			schema.CultureArticleCategory:    "culture",
		}).
		SetLatestPostURL("/infos").
		SetLatestPostSelector(&schema.NewsPostSelector{
			Title: []string{".teaser__title"},
			Image: []string{"img", "src"},
			Date:  []string{"time", "datetime"},
			Link:  []string{"a", "href"},
			List:  []string{".main-content article"},
		}).
		SetCategoryPostURL("/%v").
		SetCategoryPostSelector(&schema.NewsPostSelector{
			Title: []string{".teaser__title"},
			Image: []string{"img", "src"},
			Date:  []string{"time", "datetime"},
			Link:  []string{"a", "href"},
			List:  []string{".main-content article"},
		}).
		SetArticleSelector(&schema.NewsArticleSelector{
			Description: []string{".article-content__text p"},
		}).
		Save(context.Background())
}

func TestCreateBBCSource(t *testing.T) {
	entClient.NewsArticleSource.Create().
		SetStatus(true).
		SetName("BBC").
		SetCountry("world").
		SetURL("https://www.bbc.com").
		SetLogo("https://ichef.bbci.co.uk/favicon.ico").
		SetCategories(map[string]string{
			schema.EconomyArticleCategory:    "cnq687nr9v1t",
			schema.TechnologyArticleCategory: "cnq687nn703t",
			schema.HealthArticleCategory:     "c06gq9jxz3rt",
			schema.CultureArticleCategory:    "cnq687nrrw8t",
		}).
		SetLatestPostURL("/afrique").
		SetLatestPostSelector(&schema.NewsPostSelector{
			Title: []string{"h3"},
			Image: []string{"picture source:nth-child(1)", "srcset"},
			Date:  []string{"time", "datetime"},
			Link:  []string{"a", "href"},
			List:  []string{"main section:nth-child(1) li"},
		}).
		SetCategoryPostURL("/afrique/topics/%v?page=%v").
		SetCategoryPostSelector(&schema.NewsPostSelector{
			Title: []string{"h3"},
			Image: []string{"picture source:nth-child(1)", "srcset"},
			Date:  []string{"time", "datetime"},
			Link:  []string{"a", "href"},
			List:  []string{"main li"},
		}).
		SetArticleSelector(&schema.NewsArticleSelector{
			Description: []string{
				"main > div:not(:nth-child(-n+4)) > p, main > div:not(:nth-child(-n+4)) > figure img, main > div:not(:nth-child(-n+4)) > h2, main > div:not(:nth-child(-n+4)) > ul",
				"main noscript",
			},
		}).
		Save(context.Background())
}

func TestGetNewsArticleSources(t *testing.T) {
	entClient.NewsArticleSource.Delete().Exec(context.Background())
	// log.Println(entClient.NewsArticleSource.Query().AllX(context.Background()))
}

/// NewsTvSource
///////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////

func TestCreateRti1TV(t *testing.T) {
	entClient.NewsTvSource.Create().
		SetLive(true).
		SetStatus(true).
		SetSource("RTI 1").
		SetVideo("https://www.enovativecdn.com/rticdn/smil:rti1.smil/playlist.m3u8").
		SetLogo("https://rti.ci/_nuxt/img/6a12f31.svg").
		SetDescription("La première chaîne de télévision publique ivoirienne du Groupe RTI").
		SaveX(context.Background())
}

func TestCreateRti2TV(t *testing.T) {
	entClient.NewsTvSource.Create().
		SetLive(true).
		SetStatus(true).
		SetSource("RTI 2").
		SetVideo("https://www.enovativecdn.com/rticdn/smil:rti2.smil/playlist.m3u8").
		SetLogo("https://rti.ci/_nuxt/img/6d85e57.svg").
		SetDescription("Une nouvelle chaîne de télévision ivoirienne du Groupe RTI").
		SaveX(context.Background())
}

func TestCreateRti3TV(t *testing.T) {
	entClient.NewsTvSource.Create().
		SetLive(true).
		SetStatus(true).
		SetSource("La 3").
		SetVideo("https://www.enovativecdn.com/rticdn/smil:rti3.smil/playlist.m3u8").
		SetLogo("https://rti.ci/_nuxt/img/4da62df.svg").
		SetDescription("Appelée aussi RTI 3, une nouvelle chaîne de télévision ivoirienne du Groupe RTI").
		SaveX(context.Background())

}

func TestCreateNciTV(t *testing.T) {
	entClient.NewsTvSource.Create().
		SetLive(true).
		SetStatus(true).
		SetSource("NCI").
		SetVideo("https://nci-live.secure2.footprint.net/nci/nci.isml/.m3u8").
		SetLogo("https://static.wixstatic.com/media/f8668c_8cf416367fb743378ec26c7e7978a318~mv2_d_1692_1295_s_2.png").
		SetDescription("La Nouvelle Chaîne Ivoirienne").
		SaveX(context.Background())
}

func TestGetNewsTvSources(t *testing.T) {
	entClient.NewsTvSource.Delete().Exec(context.Background())
	// log.Println(entClient.NewsTvSource.Query().AllX(context.Background()))
}
