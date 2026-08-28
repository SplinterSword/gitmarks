local M = {}

function M.setup(opts)
	local config = opts or {}

	local width = 0
	if config.width ~= nil then
		width = config.width
	end

	vim.notify("Gitmarks started with width " .. width)

	for i = 1, 5 do
		vim.notify(tostring(i))
	end
end

function M.hello()
	vim.notify("Hello from gitmarks")
end

return M
