package db

import (
	"context"
	"fmt"
	config2 "github.com/SuanCaiYv/my-app-backend/config"
	"github.com/SuanCaiYv/my-app-backend/entity"
	"github.com/SuanCaiYv/my-app-backend/util"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"sync"
	"time"
)

const (
	PlaceHolderId = "000000000000000000000001"
)

type ArticleDao interface {
	Insert(article *entity.Article) error

	Select(id string) (*entity.Article, error)

	SelectByAuthorName(author, name string) (*entity.Article, error)

	// ListByAuthor0 未分页版本
	ListByAuthor0(author string, visibility int, equally bool) ([]entity.Article, error)

	ListByAuthor(author string, visibility int, equally bool, pgNum, pgSize int64, sort string, desc bool, tagIdList []string, searchKey string) ([]entity.Article, int64, error)

	Update(article *entity.Article) error

	Delete0(id string) error

	Delete(id string) error
}

type KindDao interface {
	Insert(kind *entity.Kind) error

	Select(id string) (*entity.Kind, error)

	SelectByName(name string) (*entity.Kind, error)

	Update(kind *entity.Kind) error

	Delete(id string) error

	ListAll() ([]entity.Kind, error)
}

type TagDao interface {
	Insert(tag *entity.Tag) error

	Select(id string) (*entity.Tag, error)

	SelectByName(name string) (*entity.Tag, error)

	Update(tag *entity.Tag) error

	Delete(id string) error

	ListAll() ([]entity.Tag, error)
}

type ArticleDaoService struct {
	collection *mongo.Collection
	logger     *logrus.Logger
}

type KindDaoService struct {
	collection *mongo.Collection
	logger     *logrus.Logger
}

type TagDaoService struct {
	collection *mongo.Collection
	logger     *logrus.Logger
}

var (
	instanceArticleDaoService *ArticleDaoService
	onceArticleDaoService     sync.Once
	instanceKindDaoService    *KindDaoService
	onceKindDaoService        sync.Once
	instanceTagDaoService     *TagDaoService
	onceTagDaoService         sync.Once
)

func NewArticleDaoService() *ArticleDaoService {
	onceArticleDaoService.Do(newInstanceArticleDaoService)
	return instanceArticleDaoService
}

func newInstanceArticleDaoService() {
	logger := util.NewLogger()
	config := config2.ApplicationConfiguration()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	url := fmt.Sprintf("%s:%d", config.DatabaseConfig.Url, config.DatabaseConfig.Port)
	clientOptions := options.Client()
	clientOptions.ApplyURI(url)
	if len(config.DatabaseConfig.User) != 0 {
		clientOptions.Auth = &options.Credential{
			AuthSource: config.DatabaseConfig.DB,
			Username:   config.DatabaseConfig.User,
			Password:   config.DatabaseConfig.Password,
		}
	}
	client, err := mongo.Connect(ctx, clientOptions)
	util.JustPanic(err)
	collection := client.Database(config.DatabaseConfig.DB).Collection(CollectionArticle)
	one := collection.FindOne(ctx, primitive.M{"_id": PlaceHolderId})
	language := "none"
	if one.Err() == mongo.ErrNoDocuments {
		_, err = collection.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys: bsonx.Doc{
				{Key: "fulltext_title", Value: bsonx.String("text")},
				{Key: "fulltext_content", Value: bsonx.String("text")},
			},
			Options: &options.IndexOptions{
				DefaultLanguage: &language,
				Weights: struct {
					Weights int32 `bson:"weights"`
				}{Weights: 1},
			},
		})
		util.JustPanic(err)
		_, err = collection.InsertOne(ctx, &entity.Article{
			Id:       PlaceHolderId,
			Name:     "Welcome to My Blog",
			Author:   "program",
			Summary:  "",
			CoverImg: "",
			Catalog: entity.Catalog{
				Name:     "",
				Children: []entity.Catalog{},
			},
			Content:         "Welcome to My Blog, this is the first article, you can edit it by yourself.",
			Kind:            entity.Kind{},
			TagList:         []entity.Tag{},
			ReleaseTime:     time.Now(),
			Visibility:      0,
			FulltextTitle:   "",
			FulltextContent: "",
			Available:       true,
			CreatedTime:     time.Now(),
			UpdatedTime:     time.Now(),
		})
		util.JustPanic(err)
	}
	instanceArticleDaoService = &ArticleDaoService{
		collection,
		logger,
	}
}

