Valist CLI

[![oclif](https://img.shields.io/badge/cli-oclif-brightgreen.svg)](https://oclif.io)

<!-- toc -->
* [Usage](#usage)
* [Commands](#commands)
<!-- tocstop -->
# Usage
<!-- usage -->
```sh-session
$ npm install -g valist-cli
$ valist COMMAND
running command...
$ valist (-v|--version|version)
valist-cli/0.0.4 darwin-arm64 node-v14.17.1
$ valist --help [COMMAND]
USAGE
  $ valist COMMAND
...
```
<!-- usagestop -->
# Commands
<!-- commands -->
* [`valist address`](#valist-address)
* [`valist help [COMMAND]`](#valist-help-command)
* [`valist publish`](#valist-publish)

## `valist address`

print current signer address

```
USAGE
  $ valist address

EXAMPLE
  $ valist address
```

_See code: [src/commands/address.ts](https://github.com/valist-io/valist/blob/v0.0.4/src/commands/address.ts)_

## `valist help [COMMAND]`

display help for valist

```
USAGE
  $ valist help [COMMAND]

ARGUMENTS
  COMMAND  command to show help for

OPTIONS
  --all  see all commands in CLI
```

_See code: [@oclif/plugin-help](https://github.com/oclif/plugin-help/blob/v3.2.2/src/commands/help.ts)_

## `valist publish`

publish a package to valist

```
USAGE
  $ valist publish

EXAMPLE
  $ valist publish
```

_See code: [src/commands/publish.ts](https://github.com/valist-io/valist/blob/v0.0.4/src/commands/publish.ts)_
<!-- commandsstop -->
