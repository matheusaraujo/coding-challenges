local M = {}

function M.process_circle(line_str, repeats)
	local line_len = #line_str
	local total_len = line_len * repeats
	local mid = math.floor(total_len / 2)

	local left = {}
	local right = {}

	for i = 1, mid do
		local char_idx = ((i - 1) % line_len) + 1
		left[i] = line_str:byte(char_idx)
	end

	for i = mid + 1, total_len do
		local char_idx = ((i - 1) % line_len) + 1
		right[i - mid] = line_str:byte(char_idx)
	end

	local left_idx, right_idx = 1, 1
	local left_len, right_len = #left, #right

	local count = 0
	local colors = { string.byte("R"), string.byte("G"), string.byte("B") }
	local color_idx = 1

	while left_len > 0 do
		count = count + 1
		local f = colors[color_idx]

		local current_left = left[left_idx]
		left_idx = left_idx + 1
		left_len = left_len - 1

		if current_left == f then
			if (left_len + right_len) % 2 ~= 0 then
				right_idx = right_idx + 1
				right_len = right_len - 1
			end
		elseif right_len > left_len then
			left[left_idx + left_len] = right[right_idx]
			left_len = left_len + 1
			right_idx = right_idx + 1
			right_len = right_len - 1
		end

		color_idx = (color_idx % 3) + 1
	end

	return count
end

return M
