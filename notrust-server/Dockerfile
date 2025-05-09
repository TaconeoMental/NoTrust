FROM bitnami/node:18.20.1 as assets

ENV UID=1000
ENV GID=1000

RUN groupadd -g "${GID}" node \
 && useradd --create-home --no-log-init -u "${UID}" -g "${GID}" node

ENV SRC_DIR /src
WORKDIR $SRC_DIR
RUN mkdir -p $SRC_DIR

COPY assets/package.json .
COPY assets/gulpfile.js .

RUN chown node:node -R $SRC_DIR

USER node

ENV PATH "$PATH:/src/node_modules/.bin/"
RUN npm install --no-audit --progress=false

ARG NODE_ENV="production"
ENV NODE_ENV="${NODE_ENV}"

FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server/main.go

# Etapa final
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 8081

ENTRYPOINT ["./server"]
