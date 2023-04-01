FROM kong/kong-gateway:latest

ADD /bin/go-hello /usr/local/bin/go-hello

USER kong

RUN ["/usr/local/bin/go-hello", "-dump"]