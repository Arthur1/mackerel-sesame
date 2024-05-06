FROM --platform=$BUILDPLATFORM golang:1.22-bookworm as builder
ARG TARGETARCH

WORKDIR /opt

COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o "mackerel-sesame" ./cmd/mackerel-sesame

FROM gcr.io/distroless/base-debian12:nonroot

COPY --from=builder --chown=nonroot:nonroot /opt/mackerel-sesame /mackerel-sesame
CMD ["/mackerel-sesame"]
