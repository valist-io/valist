export type OrgID = string;

export type OrgName = string;

export type OrgMeta = {
  name: string,
  description: string
  homepage?: string
};

export type BinaryTypes = 'binary' | 'go' | 'rust' | 'c++';

export type ProjectType = BinaryTypes | 'node' | 'npm' | 'python' | 'docker' | 'static';

export type RepoMeta = {
  name: string,
  description: string,
  projectType: ProjectType,
  homepage?: string,
  repository?: string
};

export type Organization = {
  // organization ID
  orgID: OrgID,
  // multi-party threshold
  threshold: number,
  // date threshold was set
  thresholdDate: number,
  // parsed JSON from metaCID
  meta: OrgMeta,
  // metadata CID
  metaCID: string,
  // list of repo names
  repoNames: string[]
};

export type Repository = {
  // organization ID
  orgID: OrgID,
  // multi-party threshold
  threshold: number,
  // date threshold was set
  thresholdDate: number,
  // parsed JSON from metaCID
  meta: RepoMeta,
  // metadata CID
  metaCID: string,
  // list of release tags
  tags: string[]
};

export type Release = {
  // release tag/version
  tag: string,
  // release artifact
  releaseCID: string,
  // release metadata
  metaCID: string,
  // finalized release signers
  signers?: string[],
};

export type PendingVote = {
  // expiration date of vote
  expiration: string,
  // signers that have voted on this release
  signers: string[]
};

export type PendingRelease = {
  // proposed tag
  tag: string,
  // release artifact
  releaseCID: string,
  // release metadata
  metaCID: string,
  // pending vote
  pendingVote?: PendingVote
};

export type VoteThresholdEvent = {
  _orgID: string,
  _repoName: string,
  _signer: string,
  _pendingThreshold: string,
  _sigCount: string,
  _threshold: string
};

export type VoteKeyEvent = {
  _orgID: string,
  _repoName: string,
  _signer: string,
  _operation: string,
  _key: string,
  _sigCount: string,
  _threshold: string
};

export type VoteReleaseEvent = {
  _orgID: string,
  _repoName: string,
  _tag: string,
  _releaseCID: string,
  _metaCID: string,
  _signer: string,
  _sigCount: string,
  _threshold: string
};

export type ValistCache = {
  // orgIDs: OrgID[],
  // orgNames: OrgName[],
  // orgs: Record<OrgName, Organization>,
  // repos: Record<OrgID, Record<string, Repository>>,
  orgIDs: Record<OrgName, OrgID>,
};

export type ValistConfig = {
  type: ProjectType,
  org: string,
  repo: string,
  tag: string,
  meta: string,
  image: string,
  build: string,
  install: string,
  out?: string,
  artifacts?: Record<string, string>,
};
