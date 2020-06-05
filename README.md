# TLS Trigger Resource

A resource that triggers if a TLS certificate expires within the configured time.


## Source Configuration

* `domain`: The domain whose certificate to check. e.g. `storyscript.com`

* `expires_in`: The time duration in number of days. e.g. `7`

## Behavior

### `check`
Returns current version, and a new version only if the TLS certificate provided by `domain` expires in `expires_in` days.

### `in`
no-op

### `out`
no-op

## Example

```yaml
resource_types:
  - name: tls-trigger
    type: docker-image
    source:
      repository: anukulsangwan/tls-trigger-resource

resources:
  - name: storyscript-tls-trigger
    type: tls-trigger
    check_every: 24h
    source:
      domain: storyscript.com
      expires_in: 7
```
