FROM golang:alpine AS build
RUN apk update && apk add git

WORKDIR /go/src/github.com/povilasv/pihole_exporter/

ENV GOOS=linux 
ENV GOARCH=arm 
ENV GOARM=7

ADD . .
RUN go get -t -v ./...
RUN go build -o /pihole_exporter main.go && chmod +x /pihole_exporter

FROM multiarch/alpine:armhf-edge

COPY --from=build /pihole_exporter /pihole_exporter
WORKDIR /
ENTRYPOINT ["./pihole_exporter"]
CMD ["-h"]
EXPOSE 9311
