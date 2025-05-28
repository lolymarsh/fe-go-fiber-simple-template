package render

import "csdoc_fe/config"

type Render struct {
	conf *config.Config
}

func NewRender(conf *config.Config) *Render {
	return &Render{
		conf: conf,
	}
}
