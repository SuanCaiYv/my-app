package db

import (
	context2 "context"
	"fmt"
	config2 "github.com/SuanCaiYv/my-app-backend/config"
	"github.com/SuanCaiYv/my-app-backend/entity"
	"github.com/SuanCaiYv/my-app-backend/util"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"sync"
	"time"
)

type SysUserDao interface {
	Insert(sysUser *entity.SysUser) error

	Select(id string) (*entity.SysUser, error)

	Update(sysUser *entity.SysUser) error

	Delete0(id string) error

	Delete(id string) error

	SelectByUsername(username string) (*entity.SysUser, error)

	SelectByNickname(nickname string) (*entity.SysUser, error)
}

type SysRoleDao interface {
	Insert(sysRole *entity.SysRole) error

	Select(id string) (*entity.SysRole, error)

	SelectByName(name string) (*entity.SysRole, error)
}

type SysUserDaoService struct {
	collection *mongo.Collection
	logger     *logrus.Logger
}

type SysRoleDaoService struct {
	collection *mongo.Collection
	logger     *logrus.Logger
}

var (
	instanceSysUserDaoService *SysUserDaoService
	instanceSysRoleDaoService *SysRoleDaoService
	onceSysUserDaoService     sync.Once
	onceSysRoleDaoService     sync.Once
)

func NewSysUserDaoService() *SysUserDaoService {
	onceSysUserDaoService.Do(newInstanceSysUserDaoService)
	return instanceSysUserDaoService
}

func newInstanceSysUserDaoService() {
	logger := util.NewLogger()
	config := config2.ApplicationConfiguration()
	ctx, cancel := context2.WithTimeout(context2.Background(), 2*time.Second)
	defer cancel()
	url := fmt.Sprintf("%s:%d", config.DatabaseConfig.Url, config.DatabaseConfig.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	util.JustPanic(err)
	collection := client.Database(config.DatabaseConfig.DB).Collection(CollectionSysUser)
	instanceSysUserDaoService = &SysUserDaoService{
		collection,
		logger,
	}
}

func (s *SysUserDaoService) Insert(sysUser *entity.SysUser) error {
	sysUser.Id = primitive.NewObjectID().Hex()
	sysUser.Available = true
	sysUser.CreatedTime = time.Now()
	sysUser.UpdatedTime = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := s.collection.InsertOne(ctx, sysUser)
	return err
}

// Select 不同于该有的做法，这里当Record为空时，返回nil, nil，而不是nil, ErrNotFound
// 返回error只有在真的发生了error时才会返回。
func (s *SysUserDaoService) Select(id string) (*entity.SysUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	one := s.collection.FindOne(ctx, primitive.M{"_id": id})
	if err := one.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		s.logger.Error(err)
		return nil, err
	}
	result := entity.SysUser{}
	err := one.Decode(&result)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return &result, nil
}

func (s *SysUserDaoService) Update(sysUser *entity.SysUser) error {
	sysUser.UpdatedTime = time.Now()
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	_, err := s.collection.UpdateByID(timeout, sysUser.Id, primitive.M{"$set": sysUser})
	return err
}

func (s *SysUserDaoService) Delete0(id string) error {
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	_, err := s.collection.DeleteOne(timeout, primitive.M{"_id": id})
	return err
}

func (s *SysUserDaoService) Delete(id string) error {
	sysUser, err := s.Select(id)
	if err != nil {
		return err
	}
	sysUser.Available = false
	sysUser.UpdatedTime = time.Now()
	err = s.Update(sysUser)
	return err
}

func (s *SysUserDaoService) SelectByUsername(username string) (*entity.SysUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	one := s.collection.FindOne(ctx, primitive.M{"username": username})
	if err := one.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		s.logger.Error(err)
		return nil, err
	}
	result := entity.SysUser{}
	err := one.Decode(&result)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return &result, nil
}

func (s *SysUserDaoService) SelectByNickname(nickname string) (*entity.SysUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	one := s.collection.FindOne(ctx, primitive.M{"nickname": nickname})
	if err := one.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		s.logger.Error(err)
		return nil, err
	}
	result := entity.SysUser{}
	err := one.Decode(&result)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return &result, nil
}

func NewSysRoleDaoService() *SysRoleDaoService {
	onceSysRoleDaoService.Do(newInstanceSysRoleDaoService)
	return instanceSysRoleDaoService
}

func newInstanceSysRoleDaoService() {
	logger := util.NewLogger()
	config := config2.ApplicationConfiguration()
	ctx, cancel := context2.WithTimeout(context2.Background(), 2*time.Second)
	defer cancel()
	url := fmt.Sprintf("%s:%d", config.DatabaseConfig.Url, config.DatabaseConfig.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	util.JustPanic(err)
	collection := client.Database(config.DatabaseConfig.DB).Collection(CollectionSysRole)
	instanceSysRoleDaoService = &SysRoleDaoService{
		collection,
		logger,
	}
}

func (s *SysRoleDaoService) Insert(sysRole *entity.SysRole) error {
	sysRole.Id = primitive.NewObjectID().Hex()
	sysRole.Available = true
	sysRole.CreatedTime = time.Now()
	sysRole.UpdatedTime = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := s.collection.InsertOne(ctx, sysRole)
	return err
}

func (s *SysRoleDaoService) Select(id string) (*entity.SysRole, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SysRoleDaoService) SelectByName(name string) (*entity.SysRole, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	one := s.collection.FindOne(ctx, primitive.M{"name": name})
	if err := one.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		s.logger.Error(err)
		return nil, err
	}
	result := entity.SysRole{}
	err := one.Decode(&result)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return &result, nil
}
