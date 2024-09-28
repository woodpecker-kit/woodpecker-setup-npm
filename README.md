[![ci](https://github.com/woodpecker-kit/woodpecker-setup-npm/workflows/ci/badge.svg)](https://github.com/woodpecker-kit/woodpecker-setup-npm/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/woodpecker-kit/woodpecker-setup-npm?label=go.mod)](https://github.com/woodpecker-kit/woodpecker-setup-npm)
[![GoDoc](https://godoc.org/github.com/woodpecker-kit/woodpecker-setup-npm?status.png)](https://godoc.org/github.com/woodpecker-kit/woodpecker-setup-npm)
[![goreportcard](https://goreportcard.com/badge/github.com/woodpecker-kit/woodpecker-setup-npm)](https://goreportcard.com/report/github.com/woodpecker-kit/woodpecker-setup-npm)

[![GitHub license](https://img.shields.io/github/license/woodpecker-kit/woodpecker-setup-npm)](https://github.com/woodpecker-kit/woodpecker-setup-npm)
[![codecov](https://codecov.io/gh/woodpecker-kit/woodpecker-setup-npm/branch/main/graph/badge.svg)](https://codecov.io/gh/woodpecker-kit/woodpecker-setup-npm)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/woodpecker-kit/woodpecker-setup-npm)](https://github.com/woodpecker-kit/woodpecker-setup-npm/tags)
[![GitHub release)](https://img.shields.io/github/v/release/woodpecker-kit/woodpecker-setup-npm)](https://github.com/woodpecker-kit/woodpecker-setup-npm/releases)

## for what

- this project used to woodpecker plugin

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/woodpecker-kit/woodpecker-setup-npm)](https://github.com/woodpecker-kit/woodpecker-setup-npm/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## Features

- [x] flag `npm-registry` to set custom npm registry, and support npm whoami check
- [x] args `npm-folder` to publish, which must containing `package.json`, generate `.npmrc` file will be here
- [x] support `npm-token` or `npm-username` and `npm-password` to config `.npmrc` file
    - [x] open `verdaccio-user-token-support` will use `npm-username` and `npm-password` to get token
- [x] support write `.npmrc` file at project folder with `package.json`
    - [x] auto generate `package.json` file `registries` setting `scope` at `.npmrc`
    - [x] also support `npm-scoped-list` to define scoped
- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

### workflow usage

- see [doc](doc/docs.md)

## Notice

- want dev this project, see [dev doc](doc/README.md)