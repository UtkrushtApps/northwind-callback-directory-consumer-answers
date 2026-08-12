function loadConfig() {
  return {
    upstreamBaseUrl:
      process.env.NORTHWIND_UPSTREAM_BASE_URL || 'https://jsonmock.hackerrank.com/api/countries',
    requestTimeoutMs: Number(process.env.NORTHWIND_REQUEST_TIMEOUT_MS || 5000),
  };
}

module.exports = { loadConfig };
