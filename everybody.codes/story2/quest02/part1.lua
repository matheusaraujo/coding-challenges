return {
	part1 = function(puzzleInput)
		local input = table.concat(puzzleInput, ""):gsub("%s+", "")
		local bolts = 0
		local i = 1
		local sequence = { "R", "G", "B" }

		while i <= #input do
			for _, color in ipairs(sequence) do
				while i <= #input and input:sub(i, i) == color do
					i = i + 1
				end
				if i <= #input then
					i = i + 1
					bolts = bolts + 1
				end
			end
		end
		return bolts
	end,
}
