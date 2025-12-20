local helpers = require("helpers")

return {
	part2 = function(puzzleInput)
		local left, right = helpers.parseInput(puzzleInput)

		local count = {}
		for _, x in ipairs(left) do
			count[x] = (count[x] or 0) + 1
		end

		local result = 0
		for _, item in ipairs(right) do
			if count[item] then
				result = result + (item * count[item])
			end
		end

		return result
	end,
}
