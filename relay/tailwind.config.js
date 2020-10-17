module.exports = {
  purge: {
    mode: 'layers',
    content: ['./components/**/*.tsx', './pages/**/*.tsx'],
  },

  theme: { extend: {} },
  variants: {
    display: ['responsive', 'group-hover', 'group-focus'],
  },

  plugins: [
    require('@tailwindcss/ui'),
  ],
};
