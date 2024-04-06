package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	networkingv1 "k8s.io/api/networking/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/webgamedevelop/webgame-api/internal/handlers/api"
	"github.com/webgamedevelop/webgame-api/internal/models"
	pkgclient "github.com/webgamedevelop/webgame-api/pkg/kubernetes/client"
)

var _ api.Resource = &IngressClass{}

type IngressClass struct{}

func (i IngressClass) Create(c *gin.Context) {
	BadResponse(c, http.StatusNotImplemented, fmt.Errorf("not implemented"))
	return
}

// Update ingress class
//
//	@Tags			ingressClass
//	@Summary		update ingress class
//	@Description	update ingress class
//	@Accept			json
//	@Param			ingressClass	body	models.IngressClass	true	"ingress class update request"
//	@Produce		json
//	@Success		200	{object}	detailResponse[models.IngressClass]
//	@Failure		400	{object}	simpleResponse
//	@Failure		401	{object}	simpleResponse
//	@Failure		500	{object}	simpleResponse
//	@Router			/ingressclass/update [post]
func (i IngressClass) Update(c *gin.Context) {
	var (
		class models.IngressClass
		err   error
	)

	if err = c.ShouldBindJSON(&class); err != nil {
		BadResponse(c, http.StatusBadRequest, err)
		return
	}

	if class.ID == 0 {
		err = fmt.Errorf("ingress class ID not set")
		BadResponse(c, http.StatusBadRequest, err)
		return
	}

	if err = class.Update(func() error { return nil }); err != nil {
		BadResponse(c, http.StatusInternalServerError, err)
		return
	}

	DetailResponse(c, class)
	return
}

// List returns a list of ingress class
//
//	@Tags			ingressClass
//	@Summary		list ingress classes
//	@Description	list ingress classes
//	@Param			id	query	models.Paginator	true	"ingress class list request"
//	@Produce		json
//	@Success		200	{object}	listResponse[[]models.IngressClass, models.IngressClass]
//	@Failure		400	{object}	simpleResponse
//	@Failure		401	{object}	simpleResponse
//	@Failure		500	{object}	simpleResponse
//	@Router			/ingressclass/list [get]
func (i IngressClass) List(c *gin.Context) {
	var (
		paginator models.Paginator
		classes   []models.IngressClass
		err       error
	)

	if err = c.ShouldBindQuery(&paginator); err != nil {
		BadResponse(c, http.StatusBadRequest, err)
		return
	}

	if err = models.List(&classes, &paginator, nil); err != nil {
		BadResponse(c, http.StatusInternalServerError, err)
		return
	}

	ListResponse(c, classes)
	return
}

// Detail returns the details of ingressClass
//
//	@Tags			ingressClass
//	@Summary		details of the ingressClass
//	@Description	details of the ingressClass
//	@Param			id	query	string	true	"ingress class id"
//	@Produce		json
//	@Success		200	{object}	detailResponse[models.IngressClass]
//	@Failure		400	{object}	simpleResponse
//	@Failure		401	{object}	simpleResponse
//	@Failure		500	{object}	simpleResponse
//	@Router			/ingressclass/detail [get]
func (i IngressClass) Detail(c *gin.Context) {
	var (
		query = &struct {
			ID uint `form:"id"`
		}{}
		class models.IngressClass
		err   error
	)

	if err = c.ShouldBindQuery(&query); err != nil {
		BadResponse(c, http.StatusBadRequest, err)
		return
	}

	class.ID = query.ID
	if err = class.Detail(); err != nil {
		BadResponse(c, http.StatusInternalServerError, err)
		return
	}

	DetailResponse(c, class)
	return
}

// Delete a ingressClass
//
//	@Tags			ingressClass
//	@Summary		delete a ingress class
//	@Description	delete a ingress class
//	@Param			id	query	string	true	"ingress class id"
//	@Produce		json
//	@Success		200	{object}	detailResponse[models.IngressClass]
//	@Failure		400	{object}	simpleResponse
//	@Failure		401	{object}	simpleResponse
//	@Failure		500	{object}	simpleResponse
//	@Router			/ingressclass/delete [delete]
func (i IngressClass) Delete(c *gin.Context) {
	var (
		query = &struct {
			ID uint `form:"id"`
		}{}
		class models.IngressClass
		err   error
	)

	if err = c.ShouldBindQuery(&query); err != nil {
		BadResponse(c, http.StatusBadRequest, err)
		return
	}

	fn := func() error {
		var resource networkingv1.IngressClass
		resource.SetName(class.ClassName)
		return client.IgnoreNotFound(pkgclient.Delete(context.Background(), &resource))
	}

	if err = models.Delete(query.ID, &class, fn); err != nil {
		BadResponse(c, http.StatusInternalServerError, err)
		return
	}

	DetailResponse(c, &class)
	return
}

func (i IngressClass) SyncTo(c *gin.Context) {
	BadResponse(c, http.StatusNotImplemented, fmt.Errorf("not implemented"))
	return
}

// SyncFrom sync ingress class from k8s cluster
//
//	@Tags			ingressClass
//	@Summary		sync ingress class from k8s cluster
//	@Description	sync ingress class from k8s cluster
//	@Produce		json
//	@Success		200	{object}	simpleResponse
//	@Failure		401	{object}	simpleResponse
//	@Failure		500	{object}	simpleResponse
//	@Router			/ingressclass/syncfrom [get]
func (i IngressClass) SyncFrom(c *gin.Context) {
	var (
		objects        networkingv1.IngressClassList
		ingressClasses []models.IngressClass
		err            error
	)

	if err = pkgclient.List(context.Background(), &objects); err != nil {
		BadResponse(c, http.StatusInternalServerError, err)
		return
	}

	for _, obj := range objects.Items {
		ingressClasses = append(ingressClasses, models.IngressClass{
			Name:      fmt.Sprintf("%s-synced-from-cluster-%d", obj.GetName(), time.Now().Unix()),
			ClassName: obj.GetName(),
			Synced:    true,
		})
	}

	if err = models.ImportData(ingressClasses); err != nil {
		BadResponse(c, http.StatusInternalServerError, err)
		return
	}

	EmptyResponse(c, http.StatusOK)
	return
}
