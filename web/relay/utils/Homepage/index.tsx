export type Integration = {
  name: string,
  image: string,
  actions: string[],
  docs: string,
  code: string
};

export type Action = {
  description: JSX.Element,
  code: string
};

export const GetActions = (location:string) => {
  const actions: Record<string, Action> = {
    npmInstall: {
      description: (<div>Npm packages can be <b>installed</b> by using the registry flag or by url.</div>),
      code: `# Install a package via registry
npm i --registry=${location}api/npm @valist/sdk
      
# Install a package via IPFS gateway
npm i ${location}api/npm/valist/sdk`,
    },
    npmPublish: {
      description: (<div>
        Npm packages can be <b>published</b> by using the registry flag or by setting your NPM registry.
      </div>),
      code: `# Publish a package to registry
npm publish --registry=${location}api/npm`,
    },
    dockerPush: {
      description: (<div>Docker images can be <b>tagged</b> and <b>pushed</b> with the docker pull command.</div>),
      code: `# Tag the Image
docker image tag [image_name]:latest ${location}api/oci/[org][repo]:latest
  
# Push the Image
docker image push ${location}api/oci/[org][repo]:latest`,
    },
    dockerPull: {
      description: (<div> Docker images can be <b>pulled</b> using docker push command.</div>),
      code: `# Pull container image
docker pull ${location}api/oci/[org][repo]:latest`,
    },
    gitPush: {
      description: (<div> Project source code can be committed and <b>pushed</b> using the git push command.</div>),
      code: `# Push to remote
git push ${location}api/git/[org][repo]`,
    },
    gitPull: {
      description: (<div> Project source can be <b>pulled</b> using the git push command.</div>),
      code: `# Pull from remote
git pull ${location}api/git/[org][repo]`,
    },
    goGet: {
      description: (<div> Go Modules can be <b>installed</b> with the go get command.</div>),
      code: `# Installs package from remote
go get ${location}api/git/[org][repo]`,
    },
  };

  return actions;
};

export const integrations: Integration[] = [
  {
    name: 'NodeJS Packages',
    image: '/images/npm-logo.png',
    actions: ['npmPublish', 'npmInstall'],
    docs: 'https://docs.valist.io/fetching-and-installing-software/#npm-registry',
    code: 'https://github.com/valist-io/example-projects',
  },
  {
    name: 'Docker',
    image: '/images/docker-logo.png',
    actions: ['dockerPush', 'dockerPull'],
    docs: 'https://docs.valist.io/fetching-and-installing-software',
    code: 'https://github.com/valist-io/example-projects',
  },
  {
    name: 'Source Control',
    image: '/images/git-logo.png',
    actions: ['gitPush', 'gitPull'],
    docs: 'https://docs.valist.io/fetching-and-installing-software/',
    code: 'https://github.com/valist-io/example-projects',
  },
  {
    name: 'Go Packages',
    image: '/images/go-logo.png',
    actions: ['goGet'],
    docs: 'https://docs.valist.io/fetching-and-installing-software/',
    code: 'https://github.com/valist-io/example-projects',
  },
];

export const links = [
  {
    name: 'RPC Gateway',
    href: 'https://rpc.valist.io',
  },
  {
    name: 'IPFS Gateway',
    href: 'https://gateway.valist.io/',
  },
  {
    name: 'Valist Github',
    href: 'https://github.com/valist-io',
  },
];
