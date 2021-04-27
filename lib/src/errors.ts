export class InvalidNetworkError extends Error {
  constructor(message: string) {
    super(message);
    this.name = 'InvalidNetworkError';
  }
}
