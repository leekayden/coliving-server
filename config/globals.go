package config

import "os"

type GlobalsStruct struct {
	Frontend string `env:"frontend"`
}

var Globals GlobalsStruct

func InitGlobals() {
	Globals = GlobalsStruct{
		Frontend: os.Getenv("frontend"),
	}
}
