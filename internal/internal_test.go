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
			Image: []string{"img", "data-original", "src"},
			Date:  []string{".infos"},
			Link:  []string{"href"},
			List:  []string{".grid3 > a", ".section-mea a"},
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
			List:  []string{"#tie-block_3151 .posts-items .post-item"},
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

func TestGetSources(t *testing.T) {
	entClient.NewsSource.Delete().Exec(context.Background())
	log.Println(entClient.NewsSource.Query().AllX(context.Background()))
}
