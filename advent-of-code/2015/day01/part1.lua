return {
	part1 = function(puzzleInput)
		local line = puzzleInput[1]
		local _, openParentheses = line:gsub("%(", "")
		local _, closeParentheses = line:gsub("%)", "")
		return openParentheses - closeParentheses
	end,
}
