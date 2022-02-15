package entity

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	article := Article{
		Id:       "",
		Name:     "",
		Author:   "",
		Summary:  "",
		CoverImg: "",
		Catalog: Catalog{
			Name:  "",
			Child: []Catalog{},
		},
		Content:     "",
		Tags:        []Tag{},
		Kinds:       []Kind{},
		ReleaseTime: time.Time{},
		Visibility:  0,
		Available:   false,
		CreatedTime: time.Time{},
		UpdatedTime: time.Time{},
	}
	bytes, _ := json.Marshal(article)
	fmt.Println(string(bytes))
}
