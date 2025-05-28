# Authzed API

[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg "Apache 2.0 License")](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Docs](https://img.shields.io/badge/docs-authzed.com-%234B4B6C "Authzed Documentation")](https://authzed.com/docs)
[![Build Status](https://github.com/authzed/api/workflows/Lint/badge.svg "GitHub Actions")](https://github.com/authzed/api/actions)
[![Discord Server](https://img.shields.io/discord/844600078504951838?color=7289da&logo=discord "Discord Server")](https://authzed.com/discord)
[![Twitter](https://img.shields.io/twitter/follow/authzed?color=%23179CF0&logo=twitter&style=flat-square "@authzed on Twitter")](https://twitter.com/authzed)

This project contains the definitions of [Protocol Buffers] used by Authzed.

We use [Buf] to distribute these definitions and generate source code from them. The definitions are published in [Buf Registry Authzed API repository].

You can find more info on [HTTP API usage] and the [versioning and deprecation policy] in the [Authzed Docs].

You can also use our [Postman collection] to explore the API.

See [CONTRIBUTING.md] for instructions on how to contribute.

[Protocol Buffers]: https://developers.google.com/protocol-buffers/
[Buf]: https://github.com/bufbuild/buf
[HTTP API usage]: https://authzed.com/docs/spicedb/getting-started/client-libraries#http-clients
[Authzed Docs]: https://authzed.com/docs
[versioning and deprecation policy]: https://authzed.com/blog/buf
[Postman collection]: (https://www.postman.com/authzed/spicedb/collection/m26cqyc)
[Buf Registry Authzed API repository]: https://buf.build/authzed/api/docs/main
[CONTRIBUTING.md]: https://github.com/authzed/api/blob/main/CONTRIBUTING.md

## Development

You can run `mage` to see the available commands for development. We assume you have a Mac computer.

## Warnings ⚠️

- The `version` field found in various buf YAML configuration is actually schema of the YAML of the file and is not related to the version of the definitions.
