FROM golang:1.12.5-alpine AS builder

RUN apk add --no-cache ca-certificates git

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

RUN go get github.com/vektra/mockery/cmd/mockery

COPY ./ ./

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /message-matcher .

RUN go generate ./matcher

FROM scratch AS final

COPY --from=builder /src/.matcher.yml /src/messages.json /
COPY --from=builder /message-matcher /message-matcher

ENTRYPOINT ["/message-matcher"]
