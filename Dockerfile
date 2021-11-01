FROM ubuntu:18.04 AS runtime
WORKDIR /app
# GoReleaser will automatically generate the binary in the root directory
COPY /go-cli-prototype .
ENTRYPOINT ["./go-cli-prototype"]