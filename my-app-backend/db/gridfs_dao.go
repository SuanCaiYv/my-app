package db

import (
	context2 "context"
	"fmt"
	config2 "github.com/SuanCaiYv/my-app-backend/config"
	"github.com/SuanCaiYv/my-app-backend/util"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"sync"
	"time"
)

type GridFSDao interface {
	UploadFile(fileContent []byte, filename string, metaData primitive.M) error

	ModifyFile(file os.File) error

	DownloadFile(filename string) ([]byte, primitive.M, error)

	ListByArchive(archive string, pgNum, pgSize int64) ([]string, int64, error)

	ListByArchive0(archive string) ([]string, error)

	DeleteFile(filename string) error

	ExistFile(filename string) bool
}

type GridFSDaoService struct {
	bucket *gridfs.Bucket
	logger *logrus.Logger
}

var (
	instanceGridFSDaoService *GridFSDaoService
	onceGridFSDaoService     sync.Once
)

func NewGridFSDaoService() *GridFSDaoService {
	onceGridFSDaoService.Do(newInstanceGridFSDaoService)
	return instanceGridFSDaoService
}

func newInstanceGridFSDaoService() {
	logger := util.NewLogger()
	config := config2.ApplicationConfiguration()
	ctx, cancel := context2.WithTimeout(context2.Background(), 2*time.Second)
	defer cancel()
	url := fmt.Sprintf("%s:%d", config.DatabaseConfig.Url, config.DatabaseConfig.Port)
	clientOptions := options.Client()
	clientOptions.ApplyURI(url)
	if len(config.DatabaseConfig.User) != 0 {
		clientOptions.Auth = &options.Credential{
			AuthSource: config.DatabaseConfig.GridFSDB,
			Username:   config.DatabaseConfig.User,
			Password:   config.DatabaseConfig.Password,
		}
	}
	client, err := mongo.Connect(ctx, clientOptions)
	util.JustPanic(err)
	bucket, err := gridfs.NewBucket(client.Database(config.DatabaseConfig.GridFSDB))
	util.JustPanic(err)
	instanceGridFSDaoService = &GridFSDaoService{
		bucket,
		logger,
	}
}

func (g *GridFSDaoService) UploadFile(fileContent []byte, filename string, metaData primitive.M) error {
	// ????????????????????????
	option := options.GridFSUpload()
	option.SetMetadata(metaData)
	// ??????????????????????????????Insert Files???????????????option????????????Files???meta?????????????????????????????????Files?????????????????????
	// Files??????????????????????????????chunks???
	uploadStream, err := g.bucket.OpenUploadStream(filename, option)
	defer func(uploadStream *gridfs.UploadStream) {
		_ = uploadStream.Close()
	}(uploadStream)
	if err != nil {
		g.logger.Errorf("??????GridFS???????????????: %v", err)
		return err
	}
	_, err = uploadStream.Write(fileContent)
	if err != nil {
		g.logger.Errorf("???????????????GridFS??????: %v", err)
		return err
	}
	return nil
}

func (g *GridFSDaoService) ModifyFile(file os.File) error {
	//TODO implement me
	panic("implement me")
}

func (g *GridFSDaoService) DownloadFile(filename string) ([]byte, primitive.M, error) {
	ctx, cancel := context2.WithTimeout(context2.Background(), 2*time.Second)
	defer cancel()
	// ??????????????????????????????Files???????????????????????????????????????????????????????????????
	cursor, err := g.bucket.GetFilesCollection().Find(ctx, bson.M{"filename": filename})
	if err != nil {
		g.logger.Errorf("??????files??????????????????: %s", filename)
		return nil, nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context2.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			g.logger.Error(err)
		}
	}(cursor, ctx)
	var gFile *gridfs.File = nil
	for cursor.Next(ctx) {
		file := gridfs.File{}
		err := cursor.Decode(&file)
		if err != nil {
			return nil, nil, err
		}
		gFile = &file
	}
	if gFile == nil {
		return nil, nil, nil
	}
	// ????????????Id??????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????
	downloadStream, err := g.bucket.OpenDownloadStream(gFile.ID)
	defer func(downloadStream *gridfs.DownloadStream) {
		_ = downloadStream.Close()
	}(downloadStream)
	if err != nil {
		g.logger.Errorf("?????????????????????: %v", err)
		return nil, nil, err
	}
	size := downloadStream.GetFile().Length
	data := make([]byte, size, size)
	// ?????????????????????
	_, err = downloadStream.Read(data)
	if err != nil {
		g.logger.Errorf("??????????????????: %v", err)
		return nil, nil, err
	}
	meta := gFile.Metadata
	meteData := make(map[string]interface{})
	err = bson.Unmarshal(meta, &meteData)
	if err != nil {
		g.logger.Error(err)
		return nil, nil, err
	}
	return data, meteData, nil
}

