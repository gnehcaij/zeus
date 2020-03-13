package dbutil

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

type DBConfig struct {
	//DriverName   string `yaml:"DriverName"`
	Timeout      string `yaml:"Timeout"`      //Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	ReadTimeout  string `yaml:"ReadTimeout"`  //Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	WriteTimeout string `yaml:"WriteTimeout"` //Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	User         string `yaml:"User"`
	Password     string `yaml:"Password"`
	Name         string `yaml:"Name"`
	Charset      string `yaml:"Charset"`
	Host         string `yaml:"Host"`
	Port         string `yaml:"Port"`
}

var dbs = make(map[string]DBConfig)
var dbLock sync.RWMutex

func InitDBConf(filename string) error {
	dbBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	dbLock.Lock()
	defer dbLock.Unlock()
	if err = yaml.Unmarshal(dbBytes, &dbs); err != nil {
		fmt.Printf("InitDBConf %v\n", err)
		return err
	}

	return nil
}

func (c DBConfig) GenerateConfig() string {
	if c.Charset == "" {
		c.Charset = "utf8"
	}
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local&timeout=%s&readTimeout=%s&writeTimeout=%s",
		c.User, c.Password, c.Host, c.Port, c.Name, c.Charset, c.Timeout, c.ReadTimeout, c.WriteTimeout)
	return config
}

func DBConf(dbName string) (*DBConfig, error) {
	dbLock.Lock()
	defer dbLock.Unlock()
	if c, ok := dbs[dbName]; ok {
		return &c, nil
	} else {
		return nil, fmt.Errorf("can't find DB with name %s", dbName)
	}
}
