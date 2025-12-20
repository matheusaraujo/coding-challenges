local helpers = require("helpers")

return {
	part2 = function(puzzleInput)
		local dice = {}
		local track = {}
		local reading_dice = true

		for _, line in ipairs(puzzleInput) do
			if line == "" then
				reading_dice = false
			elseif reading_dice then
				table.insert(dice, helpers.Dice.new(line))
			else
				for char in line:gsub("%s+", ""):gmatch(".") do
					table.insert(track, tonumber(char))
				end
			end
		end

		local results = {}
		for i, die in ipairs(dice) do
			local total_rolls = 0
			for _, target in ipairs(track) do
				while die:roll() ~= target do
					total_rolls = total_rolls + 1
				end
			end
			table.insert(results, { id = i, rolls = total_rolls })
		end

		table.sort(results, function(a, b)
			return a.rolls < b.rolls
		end)

		local ids = {}
		for _, res in ipairs(results) do
			table.insert(ids, tostring(res.id))
		end
		return table.concat(ids, ",")
	end,
}
