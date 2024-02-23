FROM debian:stable-slim

RUN --mount=type=cache,sharing=private,target=/var/cache/apt \
    --mount=type=cache,sharing=private,target=/var/lib/apt \
    --mount=type=cache,sharing=private,target=/root/.cache \
    --mount=type=cache,sharing=private,target=/root/go/pkg \
    apt-get update && \
    apt-get install -y bash golang curl && \
    go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest && \
    go install github.com/bojand/ghz/cmd/ghz@latest && \
    apt-get remove -y --purge golang && \
    apt-get autoremove -y --purge

ENV PATH="$PATH:/root/go/bin"

CMD ["sleep", "infinity"]
