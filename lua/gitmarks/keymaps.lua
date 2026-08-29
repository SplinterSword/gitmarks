local M = {}

function M.createFileBinding(gitmarks, number)
	vim.keymap.set({ "n", "v" }, "<leader>g" .. number - 1, function()
		gitmarks.openMarkedFile(number)
	end, {
		desc = "Open Marked file " .. number - 1,
	})
end

function M.setup(gitmarks)
	vim.keymap.set({ "n", "v" }, "<leader>g", "<Nop>", {
		desc = "[G]itmarks",
	})

	vim.keymap.set({ "n", "v" }, "<leader>gm", "<Nop>", {
		desc = "[M]ark Current File",
	})

	for i = 1, 10 do
		vim.keymap.set({ "n", "v" }, "<leader>gm" .. i - 1, function()
			gitmarks.markFile(i)
		end, {
			desc = "Mark file to " .. i - 1,
		})
	end
end

return M
