let config = {
  async headers() {
    return [
      {
        // matching all API routes
        source: "/api/:path*",
        headers: [
          { key: "Access-Control-Allow-Origin", value: "*" },
          { key: "Cache-Control", value: "s-maxage=60, stale-while-revalidate" },
        ]
      }
    ]
  },
  publicRuntimeConfig: {
    WEB3_PROVIDER: process.env.WEB3_PROVIDER || "https://rpc-mumbai.maticvigil.com",
    MAGIC_PUBKEY: "pk_test_54C6079CBEF87272",
    METATX_ENABLED: process.env.METATX_ENABLED || true,
  },
  webpack5: true,
  webpack: function (config, options) {
    config.plugins.push(new options.webpack.IgnorePlugin({ resourceRegExp: /^electron$/ }));
    return config;
  },
}

if (process.env.IPFS_BUILD == true) {
  config = {
    ...config,
    trailingSlash: true,
    exportPathMap: function() {
      return {
        '/': { page: '/' }
      };
    },
  }
}

module.exports = config;
