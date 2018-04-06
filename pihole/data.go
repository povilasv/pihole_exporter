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

// This the package for the Pi HOLE API
// See: https://github.com/pi-hole/AdminLTE

package pihole

type Metrics struct {
	DomainsBeingBlocked int                `json:"domains_being_blocked"`
	DNSQueriesToday     int                `json:"dns_queries_today"`
	AdsBlockedToday     int                `json:"ads_blocked_today"`
	AdsPercentageToday  float64            `json:"ads_percentage_today"`
	UniqueDomains       int                `json:"unique_domains"`
	QueriesForwarded    int                `json:"queries_forwarded"`
	QueriesCached       int                `json:"queries_cached"`
	ClientsEverSeen     int                `json:"clients_ever_seen"`
	UniqueClients       int                `json:"unique_clients"`
	Status              string             `json:"status"`
	TopQueries          map[string]int     `json:"top_queries"`
	TopAds              map[string]int     `json:"top_ads"`
	TopSources          map[string]int     `json:"top_sources"`
	ForwardDestinations map[string]float64 `json:"forward_destinations"`
	Querytypes          map[string]float64 `json:"querytypes"`
}
