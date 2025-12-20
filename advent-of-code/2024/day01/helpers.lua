local M = {}

function M.parseInput(puzzleInput)
	local left, right = {}, {}

	for _, line in ipairs(puzzleInput) do
		local parts = {}
		for s in line:gmatch("%S+") do
			table.insert(parts, s)
		end

		table.insert(left, tonumber(parts[1]))
		table.insert(right, tonumber(parts[2]))
	end

	table.sort(left)
	table.sort(right)

	return left, right
end

return M
