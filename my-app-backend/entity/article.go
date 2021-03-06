package entity

import "time"

type Article struct {
	Id              string    `bson:"_id" json:"article_id"`
	Name            string    `bson:"name" json:"article_name"`
	Author          string    `bson:"author" json:"author"`
	Summary         string    `bson:"summary" json:"summary"`
	CoverImg        string    `bson:"cover_img" json:"cover_img"`
	Catalog         Catalog   `bson:"catalog" json:"catalog"`
	Content         string    `bson:"content" json:"content"`
	Kind            Kind      `bson:"kind" json:"kind"`
	TagList         []Tag     `bson:"tag_list" json:"tag_list"`
	ReleaseTime     time.Time `bson:"release_time" json:"release_time"`
	Visibility      int       `bson:"visibility" json:"visibility"`
	FulltextTitle   string    `bson:"fulltext_title" json:"-"`
	FulltextContent string    `bson:"fulltext_content" json:"-"`
	Available       bool      `bson:"available" json:"-"`
	CreatedTime     time.Time `bson:"created_time" json:"-"`
	UpdatedTime     time.Time `bson:"updated_time" json:"-"`
}

type Tag struct {
	Id          string    `bson:"_id" json:"tag_id"`
	Name        string    `bson:"name" json:"tag_name"`
	Available   bool      `bson:"available" json:"-"`
	CreatedTime time.Time `bson:"created_time" json:"-"`
	UpdatedTime time.Time `bson:"updated_time" json:"-"`
}

type Kind struct {
	Id          string    `bson:"_id" json:"kind_id"`
	Name        string    `bson:"name" json:"kind_name"`
	Available   bool      `bson:"available" json:"-"`
	CreatedTime time.Time `bson:"created_time" json:"-"`
	UpdatedTime time.Time `bson:"updated_time" json:"-"`
}

type Catalog struct {
	Name     string    `bson:"name" json:"name"`
	Children []Catalog `bson:"children" json:"children"`
}

const (
	_ = iota
	VisibilityPrivate
	VisibilityPublic
	VisibilityDraft
)
