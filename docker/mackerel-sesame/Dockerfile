FROM golang:1.21-bookworm as builder

WORKDIR /opt

COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o "mackerel-sesame" ./cmd/mackerel-sesame

FROM gcr.io/distroless/base-debian12:latest

COPY --from=builder /opt/mackerel-sesame /mackerel-sesame
CMD ["/mackerel-sesame"]
