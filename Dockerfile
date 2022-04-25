FROM scratch
ENTRYPOINT ["/cloudflare-dns-updater"]
COPY cloudflare-dns-updater /