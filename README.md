# pihole_exporter

[![License Apache 2][badge-license]](LICENSE)

Fork from `nlamirault/pihole_exporter`, this fixes metrics format changes in pi-hole v3.3. 
Also adds support for arm in Dockerfile.

Docker container for Raspberry Pi 3 (arm v7):

    docker pull povilasv/arm-pihole_exporter

Building:

    docker build -t povilasv/arm-pihole_exporter:latest -f Dockerfile_arm .

Note version number matcher **pihole version** this exporter supports.

# Overview 

This Prometheus exporter check your [Pi-Hole](https://pi-hole.net/) statistics. Available metrics are :

-   Ads blocked
-   Domains blocked
-   DNS Queries
-   Top Ads
-   Top Queries
-   Top clients

## Development

-   Build tool :

          $ make build

-   Launch unit tests :

          $ make test

## Local Deployment

-   Launch Prometheus using the configuration file in this repository:

          $ prometheus -config.file=prometheus.yml

-   Launch exporter:

          $ pihole_exporter -log.level=debug

-   Check that Prometheus find the exporter on `http://localhost:9090/targets`

-   Run Grafana and import the dashboard _dashboard.json_:

          $ docker run -d --name=grafana -p 3000:3000 grafana/grafana

## Docker Deployment

-   Build Image:

        		$ docker build -t pihole-exporter .

-   Start Container
        		$ docker run -d -p 9311:9311 pihole-exporter -pihole http://192.168.1.5

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).

## License

See [LICENSE](LICENSE) for the complete license.

## Changelog

A [changelog](ChangeLog.md) is available

## Contact

Nicolas Lamirault <mailto:nicolas.lamirault@gmail.com>

[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat
