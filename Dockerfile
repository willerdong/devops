FROM registry.cn-hangzhou.aliyuncs.com/willerdong/debian:stable-slim-v1.0
ENV TZ=Asia/Shanghai

WORKDIR /app

COPY ./bin /app

CMD [ "./out" ]