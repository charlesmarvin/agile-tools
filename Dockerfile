FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Create appuser.
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

# Copy the local package files to the container's workspace.
ADD app/ /build/app
ADD static /build/static
WORKDIR /build/app

# Fetch dependencies via go get
RUN go get -d -v

# Using go mod.
# RUN go mod download
# RUN go mod verify

# build go executable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o agiletools .

# build new image from scratch with only app executable
FROM scratch
# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
# Copy our static executable.
COPY --from=builder /build/app/agiletools /app/
COPY --from=builder /build/static /static/
WORKDIR /app

# Use an unprivileged user.
USER appuser:appuser
# Run the binary.
ENTRYPOINT ["./agiletools"]

# Port on which the service will be exposed.
EXPOSE 8080