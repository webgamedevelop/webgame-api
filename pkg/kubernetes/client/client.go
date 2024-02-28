package client

import (
	"context"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

var c client.Client

func Init(config *rest.Config, options client.Options) error {
	cli, err := client.New(config, options)
	if err != nil {
		return err
	}
	c = cli
	return nil
}

func Client() client.Client {
	return c
}

func Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	return c.Get(ctx, key, obj, opts...)
}

func List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	return c.List(ctx, list, opts...)
}

func Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	return c.Create(ctx, obj, opts...)
}

func Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	return c.Delete(ctx, obj, opts...)
}

func Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	return c.Update(ctx, obj, opts...)
}

func Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	return c.Patch(ctx, obj, patch, opts...)
}

func DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	return c.DeleteAllOf(ctx, obj, opts...)
}

func CreateOrUpdate(ctx context.Context, obj client.Object, f controllerutil.MutateFn) (controllerutil.OperationResult, error) {
	return controllerutil.CreateOrUpdate(ctx, c, obj, f)
}

func CreateOrPatch(ctx context.Context, obj client.Object, f controllerutil.MutateFn) (controllerutil.OperationResult, error) {
	return controllerutil.CreateOrPatch(ctx, c, obj, f)
}
