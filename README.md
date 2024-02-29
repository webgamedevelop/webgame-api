# webgame-api

## Description
webgame api component

## Getting Started

### Prerequisites
- go version v1.20.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- helm version v3+
- mysql version 5.7.x

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/webgame-api:tag

# for example
make docker-build docker-push IMG=webgamedevelop/webgame-api:v0.0.1-alpha.3
# or
make docker-build docker-push
```

**NOTE:** This image ought to be published in the personal registry you specified.
And it is required to have access to pull the image from the working environment.
Make sure you have the proper permission to the registry if the above commands don’t work.

**Deploy the webgame-api component to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/webgame-api:tag

# for example
make deploy IMG=webgamedevelop/webgame-api:v0.0.1-alpha.3
# or
make deploy
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin
privileges or be logged in as admin.

### To Uninstall
**UnDeploy the controller from the cluster:**

```sh
make undeploy
```
