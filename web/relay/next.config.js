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
  generateBuildId: async () => {
    // use a build ID so go embed doesn't get mad
    return 'v0';
  },
}

// disable sentry source maps when not configured properly
if (!(process.env.SENTRY_ORG && process.env.SENTRY_PROJECT && process.env.SENTRY_AUTH_TOKEN)) {
  moduleExports = {
    ...moduleExports,
    sentry: {
      disableServerWebpackPlugin: true,
      disableClientWebpackPlugin: true,
    }
  }
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

// For all available options, see:
// https://github.com/getsentry/sentry-webpack-plugin#options.
const SentryWebpackPluginOptions = {
  silent: true,
};

module.exports = withSentryConfig(moduleExports, SentryWebpackPluginOptions);
