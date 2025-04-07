import { Request, Response } from 'express';
import { getWeatherWithCache } from '../services/weather.service';

export const getWeather = async (req: Request, res: Response): Promise<any> => {
  try {
    const city = req.query.city as string;
    if (!city) {
      return res.status(400).json({ message: 'City is required' });
    }

    // Call service here
    const data = await getWeatherWithCache(city);
    return res.json(data);
  } catch (err) {
    return res.status(500).json({ message: 'Something went wrong', error: err });
  }
};
