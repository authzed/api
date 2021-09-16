# Authzed API

[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg "Apache 2.0 License")](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Build Status](https://github.com/authzed/api/workflows/Build/badge.svg "GitHub Actions")](https://github.com/authzed/api/actions)
[![Mailing List](https://img.shields.io/badge/email-google%20groups-4285F4 "authzed-oss@googlegroups.com")](https://groups.google.com/g/authzed-oss)
[![Discord Server](https://img.shields.io/discord/844600078504951838?color=7289da&logo=discord "Discord Server")](https://discord.gg/jTysUaxXzM)
[![Twitter](https://img.shields.io/twitter/follow/authzed?color=%23179CF0&logo=twitter&style=flat-square "@authzed on Twitter")](https://twitter.com/authzed)

This project contains the definitions of [Protocol Buffers] used by Authzed.

[Buf] is used to lint and distribute these definitions and generate source code from them.

You can find more info on each API on the [Authzed API reference documentation].
Additionally, Protobuf API documentation can be found on the [Buf Registry Authzed API repository].

See [CONTRIBUTING.md] for instructions on how to contribute and perform common tasks like building the project and running tests.

[Protocol Buffers]: https://developers.google.com/protocol-buffers/
[Buf]: https://github.com/bufbuild/buf
[Authzed API Reference documentation]: https://docs.authzed.com/reference/api
[Buf Registry Authzed API repository]: https://buf.build/authzed/api/docs/main
[CONTRIBUTING.md]: CONTRIBUTING.md

## ⚠️ Warnings ⚠️

- The `version` field found in various buf YAML configuration is actually schema of the YAML of the file and is not related to the version of the definitions.
- `buf build` and `buf generate` do entirely different things.
   Building compiles definitions and ensures semantic validity.
   Generate builds and then produces actual source code according to `buf.gen.yaml`.
