FROM BASE_IMAGE

LABEL maintainer="<yangjinbot@gmail.com>"

WORKDIR=/opt/apiserver

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo "Asia/Shanghai" > /etc/timezone && \
      mkdir -p /var/log/app

COPY apiserver /opt/apiserver/bin/
COPY config.yaml /etc/apiserver/

ENTRYPOINT ["/opt/apiserver/bin/apiserver"]

CMD ["-c", "/etc/apiserver/config.yaml"]
