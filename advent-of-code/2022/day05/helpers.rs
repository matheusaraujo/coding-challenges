pub struct Instruction {
    pub count: usize,
    pub from: usize,
    pub to: usize,
}

pub fn parse_input(puzzle_input: &[String]) -> (Vec<Vec<char>>, Vec<Instruction>) {
    let blank_idx = puzzle_input
        .iter()
        .position(|line| line.is_empty())
        .expect("blank separator line");
    let drawing = &puzzle_input[..blank_idx];
    let instruction_lines = &puzzle_input[blank_idx + 1..];

    let label_line = drawing.last().expect("stack label line");
    let num_stacks = label_line.split_whitespace().count();
    let mut stacks: Vec<Vec<char>> = vec![Vec::new(); num_stacks];

    for line in drawing[..drawing.len() - 1].iter().rev() {
        let chars: Vec<char> = line.chars().collect();
        for (i, stack) in stacks.iter_mut().enumerate() {
            let pos = 1 + i * 4;
            if let Some(&c) = chars.get(pos)
                && c != ' '
            {
                stack.push(c);
            }
        }
    }

    let instructions = instruction_lines
        .iter()
        .filter(|line| !line.is_empty())
        .map(|line| {
            let nums: Vec<usize> = line
                .split_whitespace()
                .filter_map(|token| token.parse().ok())
                .collect();
            Instruction {
                count: nums[0],
                from: nums[1] - 1,
                to: nums[2] - 1,
            }
        })
        .collect();

    (stacks, instructions)
}

pub fn top_crates(stacks: &[Vec<char>]) -> String {
    stacks.iter().filter_map(|stack| stack.last()).collect()
}

pub fn apply_moves(
    stacks: &mut [Vec<char>],
    instructions: &[Instruction],
    move_group: impl Fn(&mut Vec<char>, &mut Vec<char>, usize),
) {
    for instruction in instructions {
        let (from, to) = if instruction.from < instruction.to {
            let (left, right) = stacks.split_at_mut(instruction.to);
            (&mut left[instruction.from], &mut right[0])
        } else {
            let (left, right) = stacks.split_at_mut(instruction.from);
            (&mut right[0], &mut left[instruction.to])
        };
        move_group(from, to, instruction.count);
    }
}
