local helpers = require("helpers")

return {
	part3 = function(puzzleInput)
		local machine, sequences = helpers.parseInput(puzzleInput)
		local numSlots = 20
		local numTokens = #sequences

		local scores = {}
		for i = 1, numTokens do
			scores[i] = {}
			for slot = 1, numSlots do
				local finalSlot = helpers.simulateToken(machine, sequences[i], slot)
				local coins = (finalSlot * 2) - slot
				scores[i][slot] = math.max(0, coins)
			end
		end

		local usedSlots = {}

		local function solve(tokenIdx, currentSum, findMax)
			if tokenIdx > numTokens then
				return currentSum
			end

			local bestVal = findMax and -1 or math.huge

			for slot = 1, numSlots do
				if not usedSlots[slot] then
					usedSlots[slot] = true
					local res = solve(tokenIdx + 1, currentSum + scores[tokenIdx][slot], findMax)
					if findMax then
						if res > bestVal then
							bestVal = res
						end
					else
						if res < bestVal then
							bestVal = res
						end
					end
					usedSlots[slot] = false
				end
			end
			return bestVal
		end

		local minScore = solve(1, 0, false)
		local maxScore = solve(1, 0, true)

		return string.format("%d %d", minScore, maxScore)
	end,
}
