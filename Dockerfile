FROM debian:stable-slim
ENV TZ=Asia/Shanghai

WORKDIR /app

COPY ./bin /app

CMD [ "./out" ]