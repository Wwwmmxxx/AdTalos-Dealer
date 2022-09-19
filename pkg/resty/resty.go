package resty

import (
	"github.com/go-resty/resty/v2"
	"sync"
)

var once sync.Once = sync.Once{}

var instance *resty.Client

func NewResty() *resty.Client {
	once.Do(func() {
		instance = resty.New()
	})
	return instance
}
