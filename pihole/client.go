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

package pihole

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	acceptHeader = "application/json"
	mediaType    = "application/json"
	Version      = "0.3.3"
)

var (
	userAgent = fmt.Sprintf("pihole-exporter/%s", Version)
)

type Client struct {
	Endpoint string
	Auth		 string
}

func NewClient(endpoint string, auth string) (*Client, error) {
	url, err := url.Parse(endpoint)
	if err != nil || url.Scheme != "http" {
		return nil, fmt.Errorf("Invalid PiHole address: %s", err)
	}
	return &Client{
		Endpoint: url.String(),
		Auth: auth,
	}, nil
}

func (c *Client) setupHeaders(request *http.Request) {
}

func (client *Client) GetMetrics() (*Metrics, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/admin/api.php?summaryRaw&overTimeData&topItems&recentItems&getQueryTypes&getForwardDestinations&getQuerySources&auth=%s", client.Endpoint, client.Auth), nil)

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", acceptHeader)
	req.Header.Add("User-Agent", userAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var metrics Metrics
	if err := json.NewDecoder(resp.Body).Decode(&metrics); err != nil {
		return nil, err
	}
	return &metrics, nil
}
