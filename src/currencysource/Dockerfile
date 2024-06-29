FROM golang:1.22.3-alpine@sha256:7e788330fa9ae95c68784153b7fd5d5076c79af47651e992a3cdeceeb5dd1df0 AS builder

WORKDIR /src
COPY . .

# Skaffold passes in debug-oriented compiler flags
ARG SKAFFOLD_GO_GCFLAGS
RUN CGO_ENABLED=0 GOOS=linux go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /currencysource .

FROM scratch
COPY data /data
WORKDIR /src
COPY --from=builder /currencysource ./server
ENV GOTRACEBACK=single

EXPOSE 8080
ENTRYPOINT ["/src/server"]
