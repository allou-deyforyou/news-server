package internal_test

import (
	"context"
	"fmt"
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

/// NewsSource
///////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////

// Cote d'ivoire

func TestCreateFratmatInfoSource(t *testing.T) {
	entClient.NewsSource.Create().
		SetStatus(true).
		SetName("Fratmat Info").
		SetURL("https://www.fratmat.info").
		SetLogo("https://www.fratmat.info/theme_fratmat/images/favicon.ico").
		SetCategories([]string{
			fmt.Sprintf("%v:politique", schema.Politics),
			fmt.Sprintf("%v:économie", schema.Economy),
			fmt.Sprintf("%v:société", schema.Society),
			fmt.Sprintf("%v:sport", schema.Sport),
			fmt.Sprintf("%v:culture", schema.Culture),
		}).
		SetLatestPostURL("/").
		SetLatestPostSelector(&schema.NewsPostSelector{
			Category: []string{".single_article_tag"},
			Title:    []string{".article-title"},
			Image:    []string{"img", "data-src"},
			Date:     []string{".publishTime"},
			Link:     []string{"a", "href"},
			List:     []string{".fratmat-more-articles .ajaxArticles .article-info"},
		}).
		SetCategoryPostURL("/morearticles/%v?pgno=%v").
		SetCategoryPostSelector(&schema.NewsPostSelector{
			Category: []string{".single_article_tag"},
			Title:    []string{".article-title"},
			Image:    []string{"img", "data-src"},
			Date:     []string{".publishTime"},
			Link:     []string{"a", "href"},
			List:     []string{".fratmat-more-articles .ajaxArticles .article-info"},
		}).
		SetArticleSelector(&schema.NewsArticleSelector{
			Description: []string{".body-desc div:nth-child(3)"},
		}).
		Save(context.Background())
}

func TestCreateAbidjanNetSource(t *testing.T) {
	entClient.NewsSource.Create().
		SetStatus(true).
		SetName("Abidjan.Net").
		SetURL("https://news.abidjan.net").
		SetLogo("https://abidjan.net/public/img/favicon-32x32.png").
		SetCategories([]string{
			fmt.Sprintf("%v:politique", schema.Politics),
			fmt.Sprintf("%v:economie", schema.Economy),
			fmt.Sprintf("%v:societe", schema.Society),
			fmt.Sprintf("%v:sport", schema.Sport),
			fmt.Sprintf("%v:art-et-culture", schema.Culture),
			fmt.Sprintf("%v:sante", schema.Health),
			fmt.Sprintf("%v:international", schema.International),
			fmt.Sprintf("%v:musique", schema.Music),
		}).
		SetLatestPostURL("/").
		SetLatestPostSelector(&schema.NewsPostSelector{
			Title: []string{".title"},
			Image: []string{"img", "data-original", "src", "picture"},
			Date:  []string{".infos"},
			Link:  []string{"href"},
			List:  []string{".section-alaune > div > a", ".sub-content .section-mea:nth-child(1) > a"},
		}).
		SetCategoryPostURL("/articles/%v?page=%v").
		SetCategoryPostSelector(&schema.NewsPostSelector{
			Title: []string{".title"},
			Image: []string{"img", "data-original", "src"},
			Date:  []string{".infos"},
			Link:  []string{"href"},
			List:  []string{".grid3 > a"},
		}).
		SetArticleSelector(&schema.NewsArticleSelector{
			Description: []string{".article-content .txt"},
		}).
		Save(context.Background())
}

func TestCreateAfrikMagSource(t *testing.T) {
	entClient.NewsSource.Create().
		SetStatus(true).
		SetName("AfrikMag").
		SetURL("https://www.afrikmag.com").
		SetLogo("https://www.afrikmag.com/favicon.ico").
		SetCategories([]string{
			fmt.Sprintf("%v:197", schema.Politics),
			fmt.Sprintf("%v:8464", schema.Economy),
			fmt.Sprintf("%v:7020", schema.Society),
			fmt.Sprintf("%v:8", schema.Sport),
			fmt.Sprintf("%v:1233", schema.Culture),
			fmt.Sprintf("%v:8464", schema.Health),
		}).
		SetLatestPostURL("/").
		SetLatestPostSelector(&schema.NewsPostSelector{
			Title: []string{".post-title"},
			Image: []string{"img", "data-src", ".big-thumb-left-box-inner", "data-lazy-bg"},
			Date:  []string{".date"},
			Link:  []string{".post-title a", "href"},
			List:  []string{"#tie-block_3151 .posts-items .post-item", "#tie-block_3151"},
		}).
		SetCategoryPostURL("/wp-admin/admin-ajax.php").
		SetCategoryPostSelector(&schema.NewsPostSelector{
			Title: []string{".post-title"},
			Image: []string{"img", "src"},
			Date:  []string{".date"},
			Link:  []string{".post-title a", "href"},
			List:  []string{"li"},
		}).
		SetArticleSelector(&schema.NewsArticleSelector{
			Description: []string{"article .entry-content > p"},
		}).
		Save(context.Background())
}

// International

func TestCreateFrance24Source(t *testing.T) {
	entClient.NewsSource.Create().
		SetStatus(true).
		SetName("France 24").
		SetURL("https://www.france24.com").
		SetLogo("https://www.france24.com/favicon.ico").
		SetCategories([]string{
			fmt.Sprintf("%v:éco-tech", schema.Economy),
			fmt.Sprintf("%v:sports", schema.Sport),
			fmt.Sprintf("%v:culture", schema.Culture),
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
			Description: []string{".t-content__chapo, .t-content__body p, .t-content__body h1"},
		}).
		Save(context.Background())
}

func TestCreateRfiSource(t *testing.T) {
	entClient.NewsSource.Create().
		SetStatus(true).
		SetName("RFI").
		SetURL("https://www.rfi.fr").
		SetLogo("https://www.rfi.fr/favicon.ico").
		SetCategories([]string{
			fmt.Sprintf("%v:afrique-foot", schema.Sport),
			fmt.Sprintf("%v:culture-médias", schema.Culture),
			fmt.Sprintf("%v:économie", schema.Economy),
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
			Description: []string{".t-content__chapo, .t-content__body p, .t-content__body h1"},
		}).
		Save(context.Background())
}

func TestCreateAfricaNewsSource(t *testing.T) {
	entClient.NewsSource.Create().
		SetStatus(true).
		SetName("Africa News").
		SetURL("https://fr.africanews.com").
		SetLogo("https://fr.africanews.com/favicon.ico").
		SetCategories([]string{
			fmt.Sprintf("%v:economie", schema.Economy),
			fmt.Sprintf("%v:science-technologie", schema.Technology),
			fmt.Sprintf("%v:sport", schema.Sport),
			fmt.Sprintf("%v:culture", schema.Culture),
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

func TestGetNewsSources(t *testing.T) {
	entClient.NewsSource.Delete().Exec(context.Background())
	log.Println(entClient.NewsSource.Query().AllX(context.Background()))
}

/// TvSource
///////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////

func TestCreateRti1TV(t *testing.T) {
	entClient.TvSource.Create().
		SetStatus(true).
		SetTitle("RTI 1").
		SetVideo("https://www.enovativecdn.com/rticdn/smil:rti1.smil/playlist.m3u8").
		SetLogo("https://rti.ci/_nuxt/img/6a12f31.svg").
		SetDescription("RTI 1 est la première chaîne de télévision généraliste publique ivoirienne qui émet en continu depuis Abidjan.").
		SaveX(context.Background())
}

func TestCreateRti2TV(t *testing.T) {
	entClient.TvSource.Create().
		SetStatus(true).
		SetTitle("RTI 2").
		SetVideo("https://www.enovativecdn.com/rticdn/smil:rti2.smil/playlist.m3u8").
		SetLogo("https://rti.ci/_nuxt/img/6d85e57.svg").
		SetDescription("RTI 2 est une chaîne de télévision généraliste publique ivoirienne. Elle est consacrée à la jeunesse.").
		SaveX(context.Background())

}

func TestCreateRti3TV(t *testing.T) {
	entClient.TvSource.Create().
		SetStatus(true).
		SetTitle("La 3").
		SetVideo("https://www.enovativecdn.com/rticdn/smil:rti3.smil/playlist.m3u8").
		SetLogo("https://rti.ci/_nuxt/img/4da62df.svg").
		SetDescription("RTI 3, baptisée La 3, est une chaîne de télévision ivoirienne du Groupe RTI lancée le 16 février 2020.").
		SaveX(context.Background())

}

func TestCreateNciTV(t *testing.T) {
	entClient.TvSource.Create().
		SetStatus(true).
		SetTitle("NCI").
		SetVideo("https://nci-live.secure2.footprint.net/nci/nci.isml/.m3u8").
		SetLogo("https://static.wixstatic.com/media/f8668c_8cf416367fb743378ec26c7e7978a318~mv2_d_1692_1295_s_2.png").
		SetDescription("La Nouvelle Chaîne Ivoirienne, plus connue sous le sigle NCI est la première chaîne de télévision privée ivoirienne.").
		SaveX(context.Background())
}

func TestGetTvSources(t *testing.T) {
	entClient.TvSource.Delete().Exec(context.Background())
	log.Println(entClient.TvSource.Query().AllX(context.Background()))
}
