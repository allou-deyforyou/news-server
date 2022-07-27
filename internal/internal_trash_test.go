package internal_test

import (
	"context"
	"news/internal/storage/custom"
)

// Cote d'ivoire

func CreateAbidjanNetSource() {
	entClient.Source.Create().
		SetStatus(true).
		SetName("Abidjan.Net").
		SetURL("https://news.abidjan.net").
		SetLogo("https://abidjan.net/public/img/favicon-32x32.png").
		SetArticleCategories(map[string]string{
			custom.PoliticsArticleCategory:      "politique",
			custom.EconomyArticleCategory:       "economie",
			custom.SocietyArticleCategory:       "societe",
			custom.SportArticleCategory:         "sport",
			custom.CultureArticleCategory:       "art-et-culture",
			custom.HealthArticleCategory:        "sante",
			custom.InternationalArticleCategory: "international",
			custom.MusicArticleCategory:         "musique",
		}).
		SetArticleFeaturedPostURL("/").
		SetArticleFeaturedPostSelector(&custom.SourcePostSelector{
			Title: []string{".title"},
			Image: []string{"img", "data-original", "src", "picture"},
			Date:  []string{".infos"},
			Link:  []string{"href"},
			List:  []string{".section-alaune > div > a", ".sub-content .section-mea:nth-child(1) > a"},
		}).
		SetArticleCategoryPostURL("/articles/%v?page=%v").
		SetArticleCategoryPostSelector(&custom.SourcePostSelector{
			Title: []string{".title"},
			Image: []string{"img", "data-original", "src"},
			Date:  []string{".infos"},
			Link:  []string{"href"},
			List:  []string{".grid3 > a"},
		}).
		SetArticleContentSelector(&custom.SourcePostSelector{
			Content: []string{".article-content .txt"},
		}).
		SaveX(context.Background())
}

func CreateAfrikMagSource() {
	entClient.Source.Create().
		SetStatus(true).
		SetName("AfrikMag").
		SetURL("https://www.afrikmag.com").
		SetLogo("https://www.afrikmag.com/favicon.ico").
		SetArticleCategories(map[string]string{
			custom.PoliticsArticleCategory: "197",
			custom.EconomyArticleCategory:  "8464",
			custom.SocietyArticleCategory:  "7020",
			custom.SportArticleCategory:    "8",
			custom.CultureArticleCategory:  "1233",
			custom.HealthArticleCategory:   "8464",
		}).
		SetArticleFeaturedPostURL("/").
		SetArticleFeaturedPostSelector(&custom.SourcePostSelector{
			Title: []string{".post-title"},
			Image: []string{"img", "data-src", ".big-thumb-left-box-inner", "data-lazy-bg"},
			Date:  []string{".date"},
			Link:  []string{".post-title a", "href"},
			List:  []string{"#tie-block_3151 .posts-items .post-item", "#tie-block_3151"},
		}).
		SetArticleCategoryPostURL("/wp-admin/admin-ajax.php").
		SetArticleCategoryPostSelector(&custom.SourcePostSelector{
			Title: []string{".post-title"},
			Image: []string{"img", "src"},
			Date:  []string{".date"},
			Link:  []string{".post-title a", "href"},
			List:  []string{"li"},
		}).
		SetArticleContentSelector(&custom.SourcePostSelector{
			Content: []string{"article .entry-content > p"},
		}).
		SaveX(context.Background())
}
