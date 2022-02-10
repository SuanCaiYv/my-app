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
	"sync"
	"time"
)

type ArticleDao interface {
	Insert(article *entity.Article) error

	Select(id string) (*entity.Article, error)

	SelectByAuthorName(author, name string) (*entity.Article, error)

	// ListByAuthor0 未分页版本
	ListByAuthor0(author string) ([]entity.Article, error)

	ListByAuthor(author string, pgNum, pgSize int64, sort string, desc bool) ([]entity.Article, int64, error)

	Update(article *entity.Article) error

	Delete0(id string) error

	Delete(id string) error
}

type ArticleDaoService struct {
	collection *mongo.Collection
	logger     *logrus.Logger
}

var (
	instanceArticleDaoService *ArticleDaoService
	onceArticleDaoService     sync.Once
)

func NewArticleDaoService() *ArticleDaoService {
	onceArticleDaoService.Do(newInstanceArticleDaoService)
	return instanceArticleDaoService
}

func newInstanceArticleDaoService() {
	logger := util.NewLogger()
	config := config2.ApplicationConfiguration()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	url := fmt.Sprintf("%s:%d", config.DatabaseConfig.Url, config.DatabaseConfig.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	util.JustPanic(err)
	collection := client.Database(config.DatabaseConfig.DB).Collection(CollectionArticle)
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
	one := a.collection.FindOne(timeout, primitive.M{"author": author, "name": name})
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

func (a *ArticleDaoService) ListByAuthor0(author string) ([]entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleDaoService) ListByAuthor(author string, pgNum, pgSize int64, sort string, desc bool) ([]entity.Article, int64, error) {
	descInt := 1
	if desc {
		descInt = -1
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	skip := (pgNum - 1) * pgNum
	cursor, err := a.collection.Find(ctx, primitive.M{"author": author}, &options.FindOptions{
		Limit: &pgSize,
		Skip:  &skip,
		Sort:  primitive.M{sort: descInt},
	})
	if err != nil {
		a.logger.Error(err)
		return nil, 0, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			a.logger.Error(err)
		}
	}(cursor, ctx)
	if err != nil {
		return nil, 0, err
	}
	results := make([]entity.Article, 0, pgSize)
	for cursor.Next(ctx) {
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
	total, err := a.collection.CountDocuments(ctx, primitive.M{"author": author})
	if err != nil {
		a.logger.Error(err)
		return nil, 0, err
	}
	return results, total, nil
}

func (a *ArticleDaoService) Update(article *entity.Article) error {
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
