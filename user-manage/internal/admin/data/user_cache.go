package data

import (
	"fmt"

	"github.com/EchoGroot/kratos-examples/pkg/cache"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type UserCacheRepo struct {
	cache.CacheRepo[User]
}

func NewUserCacheRepo(db *gorm.DB, rds *redis.ClusterClient) *UserCacheRepo {
	userCacheRepo := &UserCacheRepo{}
	userCacheRepo.CacheRepo = cache.NewCacheRepo[User](db, rds, userCacheRepo)
	return userCacheRepo
}

func (r *UserRepo) PkCacheKey(pk interface{}) string {
	return fmt.Sprintf("user-manage:cache:%s:%s:%v", User{}.TableName(), r.PrimaryKey, pk)
}

func (r *UserRepo) GetPk(t interface{}) (interface{}, error) {
	a, ok := t.(*User)
	if ok {
		return a.Id, nil
	}
	return nil, errors.Errorf("%T not a user", t)
}

func (r *UserRepo) IndexCacheKeys() []cache.IndexCacheKeyProvider {
	res := make([]cache.IndexCacheKeyProvider, 1)
	res[0] = UsernameProvider{}
	return res
}

type UsernameProvider struct{}

func (r UsernameProvider) FieldName() string {
	return "username"
}

func (r UsernameProvider) IndexCacheKeyFromStruct(s interface{}) string {
	return r.IndexCacheKey(s.(*User).Username)
}

func (r UsernameProvider) IndexCacheKey(index string) string {
	return fmt.Sprintf("user-manage:cache:%s:%s:%s", User{}.TableName(), "username", index)
}
