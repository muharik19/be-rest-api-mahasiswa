package main

import (
	"github.com/gin-gonic/gin"
	library "github.com/mimalabs/library"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	rq := gin.New()
	vg := rq.Group("mima/v1/")
	{
		vg.POST("mahasiswa-add/", library.MahasiswaAdd)
		vg.GET("mahasiswa-all/", library.MahasiswaAll)
		vg.POST("mahasiswa-edit/", library.MahasiswaUpdate)
		vg.POST("mahasiswa-delete/", library.MahasiswaDelete)
		vg.Static("/images", "./image")
	}
	rq.Run(":83")
}
