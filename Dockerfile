FROM alpine
COPY discli /bin/
ENTRYPOINT ["discli"]