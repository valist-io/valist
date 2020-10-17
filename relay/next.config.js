module.exports = {
  target: 'serverless',
  webpack: (config, { webpack }) => {
    config.plugins.push(new webpack.IgnorePlugin(/^electron$/));
    /*config.module.rules.push({
      use: [
        {
          loader: 'postcss-loader',
          options: {
            postcssOptions: {
              ident: 'postcss',
              plugins: [
                require('tailwindcss'),
                require('autoprefixer'),
              ],
            },
          }
        },
      ],
    });*/
    return config
  },
}