func (g *GridFSDaoService) ListByArchive(archive string, pgNum, pgSize int64) ([]string, int64, error) {
	skip := (pgNum - 1) * pgNum
	ctx, cancel := context2.WithTimeout(context2.Background(), 5*time.Second)
	defer cancel()
	cursor, err := g.bucket.GetFilesCollection().Find(ctx,
		primitive.M{
			"metadata.archive": archive,
		},
		&options.FindOptions{
			Limit: &pgSize,
			Skip:  &skip,
			Sort:  primitive.M{"uploadDate": -1},
		},
	)
	if err != nil {
		g.logger.Errorf("????????????????????????????????????: %s", archive)
		return nil, 0, err
	}
	defer func(cursor *mongo.Cursor, ctx context2.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			g.logger.Error(err)
		}
	}(cursor, ctx)
	results := make([]string, 0, pgNum)
	for cursor.Next(ctx) {
		file := gridfs.File{}
		err := cursor.Decode(&file)
		if err != nil {
			return nil, 0, err
		}
		results = append(results, file.Name)
	}
	total, err := g.bucket.GetFilesCollection().CountDocuments(ctx, bson.M{"metadata": primitive.M{"archive": archive}})
	if err != nil {
		g.logger.Error(err)
		return nil, 0, err
	}
	return results, total, nil
}

func (g *GridFSDaoService) ListByArchive0(archive string) ([]string, error) {
	ctx, cancel := context2.WithTimeout(context2.Background(), 5*time.Second)
	defer cancel()
	cursor, err := g.bucket.GetFilesCollection().Find(ctx,
		primitive.M{
			"metadata.archive": archive,
		},
	)
	if err != nil {
		g.logger.Errorf("????????????????????????????????????: %s", archive)
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context2.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			g.logger.Error(err)
		}
	}(cursor, ctx)
	results := make([]string, 0, 10)
	for cursor.Next(ctx) {
		file := gridfs.File{}
		err := cursor.Decode(&file)
		if err != nil {
			return nil, err
		}
		results = append(results, file.Name)
	}
	return results, nil
}

func (g *GridFSDaoService) DeleteFile(filename string) error {
	ctx, cancel := context2.WithTimeout(context2.Background(), 5*time.Second)
	defer cancel()
	cursor, err := g.bucket.GetFilesCollection().Find(ctx, bson.M{"filename": filename})
	if err != nil {
		g.logger.Errorf("??????files??????????????????: %s", filename)
		return err
	}
	defer func(cursor *mongo.Cursor, ctx context2.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			g.logger.Error(err)
		}
	}(cursor, ctx)
	var gFile *gridfs.File = nil
	for cursor.Next(ctx) {
		file := gridfs.File{}
		err := cursor.Decode(&file)
		if err != nil {
			return err
		}
		gFile = &file
	}
	if gFile != nil {
		err = g.bucket.Delete(gFile.ID)
		if err != nil {
			g.logger.Errorf("??????chunks??????????????????: %s", filename)
			return err
		}
	}
	return nil
}

func (g *GridFSDaoService) ExistFile(filename string) bool {
	ctx, cancel := context2.WithTimeout(context2.Background(), 5*time.Second)
	defer cancel()
	cursor, err := g.bucket.GetFilesCollection().Find(ctx, bson.M{"filename": filename})
	if err != nil {
		g.logger.Errorf("??????files??????????????????: %s", filename)
		return false
	}
	defer func(cursor *mongo.Cursor, ctx context2.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			g.logger.Error(err)
		}
	}(cursor, ctx)
	var gFile *gridfs.File = nil
	for cursor.Next(ctx) {
		file := gridfs.File{}
		err := cursor.Decode(&file)
		if err != nil {
			return false
		}
		gFile = &file
	}
	return gFile != nil
}
