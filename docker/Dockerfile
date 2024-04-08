FROM alpine

LABEL version="v0.25.0" \
    maintainer="wang@rehiy.com"

ADD . /srv
RUN sh /srv/deploy

ENV WCF_ADDRESS=192.168.1.2:7601

ENTRYPOINT ["/srv/entrypoint"]

EXPOSE 7600
