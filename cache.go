package cache

import "time"

type Cache struct {
	m map[string]string
}

func NewCache() Cache {
	return Cache{map[time.Now().toString()]""}
}

func (c Cache) Get(key string) (string, bool) {
	//if c.time < time.Now() {
		return c.m[key], true
		//}
	//return "", false
}

func (c Cache) Put(key, value string) {
	c.m[key] = value
}

func (c Cache) Keys() []string {
	var s []string
	for _, v := range c.m {
		s = append(s,v)
	}
	return s
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	var s []string
	for _, v := range c.m {
		//if
		s = append(s,v)
	}
	return s
}
