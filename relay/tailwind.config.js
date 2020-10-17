module.exports = {
  plugins: [
    require('@tailwindcss/ui'),
  ],
  variants: {
    display: ['responsive', 'group-hover', 'group-focus'],
  },
  purge: {
    mode: 'layers',
    layers: ['components', 'utilities'],
    content: ['./**/*.jsx'],
  },
}
