module.exports = {
  target: 'serverless',
  webpack: (config, { webpack }) => {
    // Note: we provide webpack above so you should not `require` it
    // Perform customizations to webpack config
    config.plugins.push(new webpack.IgnorePlugin(/^electron$/))

    // Important: return the modified config
    return config
  },
}
