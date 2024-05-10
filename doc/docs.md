---
name: woodpecker-setup-npm
description: "woodpecker setup npm kits, add .npmrc or other npm config file at node dir"
author: woodpecker-kit
tags: [ node, npm ]
containerImage: sinlov/woodpecker-setup-npm
containerImageUrl: https://hub.docker.com/r/sinlov/woodpecker-setup-npm
url: https://github.com/woodpecker-kit/woodpecker-setup-npm
icon: https://raw.githubusercontent.com/woodpecker-kit/woodpecker-setup-npm/main/doc/logo.png
---

woodpecker plugin template

## Settings

| Name                           | Required | Default value | Description                                        |
|--------------------------------|----------|---------------|----------------------------------------------------|
| `debug`                        | **no**   | *false*       | open debug log or open by env `PLUGIN_DEBUG`       |
| `not-empty-envs`               | **no**   | *none*        | use this args, will check env not empty            |
| `env-printer-print-keys`       | **no**   | *none*        | use this args, will print env by keys              |
| `env-printer-padding-left-max` | **no**   | *32*          | set env printer padding left max count, minimum 24 |
| `steps-transfer-demo`          | **no**   | *false*       | use this args, will print steps transfer demo      |

**Hide Settings:**

| Name                                        | Required | Default value                    | Description                                                                      |
|---------------------------------------------|----------|----------------------------------|----------------------------------------------------------------------------------|
| `timeout_second`                            | **no**   | *10*                             | command timeout setting by second                                                |
| `woodpecker-kit-steps-transfer-file-path`   | **no**   | `.woodpecker_kit.steps.transfer` | Steps transfer file path, default by `wd_steps_transfer.DefaultKitStepsFileName` |
| `woodpecker-kit-steps-transfer-disable-out` | **no**   | *false*                          | Steps transfer write disable out                                                 |

## Example

- workflow with backend `docker`

```yml
labels:
  backend: docker
steps:
  woodpecker-setup-npm:
    image: sinlov/woodpecker-setup-npm:latest
    pull: false
    settings:
      # debug: true
      # not-empty-envs: # check env not empty v1.7.+ support
      #   - WOODPECKER_AGENT_USER_HOME
      env-printer-print-keys: # print env keys
        - GOPATH
        - GOPRIVATE
        - GOBIN
        # env-printer-padding-left-max: # padding left max
        ## https://woodpecker-ci.org/docs/usage/secrets
        # from_secret: secret_printer_padding_left_max
      steps-transfer-demo: false # open this show steps transfer demo
```

- workflow with backend `local`, must install at local and effective at evn `PATH`
    - can download by [github release](https://github.com/woodpecker-kit/woodpecker-setup-npm/releases)
- install at ${GOPATH}/bin, latest

```bash
go install -a github.com/woodpecker-kit/woodpecker-setup-npm/cmd/woodpecker-setup-npm@latest
```

- install at ${GOPATH}/bin, v1.0.0

```bash
go install -v github.com/woodpecker-kit/woodpecker-setup-npm/cmd/woodpecker-setup-npm@v1.0.0
```

```yml
labels:
  backend: local
steps:
  woodpecker-setup-npm:
    image: woodpecker-setup-npm
    settings:
      # debug: true
      # not-empty-envs: # check env not empty v1.7.+ support
      #   - WOODPECKER_AGENT_USER_HOME
      env-printer-print-keys: # print env keys
        - GOPATH
        - GOPRIVATE
        - GOBIN
      env-printer-padding-left-max: 36 # padding left max
      steps-transfer-demo: false # open this show steps transfer demo
```

## Notes

- Please add notes

## Known limitations

- Please add a known issue
