const request = require('supertest');
const { createApp } = require('../src/app');
const { createDirectoryService } = require('../src/service');

function stubClient(payload) {
  return { fetchCountry: async () => payload };
}

function buildApp(payload) {
  const service = createDirectoryService(stubClient(payload));
  return createApp(service);
}

test('endpoint returns JSON result for known country', async () => {
  const app = buildApp({ data: [{ name: 'Afghanistan', callingCodes: ['93'] }] });
  const response = await request(app)
    .get('/phone-numbers')
    .query({ country: 'Afghanistan', phone: '656445445' });
  expect(response.status).toBe(200);
  expect(response.body).toEqual({ result: '+93 656445445' });
});

test('endpoint returns JSON result for multi-code country', async () => {
  const app = buildApp({ data: [{ name: 'Puerto Rico', callingCodes: ['1', '1787', '1939'] }] });
  const response = await request(app)
    .get('/phone-numbers')
    .query({ country: 'Puerto Rico', phone: '123456789' });
  expect(response.status).toBe(200);
  expect(response.body).toEqual({ result: '+1939 123456789' });
});

test('endpoint returns successful -1 JSON result for unknown country', async () => {
  const app = buildApp({ data: [] });
  const response = await request(app)
    .get('/phone-numbers')
    .query({ country: 'Atlantis', phone: '5551234' });
  expect(response.status).toBe(200);
  expect(response.body).toEqual({ result: '-1' });
});

test('endpoint rejects missing country', async () => {
  const app = buildApp({ data: [] });
  const response = await request(app).get('/phone-numbers').query({ phone: '5551234' });
  expect(response.status).toBe(400);
  expect(response.body).toEqual({ error: 'country and phone are required' });
});

test('endpoint rejects blank phone', async () => {
  const app = buildApp({ data: [] });
  const response = await request(app)
    .get('/phone-numbers')
    .query({ country: 'Afghanistan', phone: '   ' });
  expect(response.status).toBe(400);
  expect(response.body).toEqual({ error: 'country and phone are required' });
});
