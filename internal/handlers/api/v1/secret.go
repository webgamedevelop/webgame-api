package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"

	"github.com/webgamedevelop/webgame-api/internal/models"
	pkgclient "github.com/webgamedevelop/webgame-api/pkg/kubernetes/client"
	pkgsecret "github.com/webgamedevelop/webgame-api/pkg/kubernetes/secret"
)

type Secret struct{}

// Create image pull secret
//
//	@Tags			secret
//	@Summary		create image pull secret
//	@Description	create image pull secret
//	@Param			secret	body	models.ImagePullSecret	true	"secret creation request"
//	@Produce		json
//	@Success		200	{object}	detailResponse[models.ImagePullSecret]
//	@Failure		400	{object}	simpleResponse
//	@Failure		500	{object}	simpleResponse
//	@Router			/secret/create [post]
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
		result, err := pkgsecret.Create(
			context.Background(),
			pkgclient.Client(),
			secret.SecretName,
			secret.SecretNamespace,
			secret.DockerServer,
			secret.DockerUsername,
			secret.DockerPassword,
			secret.DockerEmail,
		)
		if err != nil {
			return err
		}
		klog.InfoS("create secret", "name", secret.SecretName, "namespace", secret.SecretNamespace, "result", result)
		return nil
	}

	if _, err = secret.Create(fn); err != nil {
		badResponse(c, http.StatusInternalServerError, err)
		return
	}

	DetailResponse(c, secret)
	return
}

// Update image pull secret
//
//	@Tags			secret
//	@Summary		update image pull secret
//	@Description	update image pull secret
//	@Param			secret	body	models.ImagePullSecret	true	"secret update request"
//	@Produce		json
//	@Success		200	{object}	detailResponse[models.ImagePullSecret]
//	@Failure		400	{object}	simpleResponse
//	@Failure		500	{object}	simpleResponse
//	@Router			/secret/update [post]
func (s Secret) Update(c *gin.Context) {
	var (
		secret models.ImagePullSecret
		err    error
	)

	if err = c.ShouldBindJSON(&secret); err != nil {
		badResponse(c, http.StatusBadRequest, err)
		return
	}

	if secret.ID == 0 {
		err = fmt.Errorf("secret ID not set")
		badResponse(c, http.StatusBadRequest, err)
		return
	}

	fn := func() error {
		result, err := pkgsecret.Create(
			context.Background(),
			pkgclient.Client(),
			secret.SecretName,
			secret.SecretNamespace,
			secret.DockerServer,
			secret.DockerUsername,
			secret.DockerPassword,
			secret.DockerEmail,
		)
		if err != nil {
			return err
		}
		klog.InfoS("update secret", "name", secret.SecretName, "namespace", secret.SecretNamespace, "result", result)
		return nil
	}

	if _, err = secret.Update(fn); err != nil {
		badResponse(c, http.StatusInternalServerError, err)
		return
	}

	DetailResponse(c, secret)
	return
}

// List returns a list of secret
//
//	@Tags			secret
//	@Summary		list of the secret
//	@Description	list of the secret
//	@Param			id	query	models.Paginator	true	"secret list request"
//	@Produce		json
//	@Success		200	{object}	listResponse[[]models.ImagePullSecret, models.ImagePullSecret]
//	@Failure		400	{object}	simpleResponse
//	@Failure		500	{object}	simpleResponse
//	@Router			/secret/list [get]
func (s Secret) List(c *gin.Context) {
	var (
		paginator models.Paginator
		secrets   []models.ImagePullSecret
		err       error
	)

	if err = c.ShouldBindQuery(&paginator); err != nil {
		badResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := models.List(&secrets, &paginator, nil); err != nil {
		badResponse(c, http.StatusInternalServerError, err)
		return
	}

	ListResponse(c, secrets)
	return
}

// Detail returns the details of the secret
//
//	@Tags			secret
//	@Summary		details of the secret
//	@Description	details of the secret
//	@Param			id	query	string	true	"secret id"
//	@Produce		json
//	@Success		200	{object}	detailResponse[models.ImagePullSecret]
//	@Failure		400	{object}	simpleResponse
//	@Failure		500	{object}	simpleResponse
//	@Router			/secret/detail [get]
func (s Secret) Detail(c *gin.Context) {
	var (
		query = &struct {
			ID uint `form:"id"`
		}{}
		secret models.ImagePullSecret
		err    error
	)

	if err = c.ShouldBindQuery(&query); err != nil {
		badResponse(c, http.StatusBadRequest, err)
		return
	}

	secret.ID = query.ID
	if _, err = secret.Detail(); err != nil {
		badResponse(c, http.StatusInternalServerError, err)
		return
	}

	DetailResponse(c, secret)
	return
}

func (s Secret) Delete(c *gin.Context) {
	// TODO implement me
	panic("implement me")
}

func (s Secret) Sync(c *gin.Context) {
	// TODO implement me
	panic("implement me")
}
