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
valist-cli/0.4.1 darwin-x64 node-v14.17.2
$ valist --help [COMMAND]
USAGE
  $ valist COMMAND
...
```
<!-- usagestop -->
# Commands
<!-- commands -->
* [`valist account:get`](#valist-accountget)
* [`valist account:new`](#valist-accountnew)
* [`valist build`](#valist-build)
* [`valist help [COMMAND]`](#valist-help-command)
* [`valist init`](#valist-init)
* [`valist org:get ORGNAME`](#valist-orgget-orgname)
* [`valist org:key ORGNAME OPERATION KEY`](#valist-orgkey-orgname-operation-key)
* [`valist org:new ORGNAME`](#valist-orgnew-orgname)
* [`valist org:update ORGNAME ORGMETA`](#valist-orgupdate-orgname-orgmeta)
* [`valist publish`](#valist-publish)
* [`valist repo:get ORGNAME REPONAME`](#valist-repoget-orgname-reponame)
* [`valist repo:key ORGNAME REPONAME OPERATION KEY`](#valist-repokey-orgname-reponame-operation-key)
* [`valist repo:new ORGNAME REPONAME`](#valist-reponew-orgname-reponame)
* [`valist repo:update ORGNAME REPONAME REPOMETA`](#valist-repoupdate-orgname-reponame-repometa)

## `valist account:get`

Print account info

```
USAGE
  $ valist account:get

EXAMPLE
  $ valist account:get
```

_See code: [dist/commands/account/get.ts](https://github.com/valist-io/valist/blob/v0.4.1/dist/commands/account/get.ts)_

## `valist account:new`

Create a new account

```
USAGE
  $ valist account:new

EXAMPLE
  $ valist account:new
```

_See code: [dist/commands/account/new.ts](https://github.com/valist-io/valist/blob/v0.4.1/dist/commands/account/new.ts)_

## `valist build`

Build the target valist project

```
USAGE
  $ valist build

EXAMPLE
  $ valist build
```

_See code: [dist/commands/build.ts](https://github.com/valist-io/valist/blob/v0.4.1/dist/commands/build.ts)_

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

## `valist init`

Generate a new valist project

```
USAGE
  $ valist init

EXAMPLE
  $ valist init
```

_See code: [dist/commands/init.ts](https://github.com/valist-io/valist/blob/v0.4.1/dist/commands/init.ts)_

## `valist org:get ORGNAME`

print organization info

```
USAGE
  $ valist org:get ORGNAME

EXAMPLE
  $ valist org:get valist
```

_See code: [dist/commands/org/get.ts](https://github.com/valist-io/valist/blob/v0.4.1/dist/commands/org/get.ts)_

## `valist org:key ORGNAME OPERATION KEY`

Add, remove, or rotate organization key

```
USAGE
  $ valist org:key ORGNAME OPERATION KEY

EXAMPLES
  $ valist org:key exampleOrg grant <key>
  $ valist org:key exampleOrg revoke <key>
  $ valist org:key exampleOrg rotate <key>
```

_See code: [dist/commands/org/key.ts](https://github.com/valist-io/valist/blob/v0.4.1/dist/commands/org/key.ts)_

## `valist org:new ORGNAME`

Create a Valist organization

```
USAGE
  $ valist org:new ORGNAME

EXAMPLE
  $ valist org:new exampleOrg
```

_See code: [dist/commands/org/new.ts](https://github.com/valist-io/valist/blob/v0.4.1/dist/commands/org/new.ts)_

## `valist org:update ORGNAME ORGMETA`

Update organization metadata

```
USAGE
  $ valist org:update ORGNAME ORGMETA

EXAMPLE
  $ valist org:update exampleOrg meta/orgMeta.json
```

_See code: [dist/commands/org/update.ts](https://github.com/valist-io/valist/blob/v0.4.1/dist/commands/org/update.ts)_

## `valist publish`

Publish a package to Valist

```
USAGE
  $ valist publish

EXAMPLE
  $ valist publish
```

_See code: [dist/commands/publish.ts](https://github.com/valist-io/valist/blob/v0.4.1/dist/commands/publish.ts)_

## `valist repo:get ORGNAME REPONAME`

print organization info

```
USAGE
  $ valist repo:get ORGNAME REPONAME

EXAMPLE
  $ valist repo:get exampleOrg exampleRepo
```

_See code: [dist/commands/repo/get.ts](https://github.com/valist-io/valist/blob/v0.4.1/dist/commands/repo/get.ts)_

## `valist repo:key ORGNAME REPONAME OPERATION KEY`

Add, remove, or rotate repository key

```
USAGE
  $ valist repo:key ORGNAME REPONAME OPERATION KEY

EXAMPLES
  $ valist repo:key exampleOrg exampleRepo grant <key>
  $ valist repo:key exampleOrg exampleRepo revoke <key>
  $ valist repo:key exampleOrg exampleRepo rotate <key>
```

_See code: [dist/commands/repo/key.ts](https://github.com/valist-io/valist/blob/v0.4.1/dist/commands/repo/key.ts)_

## `valist repo:new ORGNAME REPONAME`

Create a Valist repository

```
USAGE
  $ valist repo:new ORGNAME REPONAME

EXAMPLE
  $ valist repo:new exampleOrg exampleRepo
```

_See code: [dist/commands/repo/new.ts](https://github.com/valist-io/valist/blob/v0.4.1/dist/commands/repo/new.ts)_

## `valist repo:update ORGNAME REPONAME REPOMETA`

Update repository metadata

```
USAGE
  $ valist repo:update ORGNAME REPONAME REPOMETA

EXAMPLE
  $ valist repo:update exampleOrg exampleRepo meta/repoMeta.json
```

_See code: [dist/commands/repo/update.ts](https://github.com/valist-io/valist/blob/v0.4.1/dist/commands/repo/update.ts)_
<!-- commandsstop -->
