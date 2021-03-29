const colors = require('tailwindcss/colors');

module.exports = {
  purge: {
    mode: 'layers',
    content: ['./components/**/*.tsx', './pages/**/*.tsx'],
  },

  theme: {
    extend: {
      colors: {
        'light-blue': colors.lightBlue,
        teal: colors.teal,
        cyan: colors.violet,
        rose: colors.rose,
      }
    }
  },
  variants: {
    display: ['responsive', 'group-hover', 'group-focus'],
  },

  plugins: [
    require('@tailwindcss/ui'),
    require('@tailwindcss/forms'),
    require('@tailwindcss/line-clamp'),
  ],
};
