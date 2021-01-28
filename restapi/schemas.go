package restapi

type Tag struct {
	Tags []string `json:"tags"`
	// TagName  string   `json:"-"`
	// TagId    string   `json:"tag_id_123"`
	// TestName string   `json:"tag_name,omitempty"` //не появляется если пустое
}

type Article struct { //responsePayload
	ID             string                 `pg:"id,pk,type:uuid"` //json...
	CreatedAt      string                 `pg:"created_at"`
	UpdatedAt      string                 `pg:"updated_at"`
	Slug           string                 `pg:"slug"`
	Title          string                 `pg:"title"`
	Description    string                 `pg:"description"`
	Body           string                 `pg:"body"`
	TagList        []string               `pg:"many2many:tag_in_articles"`
	Favorited      bool                   `pg:"favorited"`
	FavoritesCount int                    `pg:"favorites_count"`
	Comment        map[string]interface{} `pg:"comment"`
	Score          float64                `pg:"score"`
	LikedUsers     []string               `pg:"liked_users"`
	Author         Author                 `pg:"rel:has-one"` //copy
}

type ArticleRequestPayload struct {
	Slug           string                 `pg:"slug"`
	Title          string                 `pg:"title"`
	Description    string                 `pg:"description"`
	Body           string                 `pg:"body"`
	Favorited      bool                   `pg:"favorited"`
	FavoritesCount int                    `pg:"favorites_count"`
	Comment        map[string]interface{} `pg:"comment"`
	Score          float64                `pg:"score"`
	LikedUsers     []string               `pg:"liked_users"`
	AuthorId       string                 `pg:"rel:has-one"` //copy
}

type Author struct {
	ID         string `pg:"id,pk,type:uuid"` //json
	CreatedAt  string `pg:"created_at"`
	UpdatedAt  string `pg:"updated_at"`
	Name       string `pg:"name"`
	RegisterOn string `pg:"register_on"`
	Image      string `pg:"image"`
}

type ResponsePayload struct {
	Data interface{} `json:"data"`
}
