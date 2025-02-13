
FROM node:22-alpine AS build-frontend
WORKDIR /app

COPY ./frontend/package.json ./frontend/package-lock.json ./
RUN npm install --frozen-lockfile

COPY ./frontend .

RUN ls -la

RUN npm run build

FROM golang:1.23.6 AS build-backend
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY --from=build-frontend /app/dist ./frontend/dist

ENV ENV=prod

RUN go build -o ./bin/go .

FROM debian:bookworm-slim
WORKDIR /app

# RUN apk add --no-cache gcc musl-dev libc6-compat

# RUN apt-get update && apt-get install -y --no-install-recommends \
#     libc6 libgcc-s1 libstdc++6 \
#     && rm -rf /var/lib/apt/lists/*

COPY --from=build-backend /app/bin/go /usr/bin/go

VOLUME /app/data
EXPOSE 1323

CMD ["/usr/bin/go"]