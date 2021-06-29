import * as fs from 'fs';
import { createImage, runBuild } from 'reproducible';
import { ValistConfig } from './config';

const sleep = (time: any) => new Promise((resolve) => setTimeout(resolve, time));

export const generateDockerfile = (baseImage: string,
  source: string, projectType: string, buildCommand?: string, installCommand?: string) => {
  let dockerfile = `FROM ${baseImage}
WORKDIR /opt/build/${source}
COPY ./${source} ./`;

  if (installCommand) {
    dockerfile += `\nRUN ${installCommand}`;
  }

  if (buildCommand) {
    dockerfile += `\nRUN ${buildCommand}`;
  }

  if (projectType === 'npm') {
    dockerfile += '\nRUN npm pack';
  }

  return dockerfile;
};

export const buildRelease = async ({
  image, build, type,
}: ValistConfig) => {
  const dockerfile = generateDockerfile(image, 'src', type, build);

  fs.writeFile('Dockerfile', dockerfile, async (err: any) => {
    if (err) throw err;
  });

  await createImage('build-image');
  await runBuild({
    image: 'build-image',
    outputPath: `${process.cwd()}/dist`,
    artifacts: ['main'],
  });

  // hack-y fix (needs refactor)
  await sleep(5000);

  const releaseFile = fs.createReadStream(`${process.cwd()}/dist/main`);

  return releaseFile;
};
