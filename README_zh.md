# webgame-api
<p>
  简体中文 |
  <a href="./README.md">Docs</a>
</p>

## 组件描述
webgame api 组件

## 快速入门

### 外部依赖
- go version v1.20.0+.
- docker version 17.03+.
- kubectl version v1.11.3+.
- helm version v3+.
- mysql version 5.7+

### 将 webgame-api 部署到集群中
**构建镜像并推送到由 `IMG` 指定的镜像仓库中:**

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

**使用 `IMG` 参数将指定版本的 webgame-api 组件部署到集群中:**

```sh
make deploy IMG=<some-registry>/webgame-api:tag

# for example
make deploy IMG=webgamedevelop/webgame-api:v0.0.1-alpha.3
# or
make deploy
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin
privileges or be logged in as admin.

### 卸载
**从集群中卸载 webgame-api 组件:**

```sh
make undeploy
```
