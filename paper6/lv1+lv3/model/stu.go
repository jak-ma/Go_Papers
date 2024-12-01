package model

type Student struct {
	Name     string `yaml:"name"json:"name"`
	Id       string `yaml:"id"json:"id"`
	Comefrom string `yaml:"comefrom"json:"comefrom"`
	Sex      string `yaml:"sex"json:"sex"`
	Optime   string `yaml:"optime"json:"optime"`
}
