pub fn part2(puzzle_input: &[String]) -> String {
    let mut floor = 0;

    for (i, c) in puzzle_input[0].chars().enumerate() {
        match c {
            '(' => floor += 1,
            ')' => floor -= 1,
            _ => {}
        }

        if floor == -1 {
            return (i + 1).to_string();
        }
    }

    "0".to_string()
}
