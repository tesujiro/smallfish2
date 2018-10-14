FROM alpine
MAINTAINER tesujiro <tesujiro@gmail.com>
RUN echo "now building..."
ADD ./smallfish2web.linux /smallfish2web
ADD ./public /public
ADD ./ssl /ssl
EXPOSE 80
EXPOSE 443
ENTRYPOINT ["/smallfish2web"]
