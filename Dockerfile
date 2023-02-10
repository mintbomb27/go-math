FROM golang:1.20 as builder

ENV APP_HOME /go/src/todo-app

WORKDIR "$APP_HOME"
ADD . .

RUN go build -o todoapp

FROM golang:1.20

ENV APP_HOME /go/src/todo-app
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY --from=builder "$APP_HOME"/todoapp $APP_HOME

EXPOSE 8010
CMD ["./todoapp"]