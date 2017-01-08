FROM alpine

ADD resources/bin/go-testhttpd_musl /usr/bin/go-testhttpd
ENTRYPOINT ["go-testhttpd"]
CMD ["--status-sequence", "200,201,400,500"]
