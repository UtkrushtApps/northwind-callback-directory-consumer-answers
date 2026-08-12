function createDirectoryService(client) {
  return {
    async getPhoneNumbers(country, phone) {
      const payload = await client.fetchCountry(country);
      const records = Array.isArray(payload && payload.data) ? payload.data : [];
      if (records.length === 0) {
        return '-1';
      }

      const callingCodes = Array.isArray(records[0].callingCodes) ? records[0].callingCodes : [];
      if (callingCodes.length === 0) {
        return '-1';
      }

      const callingCode = callingCodes[callingCodes.length - 1];
      return '+' + callingCode + ' ' + phone;
    },
  };
}

module.exports = { createDirectoryService };
