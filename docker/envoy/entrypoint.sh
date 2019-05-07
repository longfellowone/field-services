#!/bin/sh
set -e

cat /tmpl/envoy.yaml.tmpl | envsubst \$SERVICE > /etc/envoy/envoy.yaml

/usr/local/bin/envoy -c /etc/envoy/envoy.yaml -l trace --log-path /tmp/envoy_info.log