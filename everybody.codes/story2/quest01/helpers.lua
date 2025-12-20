local M = {}

function M.parseInput(puzzleInput)
	local lines = {}
	if type(puzzleInput) == "table" then
		lines = puzzleInput
	else
		for line in puzzleInput:gmatch("[^\r\n]+") do
			table.insert(lines, line)
		end
	end

	local machine = {}
	local sequences = {}

	for _, line in ipairs(lines) do
		if line:match("^[LR%s]+$") then
			table.insert(sequences, (line:gsub("%s+", "")))
		elseif line:find("[%*%.]") then
			table.insert(machine, line)
		end
	end
	return machine, sequences
end

function M.simulateToken(machine, sequence, startSlot)
	local col = (startSlot * 2) - 1
	local seqIdx = 1
	local maxCol = #machine[1]

	for row = 1, #machine do
		local char = machine[row]:sub(col, col)
		if char == "*" then
			local dir = sequence:sub(seqIdx, seqIdx)
			seqIdx = seqIdx + 1

			if dir == "L" then
				if col == 1 then
					col = col + 1
				else
					col = col - 1
				end
			elseif dir == "R" then
				if col == maxCol then
					col = col - 1
				else
					col = col + 1
				end
			end
		end
	end

	return (col + 1) / 2
end

return M
