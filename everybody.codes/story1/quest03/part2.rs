use super::helpers::SnailClock;

pub fn part2(puzzle_input: &[String]) -> String {
    let clock = SnailClock::from_input(puzzle_input);
    clock.find_first_alignment().to_string()
}
