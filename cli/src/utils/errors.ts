export class MissingKeyError extends Error {
  constructor() {
    super("🔎 Key not found! Use 'valist account:new' to create a new key.");
    this.name = 'MissingKeyError';
  }
}