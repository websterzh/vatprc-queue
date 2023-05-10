local threshold = ARGV[1]
local to_clean = redis.call("ZRANGEBYSCORE", "_live_check", ""..threshold, "+inf")
redis.call("ZREMRANGEBYSCORE", "_live_check", ""..threshold, "+inf")
return to_clean