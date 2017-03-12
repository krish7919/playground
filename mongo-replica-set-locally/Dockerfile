FROM mongo:3.4.2
MAINTAINER krish7919@github.com
WORKDIR /
RUN apt update && apt install -y net-tools
COPY mongod.conf.template /etc/mongod.conf.template
COPY mongod_entrypoint/mongod_entrypoint /
VOLUME /data/db /data/configdb
ENTRYPOINT ["/mongod_entrypoint"]
