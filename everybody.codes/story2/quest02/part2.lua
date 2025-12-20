local helpers = require("helpers")

return {
	part2 = function(puzzleInput)
		local input = table.concat(puzzleInput, "\n")
		return helpers.process_circle(input, 100)
	end,
}
