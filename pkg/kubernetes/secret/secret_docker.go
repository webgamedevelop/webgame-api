package secret

import (
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/kubectl/pkg/cmd/create"
	"k8s.io/kubectl/pkg/cmd/util"
)

func Create(name, namespace, dockerServer, dockerUsername, dockerPassword, dockerEmail string, enforceNamespace bool) error {
	var err error
	ioStreams := genericiooptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}

	configFlags := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag().WithDiscoveryBurst(300).WithDiscoveryQPS(50.0).WithWarningPrinter(ioStreams)
	matchVersionKubeConfigFlags := util.NewMatchVersionFlags(configFlags)
	f := util.NewFactory(matchVersionKubeConfigFlags)

	o := CreateSecretDockerRegistryOptions{create.NewSecretDockerRegistryOptions(ioStreams)}

	if err = o.Complete(f, name, namespace, dockerServer, dockerUsername, dockerPassword, dockerEmail, enforceNamespace); err != nil {
		return err
	}

	if err = o.Validate(); err != nil {
		return err
	}

	return o.Run()
}

type CreateSecretDockerRegistryOptions struct {
	*create.CreateSecretDockerRegistryOptions
}

func (o *CreateSecretDockerRegistryOptions) Complete(f util.Factory, name, namespace, dockerServer, dockerUsername, dockerPassword, dockerEmail string, enforceNamespace bool) error {
	var err error
	o.Name = name
	o.DryRunStrategy = util.DryRunNone
	o.Namespace = namespace
	o.EnforceNamespace = enforceNamespace
	o.Server = dockerServer
	o.Username = dockerUsername
	o.Password = dockerPassword
	o.Email = dockerEmail

	restConfig, err := f.ToRESTConfig()
	if err != nil {
		return err
	}

	o.Client, err = corev1client.NewForConfig(restConfig)
	if err != nil {
		return err
	}

	o.PrintObj = func(obj runtime.Object) error { return nil }
	return nil
}
