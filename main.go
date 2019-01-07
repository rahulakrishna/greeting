package main

import (
	"fmt"
	"log"

	"github.com/fogleman/gg"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
	//"net/http"
)

// Poem structure. Defines Title, Body and Author
type Poem struct {
	Title  string   `json:"title" binding:"required"`
	Body   []string `json:"body" binding:"required"`
	Author string   `json:"author" binding:"required"`
}

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./frontend/dist/greeting", true)))
	router.Use(static.Serve("/download", static.LocalFile("./output", true)))

	// Routes
	api := router.Group("/api/v1")
	{
		api.POST("/card", func(c *gin.Context) {
			// Data from post form
			var poem Poem
			c.BindJSON(&poem)

			// Read the image
			im, err := gg.LoadImage("frames/f1.png")

			if err != nil {
				log.Fatal(err)
			}

			// Create a context with dimensions same as the image
			dc := gg.NewContextForImage(im)

			// Set text color
			dc.SetRGB255(194, 24, 91)

			// Load the fontface and set the size as 24
			if fontFaceErr := dc.LoadFontFace("fonts/baybuom.ttf", 22); fontFaceErr != nil {
				panic(fontFaceErr)
			}

			// Set the text width and height
			width, height := float64(dc.Width())/2, float64(dc.Height())/3

			sep := "______________o0o______________"

			dc.DrawStringAnchored(poem.Title, width, height, -0.4, -1.5)

			dc.DrawStringAnchored(sep, width, height, 0.1, 0.1)

			for _, line := range poem.Body {
				height += 42
				dc.DrawStringAnchored(line, width, height, 0.1, -0.7)
			}

			dc.DrawStringAnchored("- "+poem.Author, width, height, -0.7, -1)

			dc.DrawImage(im, 0, 0)

			dc.Clip()

			u4, err := uuid.NewV4()

			if err != nil {
				panic(err)
			}

			fileName := fmt.Sprintf("%s.png", u4)

			dc.SavePNG("output/" + fileName)

			c.JSON(200, gin.H{
				"data":   poem,
				"fileId": fmt.Sprintf("%s", u4),
			})
		})
	}

	router.NoRoute(func(c *gin.Context) {
		c.File("./frontend/dist/greeting/index.html")
	})

	router.Run(":8899")
}
