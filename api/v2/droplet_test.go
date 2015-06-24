package apiv2

import (
	"encoding/json"
	"testing"
)

func TestNetworkV4JsonUnmarshall(t *testing.T) {
	j := `
          {
            "ip_address": "46.101.166.213",
            "netmask": "255.255.192.0",
            "gateway": "46.101.128.1",
            "type": "public"
          }
      `
	net := &Network{}
	err := json.Unmarshal([]byte(j), net)
	if err != nil {
		t.Error(err)
	}
}

func TestNetworkV6JsonUnmarshall(t *testing.T) {
	j := `
      {
        "ip_address": "2A03:B0C0:0002:00D0:0000:0000:01A3:9001",
        "netmask": 64,
        "gateway": "2A03:B0C0:0002:00D0:0000:0000:0000:0001",
        "type": "public"
      }
	`
	net := &Network{}
	err := json.Unmarshal([]byte(j), net)
	if err != nil {
		t.Error(err)
	}
}

func TestSizeUnmarshall(t *testing.T) {
	j := `{
    "slug": "512mb",
    "memory": 512,
    "vcpus": 1,
    "disk": 20,
    "transfer": 1,
    "price_monthly": 5,
    "price_hourly": 0.00744,
    "regions": [
      "nyc1",
      "sgp1",
      "ams1",
      "sfo1",
      "nyc2",
      "lon1",
      "nyc3",
      "ams3",
      "ams2",
      "fra1"
    ],
    "available": true
    }`
	obj := &Size{}
	err := json.Unmarshal([]byte(j), obj)
	if err != nil {
		t.Error(err)
	}
}

func TestRegionUnmarshall(t *testing.T) {
	j := `{
		"name": "Frankfurt 1",
        "slug": "fra1",
        "sizes": [
          "32gb",
          "16gb",
          "2gb",
          "1gb",
          "4gb",
          "8gb",
          "512mb",
          "64gb",
          "48gb"
        ],
        "features": [
          "private_networking",
          "backups",
          "ipv6",
          "metadata"
        ],
        "available": true

      }`
	obj := &Region{}
	err := json.Unmarshal([]byte(j), obj)
	if err != nil {
		t.Error(err)
	}
}

func TestImageUnmarshall(t *testing.T) {
	j := `{
    "id": 6918990,
    "name": "14.04 x64",
    "distribution": "Ubuntu",
    "slug": null,
    "public": false,
    "regions": [
      "nyc1",
      "ams1",
      "sfo1",
      "nyc2",
      "ams2",
      "sgp1",
      "lon1",
      "nyc3",
      "ams3",
      "fra1"
    ],
    "created_at": "2014-10-17T20:24:33Z",
    "min_disk_size": 20,
    "type": "snapshot"
    }
	`
	obj := &Image{}
	err := json.Unmarshal([]byte(j), obj)
	if err != nil {
		t.Error(err)
	}
}

