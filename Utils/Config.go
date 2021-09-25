package Utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	redis   redis
	db      db
	appname string `yaml:"appname"`
}

type redis struct {
	host     string
	port     int
	dataBase int
	timeout  int
}

type db struct {
	host string
	port int
	root string
	pwd  string
}

func Config() {
	var conf Conf
	config, err := ioutil.ReadFile("./Config/apps_dev")
	if err != nil {
		panic(err.Error())
	}
	err1 := yaml.Unmarshal(config, conf)
	if err1 != nil {
		panic(err1.Error())
	}
	fmt.Println(conf.appname)
	fmt.Println(conf)

}
