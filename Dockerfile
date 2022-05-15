FROM gcr.io/distroless/base-debian11
ENTRYPOINT ["/cloudflare-dns-updater"]
COPY cloudflare-dns-updater /
