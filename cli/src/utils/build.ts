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
  let outPath = out;
  let releaseFile;

  // Generate Dockerfile (uses current directory as source)
  generateDockerfile(image || defaultImages[type], './', build, install);

  if (type !== 'node') {
    // if out path is a file, cut file from mount path/get parent directory
    outPath = path.basename(path.dirname(out));
  }

  await createBuild('valist-build-image');
  await exportBuild('valist-build-image', outPath);

  // if package type is npm run npm pack
  if (type === 'node') {
    const packagePath = await npmPack();
    releaseFile = fs.createReadStream(packagePath);
  } else {
    releaseFile = fs.createReadStream(path.join(process.cwd(), out));
  }

  return releaseFile;
};
