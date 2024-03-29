package secret

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"os"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	"k8s.io/kubectl/pkg/cmd/create"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/scheme"
	"k8s.io/kubectl/pkg/util"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func Create(ctx context.Context, client client.Client, name, namespace, dockerServer, dockerUsername, dockerPassword, dockerEmail string) (controllerutil.OperationResult, error) {
	ioStreams := genericiooptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
	o := CreateSecretDockerRegistryOptions{
		CreateSecretDockerRegistryOptions: create.NewSecretDockerRegistryOptions(ioStreams),
		client:                            client,
	}

	if err := o.Complete(name, namespace, dockerServer, dockerUsername, dockerPassword, dockerEmail); err != nil {
		return controllerutil.OperationResultNone, err
	}

	if err := o.Validate(); err != nil {
		return controllerutil.OperationResultNone, err
	}

	return o.Run(ctx)
}

type CreateSecretDockerRegistryOptions struct {
	*create.CreateSecretDockerRegistryOptions
	client client.Client
}

func (o *CreateSecretDockerRegistryOptions) Complete(name, namespace, dockerServer, dockerUsername, dockerPassword, dockerEmail string) error {
	o.DryRunStrategy = cmdutil.DryRunNone
	o.CreateAnnotation = true
	o.Name = name
	o.Namespace = namespace
	o.Server = dockerServer
	o.Username = dockerUsername
	o.Password = dockerPassword
	o.Email = dockerEmail
	return nil
}

func (o *CreateSecretDockerRegistryOptions) Run(ctx context.Context) (controllerutil.OperationResult, error) {
	var secret corev1.Secret
	secret.SetNamespace(o.Namespace)
	secret.SetName(o.Name)
	mutate := func() error {
		secret.Type = corev1.SecretTypeDockerConfigJson
		secret.Data = map[string][]byte{}

		dockerConfigJSONContent, err := handleDockerCfgJSONContent(o.Username, o.Password, o.Email, o.Server)
		if err != nil {
			return err
		}
		secret.Data[corev1.DockerConfigJsonKey] = dockerConfigJSONContent
		if err = util.CreateOrUpdateAnnotation(o.CreateAnnotation, &secret, scheme.DefaultJSONEncoder()); err != nil {
			return err
		}
		return nil
	}

	for {
		result, err := controllerutil.CreateOrUpdate(ctx, o.client, &secret, mutate)
		if err != nil {
			if apierrors.IsConflict(err) {
				continue
			}
			return controllerutil.OperationResultNone, err
		}
		return result, nil
	}
}

// handleDockerCfgJSONContent serializes a ~/.docker/config.json file
func handleDockerCfgJSONContent(username, password, email, server string) ([]byte, error) {
	dockerConfigAuth := create.DockerConfigEntry{
		Username: username,
		Password: password,
		Email:    email,
		Auth:     encodeDockerConfigFieldAuth(username, password),
	}
	dockerConfigJSON := create.DockerConfigJSON{
		Auths: map[string]create.DockerConfigEntry{server: dockerConfigAuth},
	}

	return json.Marshal(dockerConfigJSON)
}

// encodeDockerConfigFieldAuth returns base64 encoding of the username and password string
func encodeDockerConfigFieldAuth(username, password string) string {
	fieldValue := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(fieldValue))
}
