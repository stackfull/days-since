package src

import (
	"github.com/kataras/iris/v12"
	"github.com/rs/zerolog/log"
)

func Router() func(iris.Party) {
	return func(r iris.Party) {
		r.Get("/", func(c iris.Context) {
			_, err := c.Write([]byte(`THIS WILL BE SOMETHING`))
			if err != nil {
				log.Warn("Couldn't talk to client", err)
			}
		})
	}
}
