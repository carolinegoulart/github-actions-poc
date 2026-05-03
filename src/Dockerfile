# Stage 1: Build the binary
FROM golang:1.22.0-alpine3.19 AS builder
WORKDIR /build 
COPY . .
# CGO_ENABLED=0 creates a binary that doesn't depend on external libraries
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app 

# Stage 2: Final lightweight image
FROM alpine:3.19 AS final
# Create a non-root user
RUN adduser --uid 1000 --disabled-password appuser
USER appuser
WORKDIR /app
# Copy only the compiled binary from the builder stage, keeping the user permissions
COPY --from=builder --chown=appuser:appuser /build/app /go/app
# Expose the port the app will listen on
EXPOSE 5000
# Start the application
CMD [ "/go/app" ]