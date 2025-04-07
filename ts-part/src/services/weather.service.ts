import axios from 'axios';
import { getCache, setCache } from './cache.service';

const apiKey = process.env.VISUAL_CROSSING_API_KEY!;
const baseUrl = 'https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline';

export const getWeatherWithCache = async (city: string) => {
  const cached = await getCache(city.toLowerCase());

  if (cached) {
    return JSON.parse(cached);
  }

  const url = `${baseUrl}/${city}?unitGroup=metric&key=${apiKey}&contentType=json`;
  const response = await axios.get(url);
  const data = response.data;

  await setCache(city.toLowerCase(), data, parseInt(process.env.CACHE_TTL || '43200'));
  return data;
};
