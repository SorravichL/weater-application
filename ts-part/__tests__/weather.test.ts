import request from 'supertest';
import app from '../src/app';

describe('GET /weather', () => {
  it('should return 400 if no city is provided', async () => {
    const res = await request(app).get('/weather');
    expect(res.statusCode).toBe(400);
    expect(res.body.message).toBe('City is required');
  });

  it('should return weather data for a valid city', async () => {
    const res = await request(app).get('/weather?city=Bangkok');
    expect(res.statusCode).toBe(200);
    expect(res.body).toHaveProperty('resolvedAddress');
  }, 10000); // Set timeout in case API call is slow
});
