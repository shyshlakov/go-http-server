package restapi

type Tag struct {
	Tags []string `json:"tags"`
	// TagName  string   `json:"-"`
	// TagId    string   `json:"tag_id_123"`
	// TestName string   `json:"tag_name,omitempty"` //не появляется если пустое
}

type Article struct { //responsePayload
	ID             string   `json:"id"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
	Slug           string   `json:"slug"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Body           string   `json:"body"`
	TagList        []string `json:"tag_list"`
	Favorited      bool     `json:"favorited"`
	FavoritesCount int      `json:"favorites_count"`
	//Comment        map[string]interface{} `json:"comment"`
	Score      float64  `json:"score"`
	LikedUsers []string `json:"liked_users"`
	Author     Author   `json:"author"`
}

type ArticleRequestPayload struct {
	Slug           string `pg:"slug"`
	Title          string `pg:"title"`
	Description    string `pg:"description"`
	Body           string `pg:"body"`
	Favorited      bool   `pg:"favorited"`
	FavoritesCount int    `pg:"favorites_count"`
	//Comment        map[string]interface{} `pg:"comment"`
	Score      float64  `pg:"score"`
	LikedUsers []string `pg:"liked_users"`
	AuthorId   string   `pg:"rel:has-one"`
}

type Author struct { //responsePayload
	ID         string `json:"id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Name       string `json:"name"`
	RegisterOn string `json:"register_on"`
	Image      string `json:"image"`
}

type AuthorRequestPayload struct {
	Name       string `pg:"name"`
	RegisterOn string `pg:"register_on"`
	Image      string `pg:"image"`
}

type ResponsePayload struct {
	Data interface{} `json:"data"`
}
