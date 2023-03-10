-- 1. 检查是不是你的锁
-- 2. 删除
-- KEYS[1] 是调用lua脚本时传入的key值，是分布式锁的key
-- ARGV[1] 是调用lua脚本时传入的value值，是预期在redis 里存的 value
if redis.call('get', KEYS[1]) == ARGV[1] then
    -- 检查是自己的锁，则删除
    return redis.call('del', KEYS[1])
else
    -- 不是自己的锁，不操作
    return 0
end