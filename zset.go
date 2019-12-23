package xredis

import (
	"github.com/go-redis/redis/v7"
)

<<<<<<< HEAD
type typeZset struct {
	*GoRedis
=======
type TypeZset struct {
	tz *GoRedis
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
}


func (gr *GoRedis) NewZset() *TypeZset {
	return &TypeZset{gr	}
}


<<<<<<< HEAD
func (tz *typeZset) ZAdd(key string, members ...*redis.Z) (int64, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZAdd(key string, members ...*redis.Z) (int64, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return 0, err
	}

	return tz.client.ZAdd(key, members...).Result()
}

<<<<<<< HEAD
func (tz *typeZset) ZCard(key string) (int64, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZCard(key string) (int64, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return 0, err
	}

	return tz.client.ZCard(key).Result()
}


<<<<<<< HEAD
func (tz *typeZset) ZCount(key, min, max string) (int64, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZCount(key, min, max string) (int64, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return 0, err
	}

	return tz.client.ZCount(key, min, max).Result()
}


<<<<<<< HEAD
func (tz *typeZset) ZIncrBy(key string, increment float64, member string) (float64, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZIncrBy(key string, increment float64, member string) (float64, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return 0, err
	}

	return tz.client.ZIncrBy(key, increment, member).Result()
}


<<<<<<< HEAD
func (tz *typeZset) ZInterStore(destination string, store *redis.ZStore) (int64, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZInterStore(destination string, store *redis.ZStore) (int64, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return 0, err
	}

	return tz.client.ZInterStore(destination, store).Result()
}


<<<<<<< HEAD
func (tz *typeZset) ZLexCount(key, min, max string) (int64, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZLexCount(key, min, max string) (int64, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return 0, err
	}

	return tz.client.ZLexCount(key, min, max).Result()
}


<<<<<<< HEAD
func (tz *typeZset) ZRange(key string, start, stop int64) ([]string, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZRange(key string, start, stop int64) ([]string, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return nil, err
	}

	return tz.client.ZRange(key, start, stop).Result()
}


<<<<<<< HEAD
func (tz *typeZset) ZRangeByLex(key string, opt *redis.ZRangeBy) ([]string, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZRangeByLex(key string, opt *redis.ZRangeBy) ([]string, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return nil, err
	}

	return tz.client.ZRangeByLex(key, opt).Result()
}


<<<<<<< HEAD
func (tz *typeZset) ZRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return nil, err
	}

	return tz.client.ZRangeByScore(key, opt).Result()
}


<<<<<<< HEAD
func (tz *typeZset) ZRangeByScoreWithScores(key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZRangeByScoreWithScores(key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return nil, err
	}

	return tz.client.ZRangeByScoreWithScores(key, opt).Result()
}


<<<<<<< HEAD
func (tz *typeZset) ZRank(key, member string) (int64, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZRank(key, member string) (int64, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return 0, err
	}

	return tz.client.ZRank(key, member).Result()
}

<<<<<<< HEAD
func (tz *typeZset) ZRem(key string, members ...interface{}) (int64, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZRem(key string, members ...interface{}) (int64, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return 0, err
	}

	return tz.client.ZRem(key, members...).Result()
}


<<<<<<< HEAD
func (tz *typeZset) ZRemRangeByLex(key, min, max string) (int64, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZRemRangeByLex(key, min, max string) (int64, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return 0, err
	}

	return tz.client.ZRemRangeByLex(key, min, max).Result()
}


<<<<<<< HEAD
func (tz *typeZset) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return 0, err
	}

	return tz.client.ZRemRangeByRank(key, start, stop).Result()
}


<<<<<<< HEAD
func (tz *typeZset) ZRemRangeByScore(key, min, max string) (int64, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZRemRangeByScore(key, min, max string) (int64, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return 0, err
	}

	return tz.client.ZRemRangeByScore(key, min, max).Result()
}


<<<<<<< HEAD
func (tz *typeZset) ZRevRange(key string, start, stop int64) ([]string, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZRevRange(key string, start, stop int64) ([]string, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return nil, err
	}

	return tz.client.ZRevRange(key, start, stop).Result()
}


<<<<<<< HEAD
func (tz *typeZset) ZRevRangeByLex(key string, opt *redis.ZRangeBy) ([]string, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZRevRangeByLex(key string, opt *redis.ZRangeBy) ([]string, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return nil, err
	}

	return tz.client.ZRevRangeByLex(key, opt).Result()
}


<<<<<<< HEAD
func (tz *typeZset) ZRevRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZRevRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return nil, err
	}

	return tz.client.ZRevRangeByScore(key, opt).Result()
}


<<<<<<< HEAD
func (tz *typeZset) ZScore(key, member string) (float64, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZScore(key, member string) (float64, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return 0, err
	}

	return tz.client.ZScore(key, member).Result()
}


<<<<<<< HEAD
func (tz *typeZset) ZUnionStore(dest string, store *redis.ZStore) (int64, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZUnionStore(dest string, store *redis.ZStore) (int64, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return 0, err
	}

	return tz.client.ZUnionStore(dest, store).Result()
}


<<<<<<< HEAD
func (tz *typeZset) ZScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	if tz == nil {
=======
func (tz *TypeZset) ZScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	if tz.tz == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tz.Ping(); err != nil {
		return nil, 0, err
	}

	return tz.client.ZScan(key, cursor, match,count).Result()
}
