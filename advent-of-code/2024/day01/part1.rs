use super::helpers::parse_input;

pub fn part1(puzzle_input: &[String]) -> String {
    let (left, right) = parse_input(puzzle_input);
    let mut sum: i32 = 0;

    for (l, r) in left.iter().zip(right.iter()) {
        sum += (l - r).abs();
    }

    sum.to_string()
}
