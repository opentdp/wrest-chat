package plugin

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
)

type Cache struct {
	mu   sync.RWMutex
	data map[string]uint
	file string
}

func NewCache(file string) *Cache {

	c := &Cache{file: file}
	return c.init()

}

func (c *Cache) Put(k string, v uint) {

	c.data[k] = v
	c.save()

}

func (c *Cache) Get(k string) uint {

	return c.data[k]

}

func (c *Cache) Del(k string) {

	delete(c.data, k)
	c.save()

}

func (c *Cache) init() *Cache {

	c.mu.Lock()
	defer c.mu.Unlock()

	data := map[string]uint{}

	temp, _ := os.ReadFile(c.file)
	list := strings.Split(string(temp), "\n")

	for _, line := range list {
		if strings.Contains(line, "=") {
			x := strings.Split(strings.TrimSpace(line), "=")
			v, _ := strconv.ParseInt(x[0], 10, 32)
			data[x[1]] = uint(v)
		}
	}

	c.data = data
	return c

}

func (c *Cache) save() *Cache {

	c.mu.Lock()
	defer c.mu.Unlock()

	list := []string{}
	for k, v := range c.data {
		list = append(list, fmt.Sprintf("%d=%s", v, k))
	}

	data := strings.Join(list, "\n")
	os.MkdirAll(path.Dir(c.file), 0755)
	os.WriteFile(c.file, []byte(data), 0644)

	return c

}
