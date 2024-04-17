package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/EchoGroot/kratos-examples/pkg/gormx/repo"
	"github.com/EchoGroot/kratos-examples/pkg/lang/reflectx"
	"github.com/EchoGroot/kratos-examples/pkg/lang/stringx"
	"github.com/EchoGroot/kratos-examples/pkg/lang/structx/copier"
	"reflect"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/samber/lo"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/go-redis/redis/v8"
)

// CacheRepo 具备缓存功能的repo
// 采用 Cache Aside 缓存模式
// 参考 go-zero https://go-zero.dev/cn/docs/blog/cache/redis-cache
type CacheRepo[T any] struct {
	CacheKey
	repo.BaseRepo[T]
	rds *redis.ClusterClient
}

type (
	// queryFn 数据库查询
	queryFn func(res interface{}) (err error)
	// indexQueryFn 根据索引查询
	indexQueryFn func() (pk interface{}, err error)
	// execFn 数据库执行，如insert、update、delete
	execFn func() (res int64, err error)
)

// CacheKey 构建缓存Key
type CacheKey interface {
	// PkCacheKey 主键的缓存Key
	PkCacheKey(pk interface{}) string
	// GetPk 从实体中获取主键值
	GetPk(t interface{}) (interface{}, error)
	// IndexCacheKeys 索引的缓存Key，没有使用索引查询的，返回nil
	IndexCacheKeys() []IndexCacheKeyProvider
}

// IndexCacheKeyProvider 索引的缓存Key构造器
type IndexCacheKeyProvider interface {
	// FieldName 索引字段名
	FieldName() string
	// IndexCacheKey 根据索引值，构建索引的缓存Key
	IndexCacheKey(index string) string
	// IndexCacheKeyFromStruct 根据结构体，构建索引的缓存Key
	IndexCacheKeyFromStruct(s interface{}) string
}

func NewCacheRepo[T any](db *gorm.DB, rds *redis.ClusterClient, cacheKey CacheKey) CacheRepo[T] {
	return CacheRepo[T]{
		BaseRepo: repo.NewBaseRepo[T](db),
		rds:      rds,
		CacheKey: cacheKey,
	}
}

// SelectOneByPK 根据主键查询
func (r *CacheRepo[T]) SelectOneByPK(ctx context.Context, pk interface{}) (*T, error) {
	var res T
	queryFn := func(res interface{}) (err error) {
		t, err := r.BaseRepo.SelectOneByPK(ctx, pk)
		if err != nil || t == nil {
			return err
		}
		return copier.Copy(res, t).Error
	}
	if err := r.take(ctx, r.PkCacheKey(pk), &res, queryFn); err != nil {
		return nil, err
	}

	// 由于res变量是结构体变量，它的零值不是nil，导致调用者无法得知是否查到了数据
	// 所以这里要判断零值，并返回nil
	if isZero(res) {
		return nil, nil
	}

	return &res, nil
}

// take 从缓存中获取数据，如果缓存中没有，则从数据库中查询，并写入缓存
func (r *CacheRepo[T]) take(
	ctx context.Context,
	key string,
	res interface{},
	query queryFn,
) error {
	v, err := r.rds.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		if err = query(res); err != nil {
			return err
		}

		// 存在缓存穿透，内网环境不考虑了
		if isZero(res) {
			return nil
		}

		return r.setCache(ctx, key, res)
	}

	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(v), res)
}

func (r *CacheRepo[T]) setCache(ctx context.Context, key string, val interface{}) error {
	log.Debugw(
		log.DefaultMessageKey, "set cache",
		"key", key,
		"value", val,
	)

	data, err := json.Marshal(val)
	if err != nil {
		return err
	}

	return r.rds.Set(ctx, key, string(data), 0).Err()
}

// SelectOneByIndex 根据索引查询
// 例如：根据用户名查询用户
func (r *CacheRepo[T]) SelectOneByIndex(
	ctx context.Context,
	index string,
	provider IndexCacheKeyProvider,
) (*T, error) {
	var (
		res           T
		indexCacheKey = provider.IndexCacheKey(index)
	)
	indexQueryFn := func() (pk interface{}, err error) {
		t, err := r.BaseRepo.SelectOneByMap(ctx, map[string]interface{}{provider.FieldName(): index})
		if err != nil {
			return "", err
		}

		if t == nil {
			return "", nil
		}

		if err = copier.Copy(&res, t).Error; err != nil {
			return "", err
		}

		return r.GetPk(t)
	}
	err := r.takeByIndex(ctx, &res, indexCacheKey, indexQueryFn)
	if err != nil {
		return nil, err
	}

	// 由于res变量是结构体变量，它的零值不是nil，导致调用者无法得知是否查到了数据
	// 所以这里要判断零值，并返回nil
	if isZero(res) {
		return nil, nil
	}

	return &res, err
}

