local helpers = {}

local Dice = {}
Dice.__index = Dice

function Dice.new(line)
	local self = setmetatable({}, Dice)
	local faces = {}

	for val in line:gmatch("-?%d+") do
		table.insert(faces, tonumber(val))
	end

	self.seed = table.remove(faces, #faces)
	table.remove(faces, 1)

	self.faces = faces
	self.pulse = self.seed
	self.index = 0
	self.roll_number = 1
	return self
end

function Dice:roll()
	local spin = self.roll_number * self.pulse
	self.index = (self.index + spin) % #self.faces
	self.pulse = (self.pulse + spin) % self.seed + (1 + self.roll_number + self.seed)
	self.roll_number = self.roll_number + 1
	return self.faces[self.index + 1]
end

helpers.Dice = Dice
return helpers
