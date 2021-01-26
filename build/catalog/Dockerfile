FROM golang:1.10.3 AS build
WORKDIR /go/src/github.com/eavillacis/velociraptor/
COPY . .
RUN CGO_ENABLED=0 go install ./services/catalog

FROM alpine
ENV TZ=America/Guayaquil
RUN apk update && apk add ca-certificates tzdata && rm -rf /var/cache/apk/*
RUN cp /usr/share/zoneinfo/America/Guayaquil /etc/localtime
RUN echo "America/Guayaquil" > /etc/timezone
RUN date
WORKDIR /root/
COPY --from=build /go/bin/catalog .
ENTRYPOINT [ "./catalog" ]