local airport = KEYS[1]
local cid = ARGV[1]
local before = ARGV[2]

if before == "-1" then
    local target_score = redis.call("ZSCORE", airport, cid)
    local max_score = math.floor(target_score)+1
    local last_element = redis.call("ZRANGE", airport, "("..max_score, target_score, "BYSCORE", "WITHSCORES", "REV", "LIMIT", 0, 1)
    if #last_element == 2 then
        redis.call("ZADD", airport, tonumber(last_element[2])+0.000001, cid)
    end
    return
end

local before_element = redis.call("ZSCORE", airport, before)
if before_element == nil then
    return
end

local start_score = tonumber(before_element)
local end_score = math.floor(start_score)+1
local target_score = start_score

local to_change = redis.call("ZRANGE", airport, start_score, "("..end_score, "BYSCORE")
local num_arg = #to_change
for i = 1, num_arg do
    target_score = target_score + 0.000001
    redis.call("ZADD", airport, target_score, to_change[i])
end

redis.call("ZADD", airport, start_score, cid)
