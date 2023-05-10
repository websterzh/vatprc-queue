for i = 1, #ARGV, 2 do
redis.call("ZADD", "_live_check_tmp", ARGV[i], ARGV[i+1])
end
redis.call("ZUNIONSTORE", "_live_check", 2, "_live_check", "_live_check_tmp", "AGGREGATE", "MIN")
redis.call("DEL", "_live_check_tmp")

local members = redis.call("ZRANGE", "_live_check", "0", "-1")
for k, member in pairs(members) do
redis.call("ZINCRBY", "_live_check", 1, member)
end