func (a *ArticleDaoService) Insert(article *entity.Article) error {
	article.Available = true
	article.CreatedTime = time.Now()
	article.UpdatedTime = time.Now()
	article.ReleaseTime = time.Now()
	article.Id = primitive.NewObjectID().Hex()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := a.collection.InsertOne(ctx, article)
	return err
}

func (a *ArticleDaoService) Select(id string) (*entity.Article, error) {
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	one := a.collection.FindOne(timeout, primitive.M{"_id": id})
	if err := one.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		a.logger.Error(err)
		return nil, err
	}
	result := entity.Article{}
	err := one.Decode(&result)
	if err != nil {
		a.logger.Error(err)
		return nil, err
	}
	return &result, nil
}

func (a *ArticleDaoService) SelectByAuthorName(author, name string) (*entity.Article, error) {
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	one := a.collection.FindOne(timeout, primitive.M{"author": author, "name": name, "available": true})
	if err := one.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		a.logger.Error(err)
		return nil, err
	}
	result := entity.Article{}
	err := one.Decode(&result)
	if err != nil {
		a.logger.Error(err)
		return nil, err
	}
	return &result, nil
}

func (a *ArticleDaoService) ListByAuthor0(author string, visibility int, equally bool) ([]entity.Article, error) {
	var v interface{}
	if equally {
		v = visibility
	} else {
		v = primitive.M{"$ne": visibility}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	cursor, err := a.collection.Find(ctx,
		primitive.M{
			"author":     author,
			"available":  true,
			"visibility": v,
		})
	if err != nil {
		a.logger.Error(err)
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			a.logger.Error(err)
		}
	}(cursor, ctx)
	if err != nil {
		return nil, err
	}
	results := make([]entity.Article, 0, 10)
	for cursor.Next(ctx) {
		tmp := entity.Article{}
		err := cursor.Decode(&tmp)
		if err != nil {
			a.logger.Error(err)
			return nil, err
		}
		results = append(results, tmp)
	}
	if err := cursor.Err(); err != nil {
		a.logger.Error(err)
		return nil, err
	}
	return results, nil
}

func (a *ArticleDaoService) ListByAuthor(author string, visibility int, equally bool, pgNum, pgSize int64, sort string, desc bool, tagIdList []string, searchKey string) ([]entity.Article, int64, error) {
	descInt := -1
	if desc {
		descInt = 1
	}
	skip := (pgNum - 1) * pgNum
	var v interface{}
	if equally {
		v = visibility
	} else {
		v = primitive.M{"$ne": visibility}
	}
	timeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var cursor *mongo.Cursor
	var err error
	var total int64
	if len(tagIdList) != 0 && searchKey != "" {
		cursor, err = a.collection.Find(timeout,
			primitive.M{
				"author":     author,
				"available":  true,
				"visibility": v,
				"tag_list._id": primitive.M{
					"$all": tagIdList,
				},
				"$text": primitive.M{
					"$search": searchKey,
				},
			},
			&options.FindOptions{
				Limit: &pgSize,
				Skip:  &skip,
				Sort:  primitive.M{sort: descInt},
			},
		)
		total, err = a.collection.CountDocuments(timeout, primitive.M{
			"author":     author,
			"available":  true,
			"visibility": v,
			"tag_list._id": primitive.M{
				"$all": tagIdList,
			},
			"$text": primitive.M{
				"$search": searchKey,
			},
		})
		if err != nil {
			a.logger.Error(err)
			return nil, 0, err
		}
	} else if len(tagIdList) != 0 {
		cursor, err = a.collection.Find(timeout,
			primitive.M{
				"author":     author,
				"available":  true,
				"visibility": v,
				"tag_list._id": primitive.M{
					"$all": tagIdList,
				},
			},
			&options.FindOptions{
				Limit: &pgSize,
				Skip:  &skip,
				Sort:  primitive.M{sort: descInt},
			},
		)
		total, err = a.collection.CountDocuments(timeout, primitive.M{
			"author":     author,
			"available":  true,
			"visibility": v,
			"tag_list._id": primitive.M{
				"$all": tagIdList,
			},
		})
		if err != nil {
			a.logger.Error(err)
			return nil, 0, err
		}
	} else if searchKey != "" {
		cursor, err = a.collection.Find(timeout,
			primitive.M{
				"author":     author,
				"available":  true,
				"visibility": v,
				"$text": primitive.M{
					"$search": searchKey,
				},
			},
			&options.FindOptions{
				Limit: &pgSize,
				Skip:  &skip,
				Sort:  primitive.M{sort: descInt},
			},
		)
		total, err = a.collection.CountDocuments(timeout, primitive.M{
			"author":     author,
			"available":  true,
			"visibility": v,
			"$text": primitive.M{
				"$search": searchKey,
			},
		})
		if err != nil {
			a.logger.Error(err)
			return nil, 0, err
		}
	} else {
		cursor, err = a.collection.Find(timeout,
			primitive.M{
				"author":     author,
				"available":  true,
				"visibility": v,
			},
			&options.FindOptions{
				Limit: &pgSize,
				Skip:  &skip,
				Sort:  primitive.M{sort: descInt},
			},
		)
		total, err = a.collection.CountDocuments(timeout, primitive.M{
			"author":     author,
			"available":  true,
			"visibility": v,
		})
		if err != nil {
			a.logger.Error(err)
			return nil, 0, err
		}
	}
	if err != nil {
		a.logger.Error(err)
		return nil, 0, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			a.logger.Error(err)
		}
	}(cursor, timeout)
	if err != nil {
		return nil, 0, err
	}
	results := make([]entity.Article, 0, pgSize)
	for cursor.Next(timeout) {
		tmp := entity.Article{}
		err := cursor.Decode(&tmp)
		if err != nil {
			a.logger.Error(err)
			return nil, 0, err
		}
		results = append(results, tmp)
	}
	if err := cursor.Err(); err != nil {
		a.logger.Error(err)
		return nil, 0, err
	}
	return results, total, nil
}

