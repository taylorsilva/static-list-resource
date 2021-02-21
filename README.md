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

## Behavior

### `check`

- first check will output the entire list
- every check after will select the next item in the list as the "latest version"

### `in` / get step

- always returns the passed in version
- stored in a file called `item` in json format

### `out` / put step

no-op

## Examples

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
