# Static List Resource

A static list resource for [Concourse](https://github.com/concourse/concourse/). It iterates over items in a given list.

![build-status](https://ci.taydev.net/api/v1/teams/main/pipelines/static-list-resource/jobs/unit-tests/badge)

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
- `unique`: _(optional)_ Pass in a random string here. Use this if you have the exact same list in multiple resource declarations. Concourse will consider them the same if you have global resources enabled.

## Behavior

### `check`

- The first check returns the first item in the list
- Every check after will return nothing

### `in` / get step

- Always returns the passed in version
- Stores version in a file called `item`

### `out` / put step

- Returns the next item in the list via the implicit `get` step. Requires the previous item to be passed in.
- If no previous item is passed in then it returns the first item in the list
- **params**
  - `previous`: _(required)_ The path to the previous item (e.g. `my-list/item`)

## Examples

A simple list of strings example. Check `example.yml` for a complex list example.

In general, you should start with a `get` step and then end with a `put` step to ensure the next time the job runs it picks up the next item in the list. The `put` step is how you iterate over the list.

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
  - get: first
    resource: example-list
  - task: display-selected-item
    config:
      platform: linux
      image_resource:
        type: registry-image
        source: { repository: busybox }
      inputs:
      - name: fist
      run:
        path: cat
        args: ["first/item"]
  - put: next-item
    resource: example-list
    params:
      previous: first/item
  - task: display-selected-item
    config:
      platform: linux
      image_resource:
        type: registry-image
        source: { repository: busybox }
      inputs:
      - name: next-item
      run:
        path: cat
        args: ["next-item/item"]
```
