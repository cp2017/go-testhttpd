FROM alpine

ENV HTTP_PORT=8080
ADD resources/bin/go-testhttpd_musl /usr/bin/go-testhttpd
ENTRYPOINT ["go-testhttpd", "--status-sequence"]
CMD ["200,200,200,500"]
