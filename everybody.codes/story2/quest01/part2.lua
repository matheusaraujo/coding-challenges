local helpers = require("helpers")

return {
	part2 = function(puzzleInput)
		local machine, sequences = helpers.parseInput(puzzleInput)

		local numSlots = 13
		local totalMaxCoins = 0

		for _, seq in ipairs(sequences) do
			local bestCoinsForToken = 0

			for tossSlot = 1, numSlots do
				local finalSlot = helpers.simulateToken(machine, seq, tossSlot)

				local coins = (finalSlot * 2) - tossSlot
				if coins < 0 then
					coins = 0
				end

				if coins > bestCoinsForToken then
					bestCoinsForToken = coins
				end
			end

			totalMaxCoins = totalMaxCoins + bestCoinsForToken
		end

		return totalMaxCoins
	end,
}
