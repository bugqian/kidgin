package Config

import (
	"github.com/Unknwon/goconfig"
)

func Config(str string) map[string]string {
	conf := make(map[string]string)
	ini, _ := goconfig.LoadConfigFile("./Config/conf.ini")
	value, err := ini.GetValue("", str)
	if err != nil {
		conf, err = ini.GetSection(str)
		if err != nil {
			panic(err)
		}
	}
	conf[str] = value
	return conf
}
