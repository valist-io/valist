import * as fs from 'fs';
import * as path from 'path';
import { createBuild, exportBuild, generateDockerfile } from 'reproducible';
import { ValistConfig } from '@valist/sdk/dist/types';
import { defaultImages } from './config';
import { npmPack } from './npm';

export const buildRelease = async ({
  image,
  install,
  build,
  out,
  type,
}: ValistConfig): Promise<fs.ReadStream> => {
  let releaseFile;

  // Generate Dockerfile (uses current directory as source)
  generateDockerfile(image || defaultImages[type], './', build, install);

  await createBuild('valist-build-image');
  await exportBuild('valist-build-image', out);

  // if package type is npm run npm pack
  if (type === 'node') {
    const packagePath = await npmPack();
    releaseFile = fs.createReadStream(packagePath);
  } else {
    releaseFile = fs.createReadStream(path.join(process.cwd(), out));
  }

  return releaseFile;
};
