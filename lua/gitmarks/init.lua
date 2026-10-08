local utils = require("gitmarks.utils")
local keymaps = require("gitmarks.keymaps")
local ui = require("gitmarks.ui")
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

	utils.save(project.url, project.file, number, function(_, err)
		if err then
			vim.notify(err, vim.log.levels.ERROR)
			return
		end
		vim.notify("Mark saved to cloud")
	end)

	vim.notify("Marked " .. project.file .. " as " .. number)
	keymaps.createFileBinding(M, number)
end

function M.deleteMark(number)
	local dir_details = utils.get_directory_details()

	if dir_details == nil then
		return
	end

	local remote = dir_details.remote
	local project_marks = M.marks[remote]

	if project_marks == nil then
		return
	end

	project_marks[number] = nil

	utils.delete(remote, number, function(_, err)
		if err then
			vim.notify(err, vim.log.levels.ERROR)
			return
		end
		vim.notify("Mark deleted")
	end)

	vim.notify("Deleted mark " .. number)
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

function M.openViewTab()
	local dir_details = utils.get_directory_details()

	if dir_details == nil then
		vim.notify("Directory Details are not found")
		return
	end

	local remote = dir_details.remote
	local project_marks = M.marks[remote]

	if project_marks == nil then
		vim.notify("No marks found for this project")
		return
	end

	local lines = {}
	local numbers = {}

	for i = 1, 9 do
		local mark = project_marks[i]

		if mark ~= nil then
			table.insert(lines, i .. " - " .. mark)
			table.insert(numbers, i)
		end
	end

	ui.showFiles(lines, numbers, M.deleteMark, keymaps.deleteFileBinding, M.openMarkedFile)
end

function M.setup()
	keymaps.setup(M)

	utils.health(function(_, err)
		if err then
			vim.notify("GitMarks backend down — run `docker start gitmarks`", vim.log.levels.ERROR)
			return
		end

		local dir_details = utils.get_directory_details()

		if dir_details == nil then
			vim.notify("Error in url fetching", vim.log.levels.ERROR)
			return
		end

		local url = dir_details.remote

		utils.get(url, function(marks, err)
			if err then
				if err:find("mongo: no documents in result", 1, true) then
					vim.notify("No marks found for this repo", vim.log.levels.INFO)
					M.marks[url] = {}
					return
				end

				vim.notify(err, vim.log.levels.ERROR)
				return
			end

			if vim.tbl_isempty(marks) then
				vim.notify("No marks found for this repo", vim.log.levels.INFO)
				return
			end

			M.marks[url] = marks

			for number, _ in pairs(M.marks[url]) do
				keymaps.createFileBinding(M, number)
			end

			vim.notify("Marks for the repo retrived successfully")
		end)
	end)
end

return M
