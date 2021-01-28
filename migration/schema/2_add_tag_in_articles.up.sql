create table tag_in_articles (
	id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    tag_id UUID REFERENCES tags(id) ON DELETE cascade,
    article_id UUID REFERENCES articles(id) ON DELETE CASCADE
);