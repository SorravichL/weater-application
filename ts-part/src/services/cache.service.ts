import Redis from "ioredis";

const redis = new Redis(process.env.REDIS_URL || 'redis://localhost:6379');

export const getCache = async (key:string) : Promise<string | null> => {
    return await redis.get(key);
}

export const setCache = async (key: string, value: any, ttl = 43200): Promise<void> => {
    await redis.set(key, JSON.stringify(value), 'EX', ttl);
}