#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# https://github.com/rancher/local-path-provisioner

kubectl apply -f https://raw.githubusercontent.com/rancher/local-path-provisioner/v0.0.26/deploy/local-path-storage.yaml
