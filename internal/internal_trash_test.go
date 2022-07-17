package internal_test

// import (
// 	"context"
// 	"news/internal/store/schema"
// 	"testing"
// )

// // Cote d'ivoire

// func TestCreateAbidjanNetSource(t *testing.T) {
// 	entClient.NewsArticleSource.Create().
// 		SetStatus(true).
// 		SetName("Abidjan.Net").
// 		SetURL("https://news.abidjan.net").
// 		SetLogo("https://abidjan.net/public/img/favicon-32x32.png").
// 		SetCategories(map[string]string{
// 			schema.PoliticsArticleCategory:      "politique",
// 			schema.EconomyArticleCategory:       "economie",
// 			schema.SocietyArticleCategory:       "societe",
// 			schema.SportArticleCategory:         "sport",
// 			schema.CultureArticleCategory:       "art-et-culture",
// 			schema.HealthArticleCategory:        "sante",
// 			schema.InternationalArticleCategory: "international",
// 			schema.MusicArticleCategory:         "musique",
// 		}).
// 		SetLatestPostURL("/").
// 		SetLatestPostSelector(&schema.NewsPostSelector{
// 			Title: []string{".title"},
// 			Image: []string{"img", "data-original", "src", "picture"},
// 			Date:  []string{".infos"},
// 			Link:  []string{"href"},
// 			List:  []string{".section-alaune > div > a", ".sub-content .section-mea:nth-child(1) > a"},
// 		}).
// 		SetCategoryPostURL("/articles/%v?page=%v").
// 		SetCategoryPostSelector(&schema.NewsPostSelector{
// 			Title: []string{".title"},
// 			Image: []string{"img", "data-original", "src"},
// 			Date:  []string{".infos"},
// 			Link:  []string{"href"},
// 			List:  []string{".grid3 > a"},
// 		}).
// 		SetArticleSelector(&schema.NewsArticleSelector{
// 			Description: []string{".article-content .txt"},
// 		}).
// 		Save(context.Background())
// }

// func TestCreateAfrikMagSource(t *testing.T) {
// 	entClient.NewsArticleSource.Create().
// 		SetStatus(true).
// 		SetName("AfrikMag").
// 		SetURL("https://www.afrikmag.com").
// 		SetLogo("https://www.afrikmag.com/favicon.ico").
// 		SetCategories(map[string]string{
// 			schema.PoliticsArticleCategory: "197",
// 			schema.EconomyArticleCategory:  "8464",
// 			schema.SocietyArticleCategory:  "7020",
// 			schema.SportArticleCategory:    "8",
// 			schema.CultureArticleCategory:  "1233",
// 			schema.HealthArticleCategory:   "8464",
// 		}).
// 		SetLatestPostURL("/").
// 		SetLatestPostSelector(&schema.NewsPostSelector{
// 			Title: []string{".post-title"},
// 			Image: []string{"img", "data-src", ".big-thumb-left-box-inner", "data-lazy-bg"},
// 			Date:  []string{".date"},
// 			Link:  []string{".post-title a", "href"},
// 			List:  []string{"#tie-block_3151 .posts-items .post-item", "#tie-block_3151"},
// 		}).
// 		SetCategoryPostURL("/wp-admin/admin-ajax.php").
// 		SetCategoryPostSelector(&schema.NewsPostSelector{
// 			Title: []string{".post-title"},
// 			Image: []string{"img", "src"},
// 			Date:  []string{".date"},
// 			Link:  []string{".post-title a", "href"},
// 			List:  []string{"li"},
// 		}).
// 		SetArticleSelector(&schema.NewsArticleSelector{
// 			Description: []string{"article .entry-content > p"},
// 		}).
// 		Save(context.Background())
// }
