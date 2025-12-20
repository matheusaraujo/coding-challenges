local helpers = require("helpers")

return {
	part3 = function(puzzleInput)
		local dice = {}
		local grid = {}
		local reading_dice = true

		for _, line in ipairs(puzzleInput) do
			if line == "" then
				reading_dice = false
			elseif reading_dice then
				table.insert(dice, helpers.Dice.new(line))
			else
				local row = {}
				for i = 1, #line do
					table.insert(row, line:sub(i, i))
				end
				table.insert(grid, row)
			end
		end

		local height = #grid
		local width = #grid[1]
		local seen = {}
		for y = 1, height do
			seen[y] = {}
		end
		local version = 0

		for _, die in ipairs(dice) do
			local roll = tostring(die:roll())
			version = version + 1
			local todo = {}

			for y = 1, height do
				for x = 1, width do
					if grid[y][x] == roll then
						table.insert(todo, { x = x, y = y })
						seen[y][x] = version
					end
				end
			end

			while #todo > 0 do
				local next_todo = {}
				local roll_val = tostring(die:roll())
				version = version + 1

				for _, p in ipairs(todo) do
					local neighbors = {
						{ x = p.x, y = p.y },
						{ x = p.x + 1, y = p.y },
						{ x = p.x - 1, y = p.y },
						{ x = p.x, y = p.y + 1 },
						{ x = p.x, y = p.y - 1 },
					}

					for _, n in ipairs(neighbors) do
						if n.y >= 1 and n.y <= height and n.x >= 1 and n.x <= width then
							if grid[n.y][n.x] == roll_val and seen[n.y][n.x] ~= version then
								seen[n.y][n.x] = version
								table.insert(next_todo, n)
							end
						end
					end
				end
				todo = next_todo
			end
		end

		local count = 0
		for y = 1, height do
			for x = 1, width do
				if (seen[y][x] or 0) > 0 then
					count = count + 1
				end
			end
		end
		return count
	end,
}
