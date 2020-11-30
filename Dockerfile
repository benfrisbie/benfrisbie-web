###############################################################################
# Stage 1 - Build executable in golang alpine image
###############################################################################
FROM golang:1.15-alpine as builder

# Get dependencies first to save time on future builds
WORKDIR /app
COPY go.mod go.sum ./
ENV GO111MODULE=on
RUN go mod download && go mod verify

# Copy rest of project and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app.o

###############################################################################
# Stage 2 - Copy executable into base alpine image and run
###############################################################################
FROM alpine:latest
COPY --from=builder /app/app.o .
EXPOSE 8080
CMD ["./app.o"]