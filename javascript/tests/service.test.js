const { createDirectoryService } = require('../src/service');

function stubClient(payload) {
  return { fetchCountry: async () => payload };
}

test('single calling code resolves', async () => {
  const service = createDirectoryService(
    stubClient({ data: [{ name: 'Afghanistan', callingCodes: ['93'] }] })
  );
  const result = await service.getPhoneNumbers('Afghanistan', '656445445');
  expect(result).toBe('+93 656445445');
});

test('multiple calling codes resolve using highest-index code', async () => {
  const service = createDirectoryService(
    stubClient({ data: [{ name: 'Puerto Rico', callingCodes: ['1', '1787', '1939'] }] })
  );
  const result = await service.getPhoneNumbers('Puerto Rico', '123456789');
  expect(result).toBe('+1939 123456789');
});

test('unknown country resolves successfully as -1', async () => {
  const service = createDirectoryService(stubClient({ data: [] }));
  const result = await service.getPhoneNumbers('Atlantis', '5551234');
  expect(result).toBe('-1');
});
