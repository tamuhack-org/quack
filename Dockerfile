# 1) BUILD GO BINARY
FROM golang:alpine as build-go
RUN apk --no-cache add git bzr mercurial
ENV D=/go/src/github.com/tamuhack-org/quack
RUN go get -d -v golang.org/x/net/html
RUN go get -d -v github.com/gorilla/handlers
RUN go get -d -v github.com/gorilla/mux
COPY ./quack.go $D/quack.go
COPY ./backend/ $D/backend/
COPY ./frontend/dist $D/frontend/dist
RUN rm -rf $D/frontend/dist/index.html
RUN rm -rf $D/frontend/dist/index.js
RUN cd $D && go build -o quack && cp quack /tmp/

# 2) BUILD UI
FROM node:alpine AS build-node 
RUN mkdir -p /src/ui
COPY ./frontend/package.json /src/ui/
RUN cd /src/ui && yarn install
COPY ./frontend /src/ui
# Replace the dev instance of index.html with the prod version.
RUN rm -rf /src/ui/dist/index.html
RUN mv /src/ui/dist/index-prod.html /src/ui/dist/index.html
RUN cd /src/ui && yarn build

# 3) BUILD FINAL IMAGE
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app/server/
COPY --from=build-go /tmp/quack /app/server/
COPY --from=build-node /src/ui/dist /app/server/frontend/dist
EXPOSE 8080
CMD ["./quack"]

