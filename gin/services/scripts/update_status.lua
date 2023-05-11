local airport = KEYS[1]
local callsign = ARGV[1]
local target_status = tonumber(ARGV[2])
local end_score = target_status+1
local next_score = target_status

local last_element = redis.call("ZRANGE", airport, "("..end_score, target_status, "BYSCORE", "WITHSCORES", "REV", "LIMIT", 0, 1)
if #last_element == 2 then
    next_score = tonumber(last_element[2]) + 0.000001
end

redis.call("ZADD", airport, next_score, callsign)