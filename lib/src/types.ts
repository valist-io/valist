export type ProjectType = 'binary' | 'npm' | 'pip' | 'docker';

export type OrgMeta = {
  name: string,
  description: string
};

export type RepoMeta = {
  name: string,
  description: string,
  projectType: ProjectType,
  homepage: string,
  repository: string
};

export type Release = {
  releaseCID: string,
  metaCID: string,
  tag: string
};
