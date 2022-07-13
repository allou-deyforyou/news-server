package schema

const (
	PoliticsArticleCategory      = "politics"
	EconomyArticleCategory       = "economy"
	SocietyArticleCategory       = "society"
	SportArticleCategory         = "sport"
	CultureArticleCategory       = "culture"
	TechnologyArticleCategory    = "technology"
	HealthArticleCategory        = "health"
	InternationalArticleCategory = "international"
	MusicArticleCategory         = "music"
)

type NewsPostSelector struct {
	Category []string `json:"category,omitempty"`
	Title    []string `json:"title,omitempty"`
	Image    []string `json:"image,omitempty"`
	Date     []string `json:"date,omitempty"`
	List     []string `json:"list,omitempty"`
	Link     []string `json:"link,omitempty"`
}

type NewsArticleSelector struct {
	Description []string `json:"description,omitempty"`
}
