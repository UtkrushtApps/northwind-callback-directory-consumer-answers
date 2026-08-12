const axios = require('axios');

function createCountriesClient(config) {
  return {
    async fetchCountry(country) {
      const response = await axios.get(config.upstreamBaseUrl, {
        params: { name: country },
        timeout: config.requestTimeoutMs,
      });
      return response.data;
    },
  };
}

module.exports = { createCountriesClient };
