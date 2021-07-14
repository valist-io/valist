const { withSentryConfig } = require('@sentry/nextjs');

let moduleExports = {
  async headers() {
    return [
      {
        // matching all API routes
        source: '/api/:path*',
        headers: [
          { key: 'Access-Control-Allow-Origin', value: '*' },
          { key: 'Cache-Control', value: 's-maxage=60, stale-while-revalidate' },
        ]
      }
    ]
  },
  publicRuntimeConfig: {
    WEB3_PROVIDER: process.env.WEB3_PROVIDER || 'https://rpc.valist.io',
    MAGIC_PUBKEY: 'pk_test_54C6079CBEF87272',
    METATX_ENABLED: process.env.METATX_ENABLED || true,
  },
  webpack5: true,
  webpack: function (config, options) {
    if (!options.isServer) {
      // polyfill events on browser. since webpack5, polyfills are not automatically included
      config.resolve.fallback.events = require.resolve('events/')
    }
    config.plugins.push(new options.webpack.IgnorePlugin({ resourceRegExp: /^electron$/ }));
    return config;
  },
  // disable sentry source maps upload when building locally
  // sentry: {
  //   disableServerWebpackPlugin: true,
  //   disableClientWebpackPlugin: true,
  // },
}

if (process.env.IPFS_BUILD == true) {
  moduleExports = {
    ...moduleExports,
    trailingSlash: true,
    exportPathMap: function() {
      return {
        '/': { page: '/' }
      };
    },
  }
}

const SentryWebpackPluginOptions = {
  // Additional config options for the Sentry Webpack plugin. Keep in mind that
  // the following options are set automatically, and overriding them is not
  // recommended:
  //   release, url, org, project, authToken, configFile, stripPrefix,
  //   urlPrefix, include, ignore

  silent: true, // Suppresses all logs
  // For all available options, see:
  // https://github.com/getsentry/sentry-webpack-plugin#options.
};

// Make sure adding Sentry options is the last code to run before exporting, to
// ensure that your source maps include changes from all other Webpack plugins
module.exports = withSentryConfig(moduleExports, SentryWebpackPluginOptions);
