local utils = require("gitmarks.utils")
local keymaps = require("gitmarks.keymaps")
local M = {}

M.marks = {}

function M.markFile(number)
	local file = vim.api.nvim_buf_get_name(0)

	if file == nil then
		vim.notify("File not found in the buffer")
	end

	local project = utils.get_file_details(file)

	if project == nil then
		vim.notify("Not inside a GitHub repository")
		return
	end

	if M.marks[project.url] == nil then
		M.marks[project.url] = {}
	end

	M.marks[project.url][number] = project.file

	vim.notify("Marked " .. project.file .. " as " .. number)
	print(vim.inspect(M.marks))
	keymaps.createFileBinding(M, number)
end

function M.openMarkedFile(number)
	local dir_details = utils.get_directory_details()

	if dir_details == nil then
		vim.notify("Directory details not found")
		return
	end

	local root = dir_details.root
	local remote = dir_details.remote

	local project_marks = M.marks[remote]

	if project_marks == nil then
		vim.notify("No marks found for this project")
		return
	end

	local file = project_marks[number]

	if file == nil then
		vim.notify("No file marked as " .. number)
		return
	end

	local path = vim.fs.joinpath(root, file)

	vim.cmd.edit(vim.fn.fnameescape(path))
end

function M.setup()
	keymaps.setup(M)
end

return M
