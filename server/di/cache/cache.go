package cache

type Cache interface {
	Get(key string) (string, bool)
	Set(key string, value string)
}

type CacheImpl struct {
	cache map[string]string
}

func NewCacheImpl() *CacheImpl {
	return &CacheImpl{cache: map[string]string{}}
}

func (c *CacheImpl) Get(key string) (string, bool) {
	v, ok := c.cache[key]
	return v, ok
}

func (c *CacheImpl) Set(key string, value string) {
	c.cache[key] = value
}

type CacheNoop struct{}

func NewCacheNoop() *CacheNoop {
	return &CacheNoop{}
}

func (c *CacheNoop) Get(key string) (string, bool) {
	return "", false
}

func (c *CacheNoop) Set(key string, value string) {}
