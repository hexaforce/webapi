FROM golang:1.12.5-alpine AS build
WORKDIR /go/src
COPY ./ .

ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN apk add git
# RUN git config --global url."https://{{GitHubPersonalAccessToken for PrivateRepository}}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
RUN go get -d -v ./...
# RUN go test -v ./...
RUN go build -a -installsuffix cgo -o webapi .

FROM scratch AS runtime
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/webapi ./
#EXPOSE 8888/tcp
ENTRYPOINT ["./webapi"]
