#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

helm repo add stable https://charts.helm.sh/stable
helm -n mysql upgrade --install --create-namespace mysql stable/mysql --version 1.6.9 \
    --set imageTag=8.3.0,mysqlRootPassword="abc123",mysqlPassword="abc123"
