---
name: woodpecker-setup-npm
description: "woodpecker setup npm kits, add .npmrc or other npm config file at package.json dir"
author: woodpecker-kit
tags: [ node, npm ]
containerImage: sinlov/woodpecker-setup-npm
containerImageUrl: https://hub.docker.com/r/sinlov/woodpecker-setup-npm
url: https://github.com/woodpecker-kit/woodpecker-setup-npm
icon: https://raw.githubusercontent.com/woodpecker-kit/woodpecker-setup-npm/main/doc/logo.png
---

# woodpecker-setup-npm

woodpecker setup npm kits, add .npmrc or other npm config file at package.json dir

## Features

- [x] flag `npm-registry` to set custom npm registry, and support npm whoami check
- [x] args `npm-folder` to publish, which must containing `package.json`, generate `.npmrc` file will be here
- [x] support `npm-token` or `npm-username` and `npm-password` to config `.npmrc` file
    - [x] open `verdaccio-user-token-support` will use `npm-username` and `npm-password` to get token
- [x] support write `.npmrc` file at project folder with `package.json`
    - [x] auto generate `package.json` file `registries` setting `scope` at `.npmrc`
    - [x] also support `npm-scoped-list` to define scoped

## Settings

| Name                           | Required | Default value | Description                                                                                        |
|--------------------------------|----------|---------------|----------------------------------------------------------------------------------------------------|
| `debug`                        | **no**   | *false*       | open debug log or open by env `PLUGIN_DEBUG`                                                       |
| `npm-registry`                 | **no**   | *none*        | NPM registry settings if empty will use https://registry.npmjs.org/                                |
| `npm-token`                    | **yes**  | *none*        | NPM token to use when publishing packages. if token is set, username and password will be ignored. |
| `npm-username`                 | **yes**  | *none*        | NPM username                                                                                       |
| `npm-password`                 | **yes**  | *none*        | NPM password                                                                                       |
| `verdaccio-user-token-support` | **no**   | *false*       | use username and password to get token, only for https://verdaccio.org/                            |
| `npm-folder`                   | **no**   | *none*        | folder containing package.json, empty will use workspace                                           |
| `npm-dry-run`                  | **no**   | *false*       | dry run mode, will add NPM registry config but will print the command only                         |
| `npm-scoped-list`              | **no**   | *none*        | auto generate `package.json` node `registries` not finding can add this to append                  |

**Hide Settings:**

| Name                                        | Required | Default value                    | Description                                                                      |
|---------------------------------------------|----------|----------------------------------|----------------------------------------------------------------------------------|
| `timeout_second`                            | **no**   | *10*                             | command timeout setting by second                                                |
| `woodpecker-kit-steps-transfer-file-path`   | **no**   | `.woodpecker_kit.steps.transfer` | Steps transfer file path, default by `wd_steps_transfer.DefaultKitStepsFileName` |
| `woodpecker-kit-steps-transfer-disable-out` | **no**   | *false*                          | Steps transfer write disable out                                                 |

## Example

- workflow with backend `docker`

[![docker hub version semver](https://img.shields.io/docker/v/sinlov/woodpecker-setup-npm?sort=semver)](https://hub.docker.com/r/sinlov/woodpecker-setup-npm/tags?page=1&ordering=last_updated)
[![docker hub image size](https://img.shields.io/docker/image-size/sinlov/woodpecker-setup-npm)](https://hub.docker.com/r/sinlov/woodpecker-setup-npm)
[![docker hub image pulls](https://img.shields.io/docker/pulls/sinlov/woodpecker-setup-npm)](https://hub.docker.com/r/sinlov/woodpecker-setup-npm/tags?page=1&ordering=last_updated)

```yml
labels:
  backend: docker
steps:
  woodpecker-setup-npm:
    image: sinlov/woodpecker-setup-npm:latest
    pull: false
    settings:
      # debug: true
      ## registry settings if empty will use https://registry.npmjs.org/
      # npm-registry: https://verdaccio.foo.com
      # npm-token: # NPM token if token is set, username and password will be ignored
      #   from_secret: npm_setup_token
      npm-username: # NPM username
        from_secret: npm_setup_username
      npm-password: # NPM password
        from_secret: npm_setup_password
      verdaccio-user-token-support: true # use username and password to get token only for https://verdaccio.org/
      ## folder containing package.json, empty will use workspace
      npm-folder: .
```

- workflow with backend `local`, must install at local and effective at evn `PATH`
    - can download by [github release](https://github.com/woodpecker-kit/woodpecker-setup-npm/releases)
- install at ${GOPATH}/bin, latest

```bash
go install -a github.com/woodpecker-kit/woodpecker-setup-npm/cmd/woodpecker-setup-npm@latest
```

[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/woodpecker-kit/woodpecker-setup-npm)](https://github.com/woodpecker-kit/woodpecker-setup-npm/tags)
[![GitHub release)](https://img.shields.io/github/v/release/woodpecker-kit/woodpecker-setup-npm)](https://github.com/woodpecker-kit/woodpecker-setup-npm/releases)

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
      ## registry settings if empty will use https://registry.npmjs.org/
      # npm-registry: https://verdaccio.foo.com
      # npm-token: # NPM token if token is set, username and password will be ignored
      #   from_secret: npm_setup_token
      npm-username: # NPM username
        from_secret: npm_setup_username
      npm-password: # NPM password
        from_secret: npm_setup_password
      verdaccio-user-token-support: true # use username and password to get token
      ## folder containing package.json, empty will use workspace
      npm-folder: .
```

- full config

```yml
labels:
  backend: docker
steps:
  woodpecker-setup-npm:
    image: sinlov/woodpecker-setup-npm:latest
    pull: false
    settings:
      # debug: true
      ## registry settings if empty will use https://registry.npmjs.org/
      # npm-registry: https://verdaccio.foo.com
      npm-token: # NPM token if token is set, username and password will be ignored
        from_secret: npm_setup_token
      npm-username: # NPM username
        from_secret: npm_setup_username
      npm-password: # NPM password
        from_secret: npm_setup_password
      verdaccio-user-token-support: true # use username and password to get token only for https://verdaccio.org/
      npm-dry-run: true # dry run mode, will add NPM registry config but will print the command only
      ## folder containing package.json, empty will use workspace
      npm-folder: .
      npm-scoped-list: # auto generate `package.json` node `registries` not finding can add this to append
        - "verdaccio.foo.com" # scoped "@verdaccio.foo.com/*": "https://verdaccio.foo.com"
        - "pkg.bar" # scoped "@pkg.bar*": "https://verdaccio.foo.com"
```

## Notes

- use scoped list to auto generate `package.json` node `registries` which `package.json` can like this

```json
{
  "name": "foo",
  "version": "1.0.0",
  "registries": {
    "@verdaccio.foo.com/*": "https://verdaccio.foo.com",
    "@pkg.bar*": "https://verdaccio.foo.com"
  }
}
```
then generate `.npmrc` file like this

```ini
//verdaccio.foo.com/:_authToken={YOUR_TOKEN}
@verdaccio.foo.com:registry=https://verdaccio.foo.com
@pkg.bar:registry=https://verdaccio.foo.com
```

## Known limitations

- `verdaccio-user-token-support` only test for [verdaccio Version: 5.x](https://verdaccio.org/)