func (a *ArticleDaoService) Update(article *entity.Article) error {
	article.UpdatedTime = time.Now()
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	_, err := a.collection.UpdateByID(timeout, article.Id, primitive.M{"$set": article})
	return err
}

func (a *ArticleDaoService) Delete0(id string) error {
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	_, err := a.collection.DeleteOne(timeout, primitive.M{"_id": id})
	return err
}

func (a *ArticleDaoService) Delete(id string) error {
	article, err := a.Select(id)
	if err != nil {
		return err
	}
	article.UpdatedTime = time.Now()
	article.Available = false
	return a.Update(article)
}

func NewKindDaoService() *KindDaoService {
	onceKindDaoService.Do(newInstanceKindDaoService)
	return instanceKindDaoService
}

func newInstanceKindDaoService() {
	logger := util.NewLogger()
	config := config2.ApplicationConfiguration()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	url := fmt.Sprintf("%s:%d", config.DatabaseConfig.Url, config.DatabaseConfig.Port)
	clientOptions := options.Client()
	clientOptions.ApplyURI(url)
	if len(config.DatabaseConfig.User) != 0 {
		clientOptions.Auth = &options.Credential{
			AuthSource: config.DatabaseConfig.DB,
			Username:   config.DatabaseConfig.User,
			Password:   config.DatabaseConfig.Password,
		}
	}
	client, err := mongo.Connect(ctx, clientOptions)
	util.JustPanic(err)
	collection := client.Database(config.DatabaseConfig.DB).Collection(CollectionKind)
	instanceKindDaoService = &KindDaoService{
		collection,
		logger,
	}
}

func (k *KindDaoService) Insert(kind *entity.Kind) error {
	kind.Available = true
	kind.CreatedTime = time.Now()
	kind.UpdatedTime = time.Now()
	kind.Id = primitive.NewObjectID().Hex()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := k.collection.InsertOne(ctx, kind)
	return err
}

func (k *KindDaoService) Select(id string) (*entity.Kind, error) {
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	one := k.collection.FindOne(timeout, primitive.M{"_id": id, "available": true})
	if err := one.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		k.logger.Error(err)
		return nil, err
	}
	result := entity.Kind{}
	err := one.Decode(&result)
	if err != nil {
		k.logger.Error(err)
		return nil, err
	}
	return &result, nil
}

func (k *KindDaoService) SelectByName(name string) (*entity.Kind, error) {
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	one := k.collection.FindOne(timeout, primitive.M{"name": name, "available": true})
	if err := one.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		k.logger.Error(err)
		return nil, err
	}
	result := entity.Kind{}
	err := one.Decode(&result)
	if err != nil {
		k.logger.Error(err)
		return nil, err
	}
	return &result, nil
}

