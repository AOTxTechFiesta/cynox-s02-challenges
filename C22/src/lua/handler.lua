local cjson = require("cjson.safe")
local upload = require("resty.upload")
local str = require("resty.string")

-- Initialize shared dictionary
local shared_config = ngx.shared.config_data

local default_config = {
    usermessages = {{
        name = "Vilgax",
        message = "I am leaving the flag in /flag.txt."
    }, {
        name = "itsyourap",
        message = "Alright Sir! I will do my best to protect the flag!"
    }}
}

-- Generate a unique token for users using a simple time-based approach
local function generate_token()
    -- Use time and a "random" number to create a unique string
    local time = ngx.time()
    local rand = math.random(1000000, 9999999)
    local token = time .. "_" .. rand
    
    -- Make it a bit harder to predict by hashing with a simple algorithm
    local hash = ngx.md5(token)
    return hash
end

-- Make sure to seed the random number generator
math.randomseed(ngx.time())

-- Helper function to get/set the shared config
local function get_user_token()
    -- Get the user token from cookie or return default
    local token = ngx.var.cookie_user_token
    if not token then
        token = "default"
    end
    return token
end

local function set_config(config, token)
    token = token or get_user_token()
    shared_config:set(token, cjson.encode(config))
end

local function get_config()
    local token = get_user_token()
    
    -- If user has a token but no config, use default
    local config_json = shared_config:get(token)
    if not config_json then
        if token ~= "default" then
            -- Set default config for this user
            set_config(default_config, token)
        end
        return default_config
    end
    
    return cjson.decode(config_json)
end

local function handle_get_data()
    local config = get_config()
    ngx.header.content_type = "application/json"
    ngx.say(cjson.encode({
        usermessages = config.usermessages
    }))
end

local function handle_export()
    local config = get_config()
    ngx.header["Content-Type"] = "text/plain"
    ngx.header["Content-Disposition"] = "attachment; filename=config.conf"

    ngx.say("config.usermessages = {}")
    for i, u in ipairs(config.usermessages) do
        ngx.say(string.format('config.usermessages[%d] = {}', i))
        ngx.say(string.format('config.usermessages[%d]["name"] = "%s"', i, u.name))
        ngx.say(string.format('config.usermessages[%d]["message"] = "%s"', i, u.message))
    end
end

local function handle_import()
    -- Check for user token or create one
    local token = ngx.var.cookie_user_token
    if not token then
        token = generate_token()
        -- Set cookie with 24 hour expiration
        ngx.header["Set-Cookie"] = "user_token=" .. token .. "; Path=/; Max-Age=86400; HttpOnly"
    end
    
    local chunk_size = 1024
    local form, err = upload:new(chunk_size)
    if not form then
        ngx.say("Failed to initialize upload: ", err)
        return
    end

    local file_content = ""
    local in_file = false

    while true do
        local typ, res, err = form:read()
        if not typ then
            ngx.say("Failed to read: ", err)
            return
        end

        if typ == "header" then
            local header_name = res[1]
            local header_value = res[2]
            if header_name == "Content-Disposition" and header_value:find("filename=") then
                in_file = true
            end
        elseif typ == "body" and in_file then
            file_content = file_content .. res
        elseif typ == "part_end" then
            in_file = false
        elseif typ == "eof" then
            break
        end
    end

    if file_content == "" then
        ngx.say("No config uploaded or empty file")
        return
    end

    local tmp_file = "/tmp/uploaded_config_" .. token .. ".lua"
    local f = io.open(tmp_file, "w")
    f:write(file_content)
    f:close()

    for line in io.lines(tmp_file) do
        if not line:match("^config") then
            ngx.say("Config file is corrupted! Invalid line: ", line)
            return
        end
    end

    -- Get current config
    local config = get_config()

    -- Inject config into global environment
    _G.config = config

    -- Execute the uploaded config file
    local ok, err = pcall(function()
        dofile(tmp_file)
    end)

    -- Update the shared config from global
    if ok then
        set_config(_G.config, token)
        ngx.say("Configuration updated!")
    else
        ngx.say("Error executing config file: ", err)
        return
    end
end

local function serve_html()
    ngx.exec("@static_html")
end

-- Router
local uri = ngx.var.uri
if uri == "/get-data" then
    handle_get_data()
elseif uri == "/export" then
    handle_export()
elseif uri == "/import" then
    handle_import()
else
    serve_html()
end
