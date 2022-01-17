package entity

import "time"

type Article struct {
	Id          string    `bson:"_id" json:"-"`
	Name        string    `bson:"name" json:"name"`
	Author      string    `bson:"author" json:"author"`
	Summary     string    `bson:"summary" json:"summary"`
	CoverImg    string    `bson:"cover_img" json:"coverImg"`
	Catalog     Catalog   `bson:"catalog" json:"catalog"`
	Content     string    `bson:"content" json:"content"`
	Tags        []Tag     `bson:"tags" json:"tags"`
	Kinds       []Kind    `bson:"kinds" json:"kinds"`
	ReleaseTime time.Time `bson:"release_time" json:"releaseTime"`
	Visibility  int       `bson:"visibility" json:"visibility"`
	Available   bool      `bson:"available" json:"-"`
	CreatedTime time.Time `bson:"created_time" json:"-"`
	UpdatedTime time.Time `bson:"updated_time" json:"-"`
}

type Tag struct {
	Id          string    `bson:"_id" json:"-"`
	Name        string    `bson:"name" json:"name"`
	Available   bool      `bson:"available" json:"-"`
	CreatedTime time.Time `bson:"created_time" json:"-"`
	UpdatedTime time.Time `bson:"updated_time" json:"-"`
}

type Kind struct {
	Id          string    `bson:"_id" json:"-"`
	Name        string    `bson:"name" json:"name"`
	Available   bool      `bson:"available" json:"-"`
	CreatedTime time.Time `bson:"created_time" json:"-"`
	UpdatedTime time.Time `bson:"updated_time" json:"-"`
}

type Catalog struct {
	Value string
	Child []Catalog
}
