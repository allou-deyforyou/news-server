package schema

type Category string

const (
	Politics      Category = "politics"
	Economy       Category = "economy"
	Society       Category = "society"
	Sport         Category = "sport"
	Culture       Category = "culture"
	Health        Category = "health"
	International Category = "international"
	Music         Category = "music"
)

type NewsPost struct {
	Category string `json:"category,omitempty"`
	Source   string `json:"source,omitempty"`
	Title    string `json:"title,omitempty"`
	Image    string `json:"image,omitempty"`
	Date     string `json:"date,omitempty"`
	Link     string `json:"link,omitempty"`
	Logo     string `json:"logo,omitempty"`
}

type NewsPostSelector struct {
	Category []string `json:"category,omitempty"`
	Title    []string `json:"title,omitempty"`
	Image    []string `json:"image,omitempty"`
	Date     []string `json:"date,omitempty"`
	List     []string `json:"list,omitempty"`
	Link     []string `json:"link,omitempty"`
}

type NewsArticle struct {
	Description string `json:"description,omitempty"`
}

type NewsArticleSelector struct {
	Description []string `json:"description,omitempty"`
}
