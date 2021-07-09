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
valist-cli/0.0.4 darwin-x64 node-v14.17.2
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
* [`valist help [COMMAND]`](#valist-help-command)
* [`valist init`](#valist-init)
* [`valist org:get ORGNAME`](#valist-orgget-orgname)
* [`valist org:key ORGNAME OPERATION KEY`](#valist-orgkey-orgname-operation-key)
* [`valist org:new ORGNAME ORGMETA`](#valist-orgnew-orgname-orgmeta)
* [`valist org:update ORGNAME ORGMETA`](#valist-orgupdate-orgname-orgmeta)
* [`valist publish`](#valist-publish)
* [`valist repo:get ORGNAME REPONAME`](#valist-repoget-orgname-reponame)
* [`valist repo:key ORGNAME REPONAME OPERATION KEY`](#valist-repokey-orgname-reponame-operation-key)
* [`valist repo:new ORGNAME REPONAME REPOMETA`](#valist-reponew-orgname-reponame-repometa)
* [`valist repo:update ORGNAME REPONAME REPOMETA`](#valist-repoupdate-orgname-reponame-repometa)

## `valist account:get`

print account info

```
USAGE
  $ valist account:get

EXAMPLE
  $ valist account:get
```

## `valist account:new`

create a new account

```
USAGE
  $ valist account:new

EXAMPLE
  $ valist account:new
```

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

## `valist org:get ORGNAME`

print organization info

```
USAGE
  $ valist org:get ORGNAME

EXAMPLE
  $ valist org:get valist
```

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

## `valist org:new ORGNAME ORGMETA`

Create a Valist organization

```
USAGE
  $ valist org:new ORGNAME ORGMETA

EXAMPLE
  $ valist org:new valist meta/metOrg.json
```

## `valist org:update ORGNAME ORGMETA`

Update organization metadata

```
USAGE
  $ valist org:update ORGNAME ORGMETA

EXAMPLE
  $ valist org:update exampleOrg meta/orgMeta.json
```

## `valist publish`

Publish a package to Valist

```
USAGE
  $ valist publish

EXAMPLE
  $ valist publish
```

## `valist repo:get ORGNAME REPONAME`

print organization info

```
USAGE
  $ valist repo:get ORGNAME REPONAME

EXAMPLE
  $ valist repo:get exampleOrg exampleRepo
```

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

## `valist repo:new ORGNAME REPONAME REPOMETA`

Create a Valist repository

```
USAGE
  $ valist repo:new ORGNAME REPONAME REPOMETA

EXAMPLE
  $ valist repo:new exampleOrg exampleRepo meta/repoMeta.json
```

## `valist repo:update ORGNAME REPONAME REPOMETA`

Update repository metadata

```
USAGE
  $ valist repo:update ORGNAME REPONAME REPOMETA

EXAMPLE
  $ valist repo:update exampleOrg exampleRepo meta/repoMeta.json
```
<!-- commandsstop -->
