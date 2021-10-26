# Multi-Factor Everything

One of the coolest features of Valist is that there is a multi-factor system for everything. With multi-factor everything you can eliminate single points of failure within your software supply chain and have all actions audited and approved by members across your organization.

## How it works

You may have heard of the [two man rule](https://en.wikipedia.org/wiki/Two-man_rule) -- an authentication mechanism designed to prevent accidental or malicious launch of nuclear weapons by a single individual. This mechanism requires the presence of two authorized individuals, **with special keys**, to perform a nuclear strike. Valist applies this same concept to software distribution. When enabled multi-party verification requires operations be verified and digitally signed by multiple members within your organization before being finalized.

![two-man-rule](./img/two-man-rule.jpeg)

## Organization vs Repo Level

On the **organization** level, organization admins are able to vote on:

* Adding new organization admins to the organization.

On the **repository** level, members are able to vote on:

* Approving a new repository threshold for votes.

* Adding a new developer key to the repository.

* Revoking an existing developer key from the repository.

## Setting up a Signature Threshold

By default each organization's and repository's threshold is set to **0**.

### CLI

#### Organization Level Threshold

Using the CLI an `Org Admin`  is able to propose a new `Organization` threshold by running:

```bash
valist org threshold [org_name] [threshold_number]
```

#### Repository Level Threshold

Using the CLI a `Repo Developer` is able to propose a new `Repository` threshold by running:

```bash
valist repo threshold [org_name] [repo_name] [threshold_number]
```

## Voting on Access Control & Releases

After a repository or organization's threshold it set to a number greater than **1**, all operations will require the target number of votes before they are finalized.

## Considerations

* By design the repository threshold can only be set to a maximum of `n-1`. This it to create a permanent buffer of 1 key so that by default if members misplace a key they do not loose access to their organization or repository.

* When revoking a key from an organization or repository if the number of remaining keys is equal to the current threshold then the threshold will be automatically decreased by one.

* A repository or organization must have greater than `2` members to enable multi-factor voting.
