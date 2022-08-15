FROM golang:1.18.5 as go-base

ENV DOCKERIZE_VERSION="v0.6.1" \
  # Tini for Docker
  TINI_VERSION="v0.19.0"


# === Builder Image ===
FROM go-base as builder
# System packages
RUN apt update && apt upgrade -y && apt install --no-install-recommends -y \
  curl \
  wget

# Dockerize Installation
RUN wget "https://github.com/jwilder/dockerize/releases/download/${DOCKERIZE_VERSION}/dockerize-linux-amd64-${DOCKERIZE_VERSION}.tar.gz" \
  && tar -C /usr/local/bin -xzf "dockerize-linux-amd64-${DOCKERIZE_VERSION}.tar.gz" \
  && rm "dockerize-linux-amd64-${DOCKERIZE_VERSION}.tar.gz" \
  && dockerize --version

# Tini Installation
RUN wget -O /usr/local/bin/tini "https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini" \
  && chmod +x /usr/local/bin/tini \
  && tini --version


# === Base Image ===
FROM go-base as base

COPY --from=builder /usr/local/bin/dockerize /usr/local/bin/dockerize
COPY --from=builder /usr/local/bin/tini /usr/local/bin/tini

COPY ./scripts/entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

WORKDIR /url-shortener
COPY ./go.mod ./go.sum ./

RUN go mod download && go mod verify


# === Development Image ===
FROM base as development

RUN go install github.com/cosmtrek/air@latest && \
  go install github.com/swaggo/swag/cmd/swag@latest

ENTRYPOINT [ "tini", "--", "/docker-entrypoint.sh" ]


# === Production Image ===
FROM base as production

# Setup Permissions
RUN groupadd -r fiber && useradd -m -d /fiber -r -g fiber fiber \
  && chown fiber:fiber -R /url-shortener

COPY --chown=fiber:fiber . /url-shortener/

RUN go build -v -o /usr/local/bin/url-shortener . \
  && chmod +x /usr/local/bin/url-shortener

USER fiber

ENTRYPOINT [ "tini", "--", "/docker-entrypoint.sh" ]

CMD [ "sh", "./scripts/start-prod.sh" ]
