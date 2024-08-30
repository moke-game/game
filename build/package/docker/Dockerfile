ARG APP_NAME=game0
ARG GIT_URL="git-ej.getqccode.com"
ARG GIT_USR="docker"
# how to get your git access token:
#https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-personal-access-token-classic
ARG GIT_PWD="your_git_access_token"
# Step 1: Modules caching
FROM golang:1.22.5-alpine as modules
ARG GIT_USR
ARG GIT_PWD
ARG GIT_URL
RUN apk add --no-cache ca-certificates git
COPY go.mod go.sum /modules/
RUN echo "machine ${GIT_URL} login ${GIT_USR} password ${GIT_PWD}" > ~/.netrc
WORKDIR /modules
ENV GO111MODULE="on"
ENV GOPROXY="https://goproxy.cn,https://gocenter.io,direct"
RUN go env -w GOPRIVATE=${GIT_URL}
RUN go mod download


# Step 2: Builder
FROM golang:1.22.5-alpine as builder
ARG APP_NAME
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./cmd/${APP_NAME}/service

# Step 3: Final
FROM alpine
RUN #sed -i -e 's/http:/https:/' /etc/apk/repositories
RUN apk --no-cache add tzdata
# you can fix it to your TZ
#ENV TZ=Asia/Shanghai
COPY --from=builder /bin/app /app
CMD ["/app"]