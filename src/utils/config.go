package utils

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Conf struct {
	DBHost     string
	DBPassword string
}

func Config(envs ...string) Conf {
	conf := Conf{}
	env := "Dev"

	if envs != nil {
		env = envs[0]
	}

	getFilename := fmt.Sprintf("./%sEnv.json", env)
	err := gonfig.GetConf(getFilename, &conf)

	if err != nil {
		panic(err)
	}

	return conf
}
