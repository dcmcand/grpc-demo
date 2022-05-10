
FROM golang:1.18-buster as builder
COPY . /app

WORKDIR /app
RUN go mod verify && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/grpc-poc server/main.go

ENV USER=appuser
ENV UID=10001 

RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

FROM scratch
ARG PORT=9000
ENV PORT="${PORT}"
EXPOSE "${PORT}/tcp"
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder --chown=appuser:appuser /app/grpc-poc /app/server
USER appuser
ADD server-test /app/
ENTRYPOINT [ "/app/server" ]