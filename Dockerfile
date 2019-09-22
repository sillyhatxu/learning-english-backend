FROM xushikuan/alpine-build:2.0 AS builder

ENV WORK_DIR=$GOPATH/src/github.com/sillyhatxu
ENV PROJECT_NAME=user-backend
WORKDIR $WORK_DIR/$PROJECT_NAME

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go


FROM xushikuan/alpine-build:1.0


ENV WORK_DIR=/go/src/github.com/sillyhatxu
ENV PROJECT_NAME=user-backend
ENV BUILDER_WORK_DIR=$WORK_DIR/$PROJECT_NAME
ENV WORK_DIR=/app

WORKDIR $WORK_DIR

ENV TIME_ZONE=Asia/Singapore
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

RUN mkdir -p logs
COPY --from=builder $BUILDER_WORK_DIR/main .
COPY --from=builder $BUILDER_WORK_DIR/config.conf .
ENTRYPOINT ./main -c config.conf