FROM golang:1.8.3 as go_builder
WORKDIR /go/src/github.com/jay-khatri/quack
RUN go get -d -v golang.org/x/net/html
RUN go get -d -v github.com/gorilla/handlers
RUN go get -d -v github.com/gorilla/mux
COPY ./main.go  .
ADD static/ static/
# Replace the dev instance of index.html with the prod version.
RUN rm -rf static/index.html
RUN mv static/index-prod.html static/index.html
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o quack .


FROM node:7 as node_builder
WORKDIR /go/src/github.com/jay-khatri/quack/frontend
ADD /frontend .
RUN rm -rf node_modules
RUN yarn install
RUN yarn run build


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Copy dependant files from go builder.
COPY --from=go_builder /go/src/github.com/jay-khatri/quack/quack .
COPY --from=go_builder /go/src/github.com/jay-khatri/quack/static/index.html ./static/index.html
# Copy dependant files from node builder.
COPY --from=node_builder /go/src/github.com/jay-khatri/quack/static/index.js ./static/index.js
CMD ["./quack"]
