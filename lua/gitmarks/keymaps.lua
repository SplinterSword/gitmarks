local M = {}

function M.createFileBinding(gitmarks, number)
	vim.keymap.set({ "n", "v" }, "<leader>g" .. number, function()
		gitmarks.openMarkedFile(number)
	end, {
		desc = "Open Marked file " .. number,
	})
end

function M.setup(gitmarks)
	vim.keymap.set({ "n", "v" }, "<leader>g", "<Nop>", {
		desc = "[G]itmarks",
	})

	vim.keymap.set({ "n", "v" }, "<leader>gm", "<Nop>", {
		desc = "[M]ark Current File",
	})

	vim.keymap.set({ "n", "v" }, "<leader>gl", function()
		gitmarks.openViewTab()
	end, {
		desc = "Open Marked file list",
	})

	for i = 1, 9 do
		vim.keymap.set({ "n", "v" }, "<leader>gm" .. i, function()
			gitmarks.markFile(i)
		end, {
			desc = "Mark file to " .. i,
		})
	end
end

return M
