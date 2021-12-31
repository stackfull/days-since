package src

import "github.com/kataras/iris/v12"

func AdminPages() func(iris.Party) {
	return func(r iris.Party) {
		r.Get("/", func(c iris.Context) {
			c.Write([]byte(`THIS WILL BE ADMIN`))
		})
	}
}
