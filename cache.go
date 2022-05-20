package cache

import "time"

type Cache struct {
	data map[string]string
	dead map[string]time.Time
}

var cnt int

func NewCache() Cache {
	return Cache{data:map[string]string{}, time:map[string]time.Time{}}
}

func (c Cache) Get(key string) (string, bool) {
	if c.dead[key].After(time.Now()) {
		return c.m[key], true
		}
	return "", false
}

func (c Cache) Put(key, value string) {
	c.data[key] = value
	c.dead[key] = 0
}

func (c Cache) Keys() []string {
	var s []string
	for k := range c.data {
		s = append(s, k)
	}
	return s
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.data[key] = value
	c.dead[key] = deadline
}
