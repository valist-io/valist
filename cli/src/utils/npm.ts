import { exec } from 'child_process';

export const npmPack = async () => {
  try {
    const out: string = await new Promise((resolve) => {
      // npm pack prints to stderr (?), switch params order
      exec('npm pack', (e, stderr, stdout) => {
        resolve(stdout);
      });
    });

    const filename = out.match(/(filename:) +(.+\.tgz)/);
    if (!filename || filename.length < 2) throw new Error('Cannot parse npm package filename');

    return filename[2];
  } catch (e) {
    const msg = 'Could not run npm pack';
    console.error(msg, e);
    throw e;
  }
};
