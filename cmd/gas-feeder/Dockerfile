FROM jrei/systemd-centos:7

WORKDIR /app

COPY gas-feeder /app
COPY GasFeeder.viper.yaml /app

ENTRYPOINT ["./gas-feeder"]

EXPOSE 32700 32701 32702

CMD ["run"]
