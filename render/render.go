package render

import "loly-eiei/config"

type Render struct {
	conf *config.Config
}

func NewRender(conf *config.Config) *Render {
	return &Render{
		conf: conf,
	}
}
