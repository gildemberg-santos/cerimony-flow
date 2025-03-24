FROM node:20 AS frontend-builder
WORKDIR /frontend

COPY ./frontend ./

RUN npm install && npm run build

FROM golang:1.24.1-alpine AS backend-builder
WORKDIR /app

COPY ./backend/main.go ./
COPY ./backend/go.mod ./
COPY ./backend/go.sum ./

COPY --from=frontend-builder /frontend/build ./frontend/build

RUN go mod tidy && go build -o main .

FROM alpine:latest
WORKDIR /

COPY --from=backend-builder /app/main /main

COPY --from=backend-builder /app/frontend/build ./

EXPOSE 8080

CMD ["/main"]
