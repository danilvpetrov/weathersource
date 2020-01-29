FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY artifacts/build/release/linux/amd64/weather /app/bin/

ENTRYPOINT ["/app/bin/weather"]
