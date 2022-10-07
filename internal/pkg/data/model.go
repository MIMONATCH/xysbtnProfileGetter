package data

// 虚研社按钮应援组件model
type Support struct {
	Uid  string `yaml:"uid"`
	Name string `yaml:"name"`
}

type Supports struct {
	Supports []Support `yaml:"supports"`
}
