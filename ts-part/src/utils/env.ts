import dotenv from 'dotenv';
dotenv.config();

const required = ['VISUAL_CROSSING_API_KEY', 'REDIS_URL'];

for (const key of required) {
  if (!process.env[key]) throw new Error(`Missing ENV variable: ${key}`);
}
