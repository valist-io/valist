export type Action = {
  description: string,
  command: string,
};

export type ProjectType = {
  actions: string[],
  default: string
};

export const GetActions = (location:string, orgName:string, repoName:string) => {
  const actions: Record<string, Action> = {
    npmInstallFromRegistry: {
      description: 'Install from Registry',
      command: `# Link Namespace to Registry
echo @${orgName}:registry=${location}/api/npm >> .npmrc

# Install from Registry
npm i @${orgName}/${repoName}`,
    },
    curlBinary: {
      description: 'Download (GET) from Url',
      command: `curl -L -o ${repoName} ${location}/api/${orgName}/${repoName}/latest

      `,
    },
    pipInstall: {
      description: 'Pip Install From Url',
      command: `pip install ${location}/api/${orgName}/${repoName}/latest

      `,
    },
    dockerLoad: {
      description: 'Load Container from Url',
      command: `curl -L ${location}/api/${orgName}/${repoName}/latest | docker load

      `,
    },
  };
  return actions;
};

export const projectTypes: Record<string, ProjectType> = {
  binary: {
    actions: ['curlBinary'],
    default: 'curlBinary',
  },
  npm: {
    actions: ['installUrl', 'npmInstallFromRegistry'],
    default: 'npmInstallFromRegistry',
  },
  node: {
    actions: ['installUrl', 'npmInstallFromRegistry'],
    default: 'npmInstallFromRegistry',
  },
  go: {
    actions: ['curlBinary'],
    default: 'curlBinary',
  },
  python: {
    actions: ['pipInstall'],
    default: 'pipInstall',
  },
  docker: {
    actions: ['dockerLoad'],
    default: 'dockerLoad',
  },
};
