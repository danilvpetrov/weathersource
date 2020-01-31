FROM alpine:latest as source
RUN apk --update add ca-certificates tzdata

FROM scratch
COPY --from=source /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=source /usr/share/zoneinfo /usr/share/zoneinfo


COPY artifacts/build/release/linux/amd64/weather /app/bin/
COPY artifacts/build/release/linux/amd64/inflateschema /app/bin/

ENTRYPOINT ["/app/bin/weather"]
