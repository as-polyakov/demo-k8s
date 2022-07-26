FROM alpine:3.14
#FROM plugins/base:linux-amd64

  #LABEL maintainer="Adon988 <adon988@gmail.com>" \
  #org.label-schema.name="go-github-action-helloworld" \
#  org.label-schema.vendor="Adon988" \
 # org.label-schema.schema-version="1.0.1"

  EXPOSE 8080

  COPY release/demo-k8s /bin/
  RUN chmod 755 /bin/demo-k8s

  ENTRYPOINT ["/bin/demo-k8s"]

