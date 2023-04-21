# sensu-http-ping Asset

Sensu Asset for HTTP Checks.

## Usage

| Arg Name | Short | Env Var | Default | Description |
| -------- | ----- | ------- | ------- | ----------- |
| `endpoint` | `e` | `HTTP_PING_ENDPOINT` | - | Required. Full address including Scheme, Host, URL and optional Port of the target Endpoint. |
| `method` | `m` | `HTTP_PING_METHOD` | `POST` | HTTP Method to use. |
| `payload` | `p` | - | | HTTP Request Body to submit. |
| `insecure` | `i` | - | `false` | Skip TLS Certificate validation. |
| `fail` | - | - | `false` | Exit with code 1 if a non-2xx HTTP Response Code has been received. Otherwise it will only exit with code 1 on malformed requests and network errors. |
| `timeout` | - | - | `30` | Timeout Seconds. |

### Examples

```sh
sensu-http-ping -e http://my-cool-endpoint -m GET
```

```sh
sensu-http-ping -e https://insecure-endpoint -m GET --insecure --fail --timeout 5
```

```sh
HTTP_PING_ENDPOINT=https://thing-to-ping.com sensu-http-ping
```

## Publish new Version

```sh
./publish_release.sh v1.x.y
```
