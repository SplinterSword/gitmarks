local M = {}

local BASE_URL = "http://localhost:34179"

function M.request(method, endpoint, body, callback)
	local args = {
		"curl",
		"-s",
		"-X",
		method,
		BASE_URL .. endpoint,
		"-H",
		"Content-Type: application/json",
	}

	if body then
		table.insert(args, "-d")
		table.insert(args, vim.json.encode(body))
	end

	vim.system(args, {
		text = true,
	}, function(result)
		vim.schedule(function()
			if result.code ~= 0 then
				callback(nil, result.stderr)
				return
			end

			local ok, data = pcall(vim.json.decode, result.stdout)

			if not ok then
				callback(nil, "Invalid JSON response: " .. result.stdout)
				return
			end

			callback(data, nil)
		end)
	end)
end

return M
