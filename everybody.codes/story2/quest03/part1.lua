local helpers = require("helpers")

return {
	part1 = function(puzzleInput)
		local dice = {}
		for _, line in ipairs(puzzleInput) do
			table.insert(dice, helpers.Dice.new(line))
		end

		local total = 0
		local rolls = 0
		while total < 10000 do
			for _, d in ipairs(dice) do
				total = total + d:roll()
			end
			rolls = rolls + 1
		end
		return rolls
	end,
}
