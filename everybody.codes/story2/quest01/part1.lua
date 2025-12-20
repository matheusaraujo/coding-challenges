local helpers = require("helpers")

return {
	part1 = function(puzzleInput)
		local machine, sequences = helpers.parseInput(puzzleInput)
		local totalCoins = 0

		for i, seq in ipairs(sequences) do
			local tossSlot = i
			local finalSlot = helpers.simulateToken(machine, seq, tossSlot)
			local coins = (finalSlot * 2) - tossSlot
			if coins < 0 then
				coins = 0
			end
			totalCoins = totalCoins + coins
		end

		return totalCoins
	end,
}
