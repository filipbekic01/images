package main

import (
	"github.com/kataras/iris/v12"
)

// Image entity
type Image struct {
	URL string `json:"url"`
}

func main() {
	images := []Image{}

	app := iris.Default()

	tmpl := iris.HTML("./", ".html")
	tmpl.Reload(true)
	tmpl.Delims("{!", "!}")
	app.RegisterView(tmpl)

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	app.Handle("GET", "/images", func(ctx iris.Context) {
		ctx.JSON(images)
	})

	app.Handle("POST", "/images", func(ctx iris.Context) {
		var image Image

		err := ctx.ReadJSON(&image)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			return
		}

		exists := false
		for i := 0; i < len(images); i++ {
			if images[i].URL == image.URL {
				exists = true
			}
		}

		if exists {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString("Image is already saved, try another one.")
			return
		}

		images = append(images, image)
	})

	app.Listen(":5050")
}
