import express from 'express';
import weatherRoute from './routes/weather.route';

export const app = express();

app.use(express.json());
app.use('/weather', weatherRoute);

export default app;