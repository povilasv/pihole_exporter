// Copyright (C) 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/povilasv/pihole_exporter/pihole"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

const (
	banner    = "pihole_exporter - %s\n"
	namespace = "pihole"
)

var (
	debug               bool
	version             bool
	listenAddress       string
	metricsPath         string
	endpoint            string
	logLevel            string
	logFormat           string
	domainsBeingBlocked = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "domains_being_blocked"),
		"Domains being blocked.",
		nil, nil,
	)
	dnsQueries = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "dns_queries_today"),
		"DNS Queries today.",
		nil, nil,
	)
	adsBlocked = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "ads_blocked_today"),
		"Ads blocked today.",
		nil, nil,
	)

	adsPercentage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "ads_percentage_today"),
		"Ads percentage today.",
		nil, nil,
	)

	domainsOverTime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "domains_over_time"),
		"Domains over time.",
		nil, nil,
	)

	adsOverTime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "ads_over_time"),
		"Ads over time.",
		nil, nil,
	)

	topQueries = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "top_queries"),
		"Top queries.",
		[]string{"domain"}, nil,
	)

	topAds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "top_ads"),
		"Top Ads.",
		[]string{"domain"}, nil,
	)
	topSources = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "top_sources"),
		"Top sources.",
		[]string{"client"}, nil,
	)
	queryTypes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "query_types"),
		"DNS Query types.",
		[]string{"type"}, nil,
	)
)

// Exporter collects Pihole stats from the given server and exports them using
// the prometheus metrics package.
type Exporter struct {
	Pihole *pihole.Client
}

// NewExporter returns an initialized Exporter.
func NewExporter(endpoint string) (*Exporter, error) {
	log.Infoln("Setup Pihole exporter using URL: %s", endpoint)
	pihole, err := pihole.NewClient(endpoint)
	if err != nil {
		return nil, err
	}
	return &Exporter{
		Pihole: pihole,
	}, nil
}

// Describe describes all the metrics ever exported by the Pihole exporter.
// It implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- domainsBeingBlocked
	ch <- dnsQueries
	ch <- adsBlocked
	ch <- adsPercentage
	ch <- domainsOverTime
	ch <- adsOverTime
	ch <- topQueries
	ch <- topAds
	ch <- topSources
	ch <- queryTypes
}

// Collect the stats from channel and delivers them as Prometheus metrics.
// It implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	resp, err := e.Pihole.GetMetrics()
	if err != nil {
		log.Errorf("Pihole error: %s", err.Error())
		return
	}
	log.Debugf("PiHole metrics: %#v", resp)
	ch <- prometheus.MustNewConstMetric(
		domainsBeingBlocked, prometheus.CounterValue, float64(resp.DomainsBeingBlocked))

	ch <- prometheus.MustNewConstMetric(
		dnsQueries, prometheus.CounterValue, float64(resp.DNSQueriesToday))

	ch <- prometheus.MustNewConstMetric(
		adsBlocked, prometheus.CounterValue, float64(resp.AdsBlockedToday))

	ch <- prometheus.MustNewConstMetric(
		adsPercentage, prometheus.CounterValue, float64(resp.AdsPercentageToday))

	for k, v := range resp.Querytypes {
		ch <- prometheus.MustNewConstMetric(
			queryTypes, prometheus.CounterValue, v, k)
	}
	for k, v := range resp.TopQueries {
		ch <- prometheus.MustNewConstMetric(
			topQueries, prometheus.CounterValue, float64(v), k)
	}
	for k, v := range resp.TopAds {
		ch <- prometheus.MustNewConstMetric(
			topAds, prometheus.CounterValue, float64(v), k)

	}
	for k, v := range resp.TopSources {
		ch <- prometheus.MustNewConstMetric(
			topSources, prometheus.CounterValue, float64(v), k)
	}
}

func init() {
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.StringVar(&listenAddress, "web.listen-address", ":9311", "Address to listen on for web interface and telemetry.")
	flag.StringVar(&metricsPath, "web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	flag.StringVar(&endpoint, "pihole", "", "Endpoint of Pihole")
	flag.StringVar(&logLevel, "log.level", "info", "Only log messages with the given severity or above. Valid levels: [debug, info, warn, error, fatal]")
	flag.StringVar(&logFormat, "log.format", "logger:stderr", `Set the log target and format. Example: "logger:syslog?appname=bob&local=7" or "logger:stdout?json=true"`)
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(banner, version))
		flag.PrintDefaults()
	}

	flag.Parse()

	if version {
		fmt.Printf("%s", pihole.Version)
		os.Exit(0)
	}
	if logLevel != "" {
		if err := log.Base().SetLevel(logLevel); err != nil {
			log.Errorf("Failed to set log level: %s", err.Error())
			os.Exit(1)
		}

	}
	if logFormat != "" {
		if err := log.Base().SetFormat(logFormat); err != nil {
			log.Errorf("Failed to set log format: %s", err.Error())
			os.Exit(1)
		}
	}

	if len(endpoint) == 0 {
		usageAndExit("Pihole endpoint cannot be empty.", 1)
	}
}

func main() {
	exporter, err := NewExporter(endpoint)
	if err != nil {
		log.Errorf("Can't create exporter : %s", err)
		os.Exit(1)
	}
	log.Infoln("Register exporter")
	prometheus.MustRegister(exporter)

	http.Handle(metricsPath, prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Pihole Exporter</title></head>
             <body>
             <h1>Pihole Exporter</h1>
             <p><a href='` + metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})

	log.Infoln("Listening on", listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}

func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n")
	}
	flag.Usage()
	os.Exit(exitCode)
}

func storeMetric(ch chan<- prometheus.Metric, value string, desc *prometheus.Desc, labels ...string) {
	if val, err := strconv.ParseFloat(value, 64); err == nil {
		ch <- prometheus.MustNewConstMetric(
			desc, prometheus.GaugeValue, val, labels...)
	} else {
		log.Errorf("Can't store metric %s into %s: %s", value, desc, err.Error())
	}
}
