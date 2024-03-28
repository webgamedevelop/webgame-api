package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/webgamedevelop/webgame-api/internal/models"
	pkgsecret "github.com/webgamedevelop/webgame-api/pkg/kubernetes/secret"
)

type Secret struct{}

func (s Secret) Create(c *gin.Context) {
	var (
		secret models.ImagePullSecret
		err    error
	)

	if err = c.ShouldBindJSON(&secret); err != nil {
		badResponse(c, http.StatusBadRequest, err)
		return
	}

	fn := func() error {
		return pkgsecret.Create(secret.SecretName, secret.SecretNamespace, secret.DockerServer, secret.DockerUsername, secret.DockerPassword, secret.DockerEmail, false)
	}

	if _, err = secret.Create(fn); err != nil {
		badResponse(c, http.StatusInternalServerError, err)
		return
	}

	okResponse(c, secret)
	return
}

func (s Secret) Update(c *gin.Context) {
	// TODO implement me
	panic("implement me")
}

func (s Secret) List(c *gin.Context) {
	// TODO implement me
	panic("implement me")
}

func (s Secret) Detail(c *gin.Context) {
	// TODO implement me
	panic("implement me")
}

func (s Secret) Delete(c *gin.Context) {
	// TODO implement me
	panic("implement me")
}
