/* eslint-disable max-classes-per-file */

export class InvalidNetworkError extends Error {
  constructor(message: string) {
    super(message);
    this.name = 'InvalidNetworkError';
  }
}

export class ValistSDKError extends Error {
  constructor(message: string) {
    super(message);
    this.name = 'ValistSDKError';
  }
}
