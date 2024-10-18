Huawei Cloud DNS module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with [Huawei Cloud DNS](https://www.huaweicloud.com/product/dns.html).

## Caddy module name

```
dns.providers.huaweicloud
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "huaweicloud",
				"access_key_id": "HUAWEICLOUD_ACCESS_KEY_ID",
				"secret_access_key": "HUAWEICLOUD_SECRET_ACCESS_KEY"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns huaweicloud {
		access_key_id {env.HUAWEICLOUD_ACCESS_KEY_ID}
		secret_access_key {env.HUAWEICLOUD_SECRET_ACCESS_KEY}
	}
}
```

```
# one site
tls {
  dns huaweicloud {
    access_key_id {env.HUAWEICLOUD_ACCESS_KEY_ID}
    secret_access_key {env.HUAWEICLOUD_SECRET_ACCESS_KEY}
  }
}
```