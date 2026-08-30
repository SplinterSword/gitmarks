local M = {}

function M.showFiles(lines, numbers, delete_mark, delete_mark_keybinding, open_mark)
	local width = 50
	local height = #lines

	local buf = vim.api.nvim_create_buf(false, true)

	vim.api.nvim_buf_set_lines(buf, 0, -1, false, lines)

	local win = vim.api.nvim_open_win(buf, true, {
		relative = "editor",
		width = width,
		height = height,
		row = math.floor((vim.o.lines - height) / 2),
		col = math.floor((vim.o.columns - width) / 2),
		style = "minimal",
		border = "rounded",
	})

	-- Close window
	local function close()
		if vim.api.nvim_win_is_valid(win) then
			vim.api.nvim_win_close(win, true)
		end
	end

	-- Open selected mark
	vim.keymap.set("n", "<CR>", function()
		local row = vim.api.nvim_win_get_cursor(win)[1]
		local number = numbers[row]

		close()
		open_mark(number)
	end, {
		buffer = buf,
		desc = "Open marked file",
	})

	-- Delete selected mark
	vim.keymap.set("n", "d", function()
		local row = vim.api.nvim_win_get_cursor(win)[1]
		local number = numbers[row]

		delete_mark(number)
		delete_mark_keybinding(number)

		table.remove(numbers, row)
		table.remove(lines, row)

		vim.api.nvim_buf_set_lines(buf, 0, -1, false, lines)

		if #lines == 0 then
			close()
		end
	end, {
		buffer = buf,
		desc = "Delete mark",
	})

	-- Escape / q
	vim.keymap.set("n", "q", close, {
		buffer = buf,
		desc = "Close GitMarks",
	})

	vim.keymap.set("n", "<Esc>", close, {
		buffer = buf,
		desc = "Close GitMarks",
	})

	return buf, win
end

return M
