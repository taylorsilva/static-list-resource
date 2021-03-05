# Static List Resource

A static list resource for [Concourse](https://github.com/concourse/concourse/).

Add it to your pipeline:
```yaml
resource_types:
- name: static-list-resource
  type: registry-image
  source:
    repository: taylorsilva/static-list-resource
```

## Source Configuration

- `list`: _(required)_ A list of strings. Can be single or multi-line strings.

## Behavior

### `check`

- The first check returns the first item in the list
- every check after will return the next item in the list as the "latest version"
- You'll probably want to set the resource to `check_every: never` to avoid having the list shift on you

### `in` / get step

- always returns the passed in version
- stores version in a file called `item`

### `out` / put step

no-op

## Examples

A simple list of strings example. Check `example.yml` for a complex list example.

```yaml
resources:
- name: example-list
  type: static-list-resource
  source:
    list:
    - Kaladin
    - Shallan
    - Dalinar
    - Navani

jobs:
- name: list-job
  plan:
  - get: example-list
  - task: display-selected-item
    config:
      platform: linux
      image_resource:
        type: registry-image
        source: { repository: busybox }
      inputs:
      - name: example-list
      run:
        path: cat
        args: ["example-list/item"]
```
