FROM golang:1.20 AS builder

WORKDIR /go/src/pixelart
COPY backend .

ENV CGO_ENABLED=0

RUN go mod download
RUN go build -o /go/bin/pixelart ./cmd/pixelart

FROM gcr.io/distroless/static

COPY --from=builder /go/bin/pixelart /
COPY frontend /frontend

EXPOSE 8080
CMD ["/pixelart"]
