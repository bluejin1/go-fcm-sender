FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    DEBUG_LEVEL=TRACE \
    ENVIRONMENT=development

WORKDIR /build

#COPY go.mod go.sum main.go ./
COPY . ./

RUN go mod download

ARG GIT_SHORT_HASH
RUN echo "Build GitShortHash: $GIT_SHORT_HASH"
RUN echo "Build Date: `date +%Y/%m/%d/%H:%M`"

#RUN go build -o main .
RUN go build -v -o fcm-sender -ldflags "-X 'main.GitCommit=$GIT_SHORT_HASH' -X 'main.BuildTime=`date +%Y/%m/%d/%H:%M`'"

WORKDIR /dist

#RUN cp /build/main .
RUN cp /build/fcm-sender .

#FROM scratch
FROM alpine:3.16

#RUN apk add --no-cache ca-certificates && \
#    apk add --no-cache tcpdump && \
#    apk add --no-cache curl && \
#    apk add --no-cache busybox-extras \


#COPY --from=builder /dist/main .
#COPY --from=builder /dist/fcm-sender /fcm-sender
COPY --from=builder /dist/fcm-sender .


ENTRYPOINT ["/fcm-sender"]