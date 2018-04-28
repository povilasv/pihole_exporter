# pihole_exporter

[![License Apache 2][badge-license]](LICENSE)

Fork from `nlamirault/pihole_exporter`, this fixes metrics format changes in pi-hole v3.3. 
Also adds support for arm in Dockerfile.

Docker container for Raspberry Pi 3 (arm v7):

    docker pull povilasv/arm-pihole_exporter

Building:

    docker build -t povilasv/arm-pihole_exporter:latest -f Dockerfile_arm .

Note version number matches **pihole version** this exporter supports.

# Overview 

This Prometheus exporter check your [Pi-Hole](https://pi-hole.net/) statistics. Available metrics are :

-   Ads blocked
-   Domains blocked
-   DNS Queries
-   Top Ads
-   Top Queries
-   Top clients

## Grafana dashboard

You can import grafana dashboard from [https://grafana.com/dashboards/5855](https://grafana.com/dashboards/5855)

## Development

- Build tool:

    make build

- Launch unit tests:

    make test

## Docker Deployment

-   Build Image:

    docker build -t pihole-exporter .

-   Start Container

    docker run -d -p 9311:9311 pihole-exporter -pihole http://192.168.1.5

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).

## License

See [LICENSE](LICENSE) for the complete license.

## Changelog

A [changelog](ChangeLog.md) is available

[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat
