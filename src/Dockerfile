FROM golang:latest
ENV REFRESHED_AT 2017-12-19
WORKDIR /mnt
COPY ./my_first /mnt/my_first
ADD start.sh /mnt/start.sh
RUN chmod +x /mnt/start.sh

ENV START_ENV "dev"

ENTRYPOINT ["sh"]
CMD ["/mnt/start.sh"]