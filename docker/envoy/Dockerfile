FROM envoyproxy/envoy:latest

COPY /docker/envoy/envoy.yaml /tmpl/envoy.yaml.tmpl
COPY /docker/envoy/entrypoint.sh /

RUN chmod 500 /entrypoint.sh

RUN apt-get update && \
    apt-get install gettext -y

ENTRYPOINT ["/entrypoint.sh"]