func TestDropletUnmarshall1(t *testing.T) {
	j := `
{
  "id": 2092639,
  "name": "lalyos",
  "memory": 512,
  "vcpus": 1,
  "disk": 20,
  "locked": false,
  "status": "active",
  "kernel": {
    "id": 1221,
    "name": "Ubuntu 14.04 x64 vmlinuz-3.13.0-24-generic (1221)",
    "version": "3.13.0-24-generic"
  },
  "created_at": "2014-07-15T17:17:31Z",
  "features": [
    "virtio"
  ],
  "backup_ids": [],
  "next_backup_window": null,
  "snapshot_ids": [],
  "image": {
    "id": 3240036,
    "name": "Ubuntu 14.04 x64",
    "distribution": "Ubuntu",
    "slug": null,
    "public": false,
    "regions": [
      "nyc1",
      "ams1",
      "sfo1",
      "nyc2",
      "ams2",
      "sgp1",
      "lon1",
      "nyc3",
      "ams3",
      "fra1"
    ],
    "created_at": "2014-04-18T15:59:36Z",
    "min_disk_size": 20,
    "type": "snapshot"
  },
  "size": {
    "slug": "512mb",
    "memory": 512,
    "vcpus": 1,
    "disk": 20,
    "transfer": 1,
    "price_monthly": 5,
    "price_hourly": 0.00744,
    "regions": [
      "nyc1",
      "sgp1",
      "ams1",
      "sfo1",
      "nyc2",
      "lon1",
      "nyc3",
      "ams3",
      "ams2",
      "fra1"
    ],
    "available": true
  },
  "size_slug": "512mb",
  "networks": {
    "v4": [
      {
        "ip_address": "178.62.59.231",
        "netmask": "255.255.192.0",
        "gateway": "178.62.0.1",
        "type": "public"
      }
    ],
    "v6": []
  },
  "region": {
    "name": "London 1",
    "slug": "lon1",
    "sizes": [
      "32gb",
      "16gb",
      "2gb",
      "1gb",
      "4gb",
      "8gb",
      "512mb",
      "64gb",
      "48gb"
    ],
    "features": [
      "private_networking",
      "backups",
      "ipv6",
      "metadata"
    ],
    "available": true
  }
}`
	obj := &Droplet{}
	err := json.Unmarshal([]byte(j), obj)
	if err != nil {
		t.Error(err)
	}
}

func TestDrolpletUnmarshall2(t *testing.T) {
	j := `
{
  "id": 3232708,
  "name": "ipv6",
  "memory": 512,
  "vcpus": 1,
  "disk": 20,
  "locked": false,
  "status": "active",
  "kernel": {
    "id": 2233,
    "name": "Ubuntu 14.04 x64 vmlinuz-3.13.0-37-generic",
    "version": "3.13.0-37-generic"
  },
  "created_at": "2014-11-21T06:24:46Z",
  "features": [
    "ipv6",
    "virtio"
  ],
  "backup_ids": [],
  "next_backup_window": null,
  "snapshot_ids": [],
  "image": {
    "id": 6918990,
    "name": "14.04 x64",
    "distribution": "Ubuntu",
    "slug": null,
    "public": false,
    "regions": [
      "nyc1",
      "ams1",
      "sfo1",
      "nyc2",
      "ams2",
      "sgp1",
      "lon1",
      "nyc3",
      "ams3",
      "fra1"
    ],
    "created_at": "2014-10-17T20:24:33Z",
    "min_disk_size": 20,
    "type": "snapshot"
  },
  "size": {
    "slug": "512mb",
    "memory": 512,
    "vcpus": 1,
    "disk": 20,
    "transfer": 1,
    "price_monthly": 5,
    "price_hourly": 0.00744,
    "regions": [
      "nyc1",
      "sgp1",
      "ams1",
      "sfo1",
      "nyc2",
      "lon1",
      "nyc3",
      "ams3",
      "ams2",
      "fra1"
    ],
    "available": true
  },
  "size_slug": "512mb",
  "networks": {
    "v4": [
      {
        "ip_address": "128.199.54.228",
        "netmask": "255.255.224.0",
        "gateway": "128.199.32.1",
        "type": "public"
      }
    ],
    "v6": [
      {
        "ip_address": "2A03:B0C0:0002:00D0:0000:0000:01A3:9001",
        "netmask": 64,
        "gateway": "2A03:B0C0:0002:00D0:0000:0000:0000:0001",
        "type": "public"
      }
    ]
  },
  "region": {
    "name": "Amsterdam 3",
    "slug": "ams3",
    "sizes": [
      "32gb",
      "16gb",
      "2gb",
      "1gb",
      "4gb",
      "8gb",
      "512mb",
      "64gb",
      "48gb"
    ],
    "features": [
      "private_networking",
      "backups",
      "ipv6",
      "metadata"
    ],
    "available": true
  }
}
	`
	obj := &Droplet{}
	err := json.Unmarshal([]byte(j), obj)
	if err != nil {
		t.Logf("unmarshall error: %#v \n", err)
		if jerr, ok := err.(*json.UnmarshalTypeError); ok {
			t.Log("type error, on value:", jerr.Value)
		}
		t.Error(err)
	}
}