// takeByIndex 根据索引从缓存中获取数据，如果缓存中没有，则从数据库中查询，并写入缓存
// 流程：先根据索引查出主键，再通过主键查询，类似 mysql
func (r *CacheRepo[T]) takeByIndex(
	ctx context.Context,
	res interface{},
	indexCacheKey string,
	indexQuery indexQueryFn,
) error {
	var (
		primaryKey interface{}
		found      bool
	)
	queryFn := func(val interface{}) (err error) {
		primaryKey, err = indexQuery()
		if err != nil || isZero(primaryKey) {
			return
		}
		found = true
		return r.setCache(ctx, r.PkCacheKey(primaryKey), res)
	}

	// 在缓存中根据索引查询主键
	err := r.take(ctx, indexCacheKey, &primaryKey, queryFn)
	if err != nil {
		return err
	}

	// 没找到主键，从库里查出数据
	if found {
		return nil
	}

	if isZero(primaryKey) {
		return nil
	}

	// 找到主键，再根据主键查询
	t, err := r.SelectOneByPK(ctx, primaryKey)
	if err != nil {
		return err
	}
	return copier.Copy(res, t).Error
}

// UpdateByPKWithMap 根据id更新，支持零值
// 更新数据库后，删除缓存
func (r *CacheRepo[T]) UpdateByPKWithMap(ctx context.Context, pk interface{}, updateData map[string]interface{}) (int64, error) {
	var res int64

	cacheKeys, err := r.needDeleteCacheKey(ctx, pk, updateData)
	if err != nil {
		return 0, err
	}

	// 避免缓存不一致，使用事务
	// 带来的影响是：原本是先更新数据库，再更新缓存，现在是先删除缓存，再更新数据库
	// 使用延迟双删来减小缓存不一致的窗口
	err = r.BaseRepo.ExecTx(ctx, func(ctx context.Context) error {
		row, err := r.exec(ctx,
			func() (res int64, err error) {
				return r.BaseRepo.UpdateByPKWithMap(ctx, pk, updateData)
			},
			cacheKeys...,
		)
		if err != nil {
			return err
		}

		return copier.Copy(&res, row).Error
	})

	return res, err
}

func (r *CacheRepo[T]) needDeleteCacheKey(ctx context.Context, pk interface{}, updateData map[string]interface{}) ([]string, error) {
	pkCacheKey := r.PkCacheKey(pk)
	t, err := r.SelectOneByPK(ctx, pk)
	if err != nil {
		return nil, err
	}

	providers := r.IndexCacheKeys()
	cacheKeys := make([]string, 0, len(providers)+1)
	cacheKeys = append(cacheKeys, pkCacheKey)

	for _, provider := range providers {
		// 如果更新了索引字段，需要删除索引缓存
		if _, ok := updateData[stringx.FirstUpper(provider.FieldName())]; ok {
			cacheKeys = append(cacheKeys, provider.IndexCacheKeyFromStruct(t))
		}
	}
	return cacheKeys, nil
}

func (r *CacheRepo[T]) exec(ctx context.Context, exec execFn, keys ...string) (int64, error) {
	res, err := exec()
	if res == 0 || err != nil {
		return 0, err
	}

	return res, r.delCache(ctx, keys...)
}

func (r *CacheRepo[T]) delCache(ctx context.Context, keys ...string) error {
	// redis集群下，key可能不在同一个节点，所以不能批量去删除
	for _, key := range keys {
		if err := r.rds.Del(ctx, key).Err(); err != nil {
			return errors.Wrapf(err, "delate cache error, key:%v", key)
		}
	}

	// 延迟双删，减小缓存不一致的窗口
	go func() {
		ctx := context.Background()
		time.Sleep(time.Second)
		for _, key := range keys {
			if err := r.rds.Del(ctx, key).Err(); err != nil {
				log.Error(errors.Wrapf(err, "delay delate cache error, key:%v", key))
			}
		}
	}()

	return nil
}

// DeleteByPK 根据主键删除，支持单个主键或者一个主键数组
func (r *CacheRepo[T]) DeleteByPK(ctx context.Context, pks interface{}) (int64, error) {
	var res int64

	cacheKeys, err := r.allCacheKey(ctx, pks)
	if err != nil {
		return 0, err
	}

	// 避免缓存不一致，使用事务
	// 带来的影响是：原本是先更新数据库，再更新缓存，现在是先删除缓存，再更新数据库
	// 使用延迟双删来减小缓存不一致的窗口
	err = r.BaseRepo.ExecTx(ctx, func(ctx context.Context) error {
		row, err := r.exec(ctx,
			func() (res int64, err error) {
				return r.BaseRepo.DeleteByPK(ctx, pks)
			},
			cacheKeys...,
		)
		if err != nil {
			return err
		}

		return copier.Copy(&res, row).Error
	})

	return res, err
}

// allCacheKey 包括主键缓存和索引缓存
func (r *CacheRepo[T]) allCacheKey(ctx context.Context, pks interface{}) ([]string, error) {
	pkArr := reflectx.Interface2Array(pks)

	cacheKeys := lo.Map(pkArr, func(pk interface{}, i int) string {
		return r.PkCacheKey(pk)
	})

	ts, err := r.BaseRepo.SelectByPK(ctx, pks)
	if err != nil {
		return nil, err
	}

	providers := r.IndexCacheKeys()
	for _, t := range ts {
		for _, provider := range providers {
			cacheKeys = append(cacheKeys, provider.IndexCacheKeyFromStruct(t))
		}
	}

	return cacheKeys, nil
}

func isZero(i interface{}) bool {
	// 判断interface{}类型的空字符串
	if s, ok := i.(*interface{}); ok && fmt.Sprintf("%s", *s) == "" {
		return true
	}

	// 判断结构体、数字的零值
	return reflectx.Indirect(reflect.ValueOf(i)).IsZero()
}