func (k *KindDaoService) Update(kind *entity.Kind) error {
	kind.UpdatedTime = time.Now()
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	_, err := k.collection.UpdateByID(timeout, kind.Id, primitive.M{"$set": kind})
	return err
}

func (k *KindDaoService) Delete(id string) error {
	kind, err := k.Select(id)
	if err != nil {
		return err
	}
	kind.UpdatedTime = time.Now()
	kind.Available = false
	return k.Update(kind)
}

func (k *KindDaoService) ListAll() ([]entity.Kind, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	cursor, err := k.collection.Find(ctx, primitive.M{"available": true})
	if err != nil {
		k.logger.Error(err)
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			k.logger.Error(err)
		}
	}(cursor, ctx)
	if err != nil {
		return nil, err
	}
	results := make([]entity.Kind, 0, 10)
	for cursor.Next(ctx) {
		tmp := entity.Kind{}
		err := cursor.Decode(&tmp)
		if err != nil {
			k.logger.Error(err)
			return nil, err
		}
		results = append(results, tmp)
	}
	if err := cursor.Err(); err != nil {
		k.logger.Error(err)
		return nil, err
	}
	return results, nil
}

func NewTagDaoService() *TagDaoService {
	onceTagDaoService.Do(newInstanceTagDaoService)
	return instanceTagDaoService
}

func newInstanceTagDaoService() {
	logger := util.NewLogger()
	config := config2.ApplicationConfiguration()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	url := fmt.Sprintf("%s:%d", config.DatabaseConfig.Url, config.DatabaseConfig.Port)
	clientOptions := options.Client()
	clientOptions.ApplyURI(url)
	if len(config.DatabaseConfig.User) != 0 {
		clientOptions.Auth = &options.Credential{
			AuthSource: config.DatabaseConfig.DB,
			Username:   config.DatabaseConfig.User,
			Password:   config.DatabaseConfig.Password,
		}
	}
	client, err := mongo.Connect(ctx, clientOptions)
	util.JustPanic(err)
	collection := client.Database(config.DatabaseConfig.DB).Collection(CollectionTag)
	instanceTagDaoService = &TagDaoService{
		collection,
		logger,
	}
}

func (t *TagDaoService) Insert(tag *entity.Tag) error {
	tag.Available = true
	tag.CreatedTime = time.Now()
	tag.UpdatedTime = time.Now()
	tag.Id = primitive.NewObjectID().Hex()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := t.collection.InsertOne(ctx, tag)
	return err
}

func (t *TagDaoService) Select(id string) (*entity.Tag, error) {
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	one := t.collection.FindOne(timeout, primitive.M{"_id": id, "available": true})
	if err := one.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		t.logger.Error(err)
		return nil, err
	}
	result := entity.Tag{}
	err := one.Decode(&result)
	if err != nil {
		t.logger.Error(err)
		return nil, err
	}
	return &result, nil
}

func (t *TagDaoService) SelectByName(name string) (*entity.Tag, error) {
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	one := t.collection.FindOne(timeout, primitive.M{"name": name, "available": true})
	if err := one.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		t.logger.Error(err)
		return nil, err
	}
	result := entity.Tag{}
	err := one.Decode(&result)
	if err != nil {
		t.logger.Error(err)
		return nil, err
	}
	return &result, nil
}

func (t *TagDaoService) Update(tag *entity.Tag) error {
	tag.UpdatedTime = time.Now()
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	_, err := t.collection.UpdateByID(timeout, tag.Id, primitive.M{"$set": tag})
	return err
}

func (t *TagDaoService) Delete(id string) error {
	tag, err := t.Select(id)
	if err != nil {
		return err
	}
	tag.UpdatedTime = time.Now()
	tag.Available = false
	return t.Update(tag)
}

func (t *TagDaoService) ListAll() ([]entity.Tag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	cursor, err := t.collection.Find(ctx, primitive.M{"available": true})
	if err != nil {
		t.logger.Error(err)
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			t.logger.Error(err)
		}
	}(cursor, ctx)
	if err != nil {
		return nil, err
	}
	results := make([]entity.Tag, 0, 10)
	for cursor.Next(ctx) {
		tmp := entity.Tag{}
		err := cursor.Decode(&tmp)
		if err != nil {
			t.logger.Error(err)
			return nil, err
		}
		results = append(results, tmp)
	}
	if err := cursor.Err(); err != nil {
		t.logger.Error(err)
		return nil, err
	}
	return results, nil
}
