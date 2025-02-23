# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:1.21 as builder
ARG GRPC_HEALTH_PROBE_VERSION=v0.4.21
ARG TARGETARCH
WORKDIR /workspace
RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-${TARGETARCH} && chmod +x /bin/grpc_health_probe

FROM gcr.io/distroless/static:latest
ARG MAIN_ENTRY
COPY --from=baseimg /${MAIN_ENTRY} /
COPY --from=builder /bin/grpc_health_probe /usr/local/bin/grpc_health_probe
ENTRYPOINT [/${MAIN_ENTRY}]
