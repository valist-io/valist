module.exports = {
  publicRuntimeConfig: {
    WEB3_PROVIDER: process.env.WEB3_PROVIDER || "http://127.0.0.1:9545",
    MAGIC_PUBKEY: "pk_test_54C6079CBEF87272",
    METATX_ENABLED: process.env.METATX_ENABLED || true,
  },
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
  trailingSlash: true,
  exportPathMap: function() {
    return {
      '/': { page: '/' }
    };
  }
}
