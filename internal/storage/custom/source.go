package custom

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

	FeaturedArticleCategory = "featured"
)

const (
	BulletinMediaCategory = "bulletin"
	InternationalMediaCategory = "international"
)

type SourcePostSelector struct {
	Content []string `json:"content,omitempty"`
	Title   []string `json:"title,omitempty"`
	Image   []string `json:"image,omitempty"`
	Date    []string `json:"date,omitempty"`
	List    []string `json:"list,omitempty"`
	Link    []string `json:"link,omitempty"`
}
