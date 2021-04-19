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
    'airbnb-typescript/base'
  ],
  rules: {
    'no-console': 'off',
    'import/prefer-default-export': 'off',
  }
};
