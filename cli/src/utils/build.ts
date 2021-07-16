import * as fs from 'fs';
import * as path from 'path';
import { createBuild, exportBuild, generateDockerfile } from 'reproducible';
import { ValistConfig } from '@valist/sdk/dist/types';
import { defaultImages } from './config';
import { npmPack } from './npm';

export const buildRelease = async (config : ValistConfig): Promise<fs.ReadStream[]> => {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const releaseFiles: fs.ReadStream[] = [];

  // Generate Dockerfile (uses current directory as source)
  generateDockerfile(config.image || defaultImages[config.type], './', config.build, config.install);

  await createBuild('valist-build-image');
  await exportBuild('valist-build-image', config.out);

  // if package type is npm run npm pack
  if (config.type === 'node') {
    const packagePath = await npmPack();
    releaseFiles.push(fs.createReadStream(packagePath));
  } else {
    config.artifacts.forEach((artifact) => {
      releaseFiles.push(fs.createReadStream(path.join(process.cwd(), '/dist/', artifact)));
    });
  }
  return releaseFiles;
};
