local helpers = require("helpers")

return {
	part1 = function(puzzleInput)
		local left, right = helpers.parseInput(puzzleInput)
		local sum = 0

		for i = 1, #left do
			sum = sum + math.abs(left[i] - right[i])
		end

		return sum
	end,
}
