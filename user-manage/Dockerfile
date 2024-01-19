FROM alpine:latest
MAINTAINER zhangsan
ARG SUB_MODULE
USER root
COPY bin/$SUB_MODULE /app/server
COPY cmd/$SUB_MODULE/configs /app/configs
RUN chmod +x /app/server
EXPOSE 8000
EXPOSE 9000
WORKDIR /app/
CMD ["./server","-conf","./configs/config.yaml","-logLevel","info"]
