# CLI Commands

valist

```
[--account]=[value]
```

**Usage**:

```
valist [GLOBAL OPTIONS] command [COMMAND OPTIONS] [ARGUMENTS...]
```

# GLOBAL OPTIONS

**--account**="": Account to transact with


# COMMANDS

## account

Create, update, or list accounts

### create

Create an account

### list

List all accounts

### default

Set the default account

### pin

Pin an account from an external wallets

### unpin

Unpin an account from an external wallet

## organization, org

Create, update, or fetch organizations

### fetch, get

Fetch organization info

### create

Create an organization

### update

Update organization metadata

### threshold

Vote for organization threshold

### key

Manage keys at an organization level

#### add

Add a new key to an organization

#### revoke

Remove a key from an organization

#### rotate

Replace a key on an organization

## repository, repo

Create, update, or fetch repositories

### fetch, get

Fetch repository info

### create

Create a repository

### update

Update repository metadata

### threshold

Vote for repository threshold

### key

Manage keys at a repository level

#### add

Add a new key to a repository

#### revoke

Remove a key from a repository

#### rotate

Replace a key on a repository

## daemon

Runs a relay node

**--account**="": Account to authenticate with (default: default)

## build

Build the target valist project

## init

Generate a new Valist project

**--wizard, -i**: Enable interactive wizard

## publish

Publish a package

**--dryrun**: Build and skip publish
