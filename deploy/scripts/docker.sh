#!/bin/bash

# relative paths for services
reacttv_base="../.."
reacttv_services=$reacttv_base/services

# Locate Dockerfiles within ReactTV and return absolute directories
services=($(find $reacttv_services -iname 'dockerfile' -not -path "*/.garden/*" | xargs dirname | xargs readlink -f))

registry=ghcr.io/reacttv

# Run "docker build" against the services/* dirs.
function build() {
    version=$1

    cd $reacttv_base
    for service in "${services[@]}"; do
        service_image_name=$registry/$(basename $service):$version
        echo "docker build: $service_image_name"
        docker build -f $service/Dockerfile . -t $service_image_name --build-arg VERSION="itworks" --target prod
    done
}

function push() {
    version=$1

    for service in "${services[@]}"; do
        service_image_name=$registry/$(basename $service):$version
        echo "docker push: $service_image_name"
        docker push $service_image_name
    done
}
