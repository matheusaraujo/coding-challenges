local helpers = require("helpers")

return {
	part3 = function(puzzleInput)
		local input = table.concat(puzzleInput, "\n")
		return helpers.process_circle(input, 100000)
	end,
}
