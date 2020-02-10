package singleton

// ============================================================================

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once
var conf = make(map[string]*Conf)

func GetConfig() map[string]*Conf {
	once.Do(func() {
		if e := config.ParseConfigWithoutDefaults(&conf); e != nil {
			log.Output(1, "client req parse config fail:"+e.Error())
		}
	})
	return conf
}
