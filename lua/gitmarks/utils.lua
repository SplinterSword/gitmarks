local M = {}

function M.get_git_root()
	local cwd = vim.fn.getcwd()

	local git_dir = vim.fs.find(".git", {
		path = cwd,
		upward = true,
		type = "directory",
	})[1]

	if git_dir == nil then
		return nil
	end

	return vim.fs.dirname(git_dir)
end

function M.get_remote(root)
	local result = vim.system({
		"git",
		"-C",
		root,
		"remote",
		"get-url",
		"origin",
	}, {
		text = true,
	}):wait()

	if result.code ~= 0 then
		return nil
	end

	return vim.trim(result.stdout)
end

function M.normalize_remote(remote)
	remote = remote:gsub("%.git$", "")

	remote = remote:gsub("^git@github.com:", "https://github.com/")
	remote = remote:gsub("^ssh://git@github.com/", "https://github.com/")

	return remote
end

function M.get_file_details(file)
	if file == "" then
		return nil
	end

	local root = M.get_git_root()

	if root == nil then
		return nil
	end

	local remote = M.get_remote(root)

	if remote == nil then
		return nil
	end

	remote = M.normalize_remote(remote)

	local relative_path = vim.fs.relpath(root, file)

	return {
		url = remote,
		root = root,
		file = relative_path,
	}
end

function M.get_directory_details()
	local root = M.get_git_root()

	if root == nil then
		return nil
	end

	local remote = M.get_remote(root)

	if remote == nil then
		return nil
	end

	remote = M.normalize_remote(remote)

	return {
		root = root,
		remote = remote,
	}
end

return M
