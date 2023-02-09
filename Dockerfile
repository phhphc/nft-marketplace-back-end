# Build stage
FROM golang:1.19 as builder
WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
RUN make build

# Image
FROM ubuntu:20.04 as runner
WORKDIR /app

# install root CA certificates
RUN apt-get update \
    && apt-get install -y ca-certificates \
    && rm -rf /var/cache/apt/archives /var/lib/apt/lists

ENV PORT=9090 ENV=Production
EXPOSE 9090
COPY --from=builder /src/bin/ ./
CMD [ "/app/marketplace" ]
