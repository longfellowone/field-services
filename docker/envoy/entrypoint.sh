#!/bin/sh
set -e

cat /tmpl/envoy.yaml.tmpl | envsubst \$SERVICE > /etc/envoy/envoy.yaml

/usr/local/bin/envoy -c /etc/envoy/envoy.yaml # -l trace

>&2 echo "Routing requests from :8080 to $SERVICE:9090"
