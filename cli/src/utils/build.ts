import * as fs from 'fs';
import * as path from 'path';
import { createBuild, exportBuild, generateDockerfile } from 'reproducible';
import { ValistConfig } from '@valist/sdk/dist/types';
import { defaultImages, parsePackageJson } from './config';

export const buildRelease = async ({
  image,
  install,
  build,
  out,
  type,
}: ValistConfig): Promise<fs.ReadStream> => {
  let buildCommand = build;
  let outArtifact = out;

  if (type === 'node') {
    const packageJson = parsePackageJson();
    buildCommand = `${build} && npm pack`;
    outArtifact = `${packageJson.name}-${packageJson.version}.tgz`;
  }

  // Generate Dockerfile (uses current directory as source)
  generateDockerfile(image || defaultImages[type], './', buildCommand, install);

  await createBuild('valist-build-image');
  await exportBuild('valist-build-image', outArtifact);

  // if package type is npm run npm pack
  const releaseFile = fs.createReadStream(path.join(process.cwd(), outArtifact));

  return releaseFile;
};
