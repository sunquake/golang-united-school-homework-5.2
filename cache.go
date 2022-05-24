package cache

import "time"

type ctime struct {
	val string
	tim time.Time
}

type Cache struct {
	data map[string]ctime
}

func NewCache() Cache {
	return Cache{data:map[string]ctime{}}
}

func (c Cache) Get(key string) (string, bool) {
	if c.data[key].tim.After(time.Now()) {
		return c.data[key].val, true
		}
	delete(c.data, key)
	return "", false
}

func (c Cache) Put(key, value string) {
	c.data[key].val = value
	c.dead[key].tim = time.Now().AddDate(100,0,0)
}

func (c Cache) Keys() []string {
	var s []string
	for k := range c.data {
		s = append(s, k)
	}
	return s
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.data[key].val = value
	c.dead[key].tim = deadline
}
