module.exports = {
  purge: {
    mode: 'layers',
    content: ['./components/**/*.{js,ts,jsx,tsx}', './pages/**/*.{js,ts,jsx,tsx}'],
  },

  theme: { extend: {} },
  variants: {
    display: ['responsive', 'group-hover', 'group-focus'],
  },

  plugins: [
    require('@tailwindcss/ui'),
  ],
};
