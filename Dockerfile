FROM debian
COPY gredis ./
EXPOSE 6379 
VOLUME /tmp/rosedb_server
CMD ./gredis