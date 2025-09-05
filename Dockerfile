# -----------------------
# Build stage
# -----------------------
FROM public.ecr.aws/docker/library/golang:1.24 AS buildstage

WORKDIR /app

# Copy everything so we have Makefile, migrations, source
COPY . .

# Load env vars only for build step
RUN set -a && . /app/local.envrc && set +a && make build

# -----------------------
# Runtime stage
# -----------------------
FROM public.ecr.aws/docker/library/alpine:3.17.3

WORKDIR /app

# Install goose (needed for migrations)
RUN apk add --no-cache bash make curl git \
    && curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh | sh

# Copy binary + Makefile + migrations + envrc
COPY --from=buildstage /app/target/build ./build
COPY Makefile ./
COPY migrations ./migrations
COPY local.envrc ./
# # Copy wait-for script into runtime image
# COPY wait-for-db.sh ./wait-for-db.sh
# RUN chmod +x ./wait-for-db.sh
# Copy entrypoint script into runtime image
COPY entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh


# Set env vars
ENV GOMAXPROCS=1

# Default command (will be overridden by docker-compose)
CMD ["./build"]
