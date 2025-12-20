return {
	part2 = function(puzzleInput)
		local floor = 0
		local input = puzzleInput[1]

		for i = 1, #input do
			local c = input:sub(i, i)

			if c == "(" then
				floor = floor + 1
			elseif c == ")" then
				floor = floor - 1
			end

			if floor == -1 then
				return i
			end
		end

		return 0
	end,
}
