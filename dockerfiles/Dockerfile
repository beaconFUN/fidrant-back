FROM golang:1.20.5-alpine3.18 as builder

# update
RUN apk update && apk upgrade
RUN rm -rf /var/cache/apk/*

# copy
COPY ./go.mod ./go.sum /src/
COPY ./main.go /src/

WORKDIR /src

# build
ENV GOOS=linux
ENV GOARCH=amd64
RUN go mod download
RUN go build -o main ./main.go

# Lambda
FROM gcr.io/distroless/static as lambda

COPY --from=builder /src/main /bin/main
CMD [ "/bin/main" ]

