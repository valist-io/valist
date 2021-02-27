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
    WEB3_PROVIDER: process.env.WEB3_PROVIDER || "http://127.0.0.1:9545",
    MAGIC_PUBKEY: "pk_test_54C6079CBEF87272",
    METATX_ENABLED: process.env.METATX_ENABLED || true,
  },
  webpack: (config, { webpack }) => {
    // Note: we provide webpack above so you should not `require` it
    // Perform customizations to webpack config
    config.plugins.push(new webpack.IgnorePlugin(/^electron$/))

    // Important: return the modified config
    return config
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
