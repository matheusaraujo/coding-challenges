pub fn part1(puzzle_input: &[String]) -> String {
    let line = &puzzle_input[0];

    let open_parentheses = line.matches('(').count();
    let close_parentheses = line.matches(')').count();

    let result = open_parentheses as i32 - close_parentheses as i32;
    result.to_string()
}
