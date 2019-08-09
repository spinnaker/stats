FROM mirror.gcr.io/library/golang as builder

ADD . /workspace
WORKDIR /workspace

RUN ./proto/regen-protos.sh

ENV GOPROXY=https://proxy.golang.org
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go install cmd/server/main.go

# -------

FROM mirror.gcr.io/library/alpine
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/bin/main /server

ENTRYPOINT ["/server"]
