# tart by building the application.
# Build em gologs com distroless
FROM golang:1.14 as builder

## install xz
RUN apt-get update && apt-get install -y \
    xz-utils \
&& rm -rf /var/lib/apt/lists/*
## install UPX
ADD https://github.com/upx/upx/releases/download/v3.94/upx-3.94-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.94-amd64_linux.tar.xz | \
    tar -xOf - upx-3.94-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx

WORKDIR /go/src/gologs
ENV GO111MODULE=on
COPY . .
#RUN go install -v ./...
RUN GOOS=linux go  build -ldflags="-s -w" -o gologs main.go
RUN upx --brute gologs
#RUN upx gologs
RUN cp gologs /go/bin/gologs
RUN ls -lh

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/gologs /
CMD ["/gologs"]