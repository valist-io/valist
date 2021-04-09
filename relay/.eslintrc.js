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
    'plugin:react/recommended',
    'plugin:@typescript-eslint/recommended',
    'airbnb-typescript/base',
  ],
  rules: {
    'no-console': 'off',
    'no-alert': 'off',
    'import/extensions': 'off',
    'max-len': ['error', { 'code': 120 }],
    'react/react-in-jsx-scope': 'off',
    '@typescript-eslint/ban-ts-comment': 'off',
    '@typescript-eslint/explicit-module-boundary-types': 'off',
    '@typescript-eslint/no-explicit-any': 'off',
  }
};
