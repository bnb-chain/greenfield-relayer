FROM golang:1.19-alpine as builder

# Set up apk dependencies
ENV PACKAGES make git libc-dev bash gcc linux-headers eudev-dev curl ca-certificates build-base

# Set working directory for the build
WORKDIR /opt/app

# Add source files
COPY . .

# Install minimum necessary dependencies, remove packages
RUN apk add --no-cache $PACKAGES

# For Private REPO
ARG GH_TOKEN=""
RUN go env -w GOPRIVATE="github.com/bnb-chain/*"
RUN git config --global url."https://${GH_TOKEN}@github.com".insteadOf "https://github.com"

RUN make build

# Pull binary into a second stage deploy alpine container
FROM alpine:3.17

ARG USER=app
ARG USER_UID=1000
ARG USER_GID=1000

ENV GREENFIELD_RELAYER_HOME /opt/app
ENV CONFIG_FILE_PATH $GREENFIELD_RELAYER_HOME/config/config.json
ENV CONFIG_TYPE "local"
ENV PRIVATE_KEY ""
ENV BLS_PRIVATE_KEY ""
ENV DB_PASS ""
# You need to specify aws s3 config if you want to load config from s3
ENV AWS_REGION ""
ENV AWS_SECRET_KEY ""

ENV PACKAGES ca-certificates bash curl libstdc++
ENV WORKDIR=/app

RUN apk add --no-cache $PACKAGES \
  && rm -rf /var/cache/apk/* \
  && addgroup -g ${USER_GID} ${USER} \
  && adduser -u ${USER_UID} -G ${USER} --shell /sbin/nologin --no-create-home -D ${USER} \
  && addgroup ${USER} tty \
  && sed -i -e "s/bin\/sh/bin\/bash/" /etc/passwd

RUN echo "[ ! -z \"\$TERM\" -a -r /etc/motd ] && cat /etc/motd" >> /etc/bash/bashrc

WORKDIR ${WORKDIR}

COPY --from=builder /opt/app/build/greenfield-relayer ${WORKDIR}/
RUN chown -R ${USER_UID}:${USER_GID} ${WORKDIR}
USER ${USER_UID}:${USER_GID}

VOLUME [ $GREENFIELD_RELAYER_HOME ]

# Run the app
CMD /app/greenfield-relayer --config-type $CONFIG_TYPE --config-path $CONFIG_FILE_PATH --private-key $PRIVATE_KEY --bls-private-key $BLS_PRIVATE_KEY -db-pass $DB_PASS --aws-region $AWS_REGION --aws-secret-key $AWS_SECRET_KEY