package cache

import "time"

type Cache struct {
	m map[string]time.Time
}

func NewCache() Cache {
	return Cache{map[string]time.Time{time.Now().String():time.Now()}}
}

func (c Cache) Get(key string) (string, bool) {
	//if c.time < time.Now() {
		return c.m[key].String(), true
		//}
	//return "", false
}

func (c Cache) Put(key, value string) {
	c.m[key],_ = time.Parse(time.RFC3339, value)
}

func (c Cache) Keys() []string {
	var s []string
	for _, v := range c.m {
		s = append(s,v.String())
	}
	return s
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	
}
