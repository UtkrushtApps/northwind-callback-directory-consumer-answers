const express = require('express');

function isBlank(value) {
  return typeof value !== 'string' || value.trim() === '';
}

function createRouter(service) {
  const router = express.Router();

  router.get('/phone-numbers', async (req, res) => {
    const { country, phone } = req.query;

    if (isBlank(country) || isBlank(phone)) {
      return res.status(400).json({ error: 'country and phone are required' });
    }

    try {
      const resolved = await service.getPhoneNumbers(country, phone);
      return res.json({ result: resolved });
    } catch (error) {
      return res.status(500).json({ error: error.message || 'internal server error' });
    }
  });

  return router;
}

module.exports = { createRouter };
