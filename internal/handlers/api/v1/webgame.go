package v1

import "github.com/gin-gonic/gin"

type Webgame struct{}

func (*Webgame) Detail(c *gin.Context) {}

func (*Webgame) Create(c *gin.Context) {}

func (*Webgame) Delete(c *gin.Context) {}

func (*Webgame) List(c *gin.Context) {}

func (*Webgame) Update(c *gin.Context) {}

func (*Webgame) SyncTo(c *gin.Context) {}

func (*Webgame) SyncFrom(c *gin.Context) {}
