# Image to get some root ca-certificates
FROM alpine:3.10 as certs
RUN apk --update add ca-certificates

# Build image
FROM golang:1.13.4 as go-builder

# API test & build
ENV API_NAME="project_emitter"
COPY . /build/project/${API_NAME}
WORKDIR /build/project/${API_NAME}

ENV CGO_ENABLED 0
# GOGC=off disables the garbage collector during compile time (with no effect on runtime)
ENV GOGC off
RUN go build -o main -ldflags="-s -w"
RUN mv main /build/main

# Final image
FROM scratch

USER 1000:1000
COPY --chown=1000:1000 --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --chown=1000:1000 --from=go-builder /build/main /main

EXPOSE 8080

CMD [ "/main" ]
