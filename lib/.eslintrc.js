module.exports = {
  root: true,
  parser: '@typescript-eslint/parser',
  parserOptions: {
    project: './tsconfig.json',
  },
  plugins: [
    '@typescript-eslint/eslint-plugin',
  ],
  extends: [
    'plugin:@typescript-eslint/recommended',
    'airbnb-typescript/base'
  ],
  rules: {
    'no-console': 'off',
    'no-plusplus': 'off',
    'import/prefer-default-export': 'off',
    'max-len': ['error', { 'code': 120 }],
  }
};
