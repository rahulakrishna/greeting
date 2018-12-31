package main

import (
  "github.com/fogleman/gg"
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/static"
  "github.com/nu7hatch/gouuid"
  "log"
  "fmt"
  //"net/http"
)

type Poem struct {
  Title string `json:"title" binding:"required"`
  Body []string `json:"body" binding:"required"`
  Author string `json:"body" binding:"required"`
}

func main() {
  router := gin.Default()

  router.Use(static.Serve("/", static.LocalFile("./frontend/dist/greeting", true)))
  router.Use(static.Serve("/download", static.LocalFile("./output", true)))

  api := router.Group("/api/v1")
    {
    api.POST("/card", func(c *gin.Context) {
      var poem Poem
      c.BindJSON(&poem)
      log.Print(poem)
      im, err := gg.LoadImage("frames/f1.png")

      if err != nil {
        log.Fatal(err)
      }

      dc := gg.NewContextForImage(im)

      dc.SetRGB255(220, 94, 94)

      if err := dc.LoadFontFace("fonts/bucthu.ttf", 44); err != nil {
        panic(err)
      }

      width, height := float64(dc.Width()) / 1.8, float64(dc.Height()) / 3

      sep := "-------------------o0o-------------------"

      dc.DrawStringAnchored(poem.Title, width, height, 0.5, 0.5)

      for _, line := range poem.Body {
        height += 42
        dc.DrawStringAnchored(line, width, height, 0.5, 0.5)
      }

      dc.DrawStringAnchored(sep, width, height, 0.5, 0.5)

      dc.DrawStringAnchored(poem.Author, width, height, 0.5, 0.5)

      dc.DrawImage(im, 0, 0)

      dc.Clip()

      u4, err := uuid.NewV4()

      if err != nil {
        panic(err)
      }

      fileName := fmt.Sprintf("%s.png", u4)

      dc.SavePNG("output/" + fileName)

      c.JSON(200, gin.H{
        "data": poem,
        "fileId": fmt.Sprintf("%s", u4),
      })
    })
  }

  router.NoRoute(func(c *gin.Context) {
    c.File("./frontend/dist/greeting/index.html")
  })

  router.Run(":8899")
}
