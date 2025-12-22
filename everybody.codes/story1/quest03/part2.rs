use crate::{Any, any};
use super::helpers::SnailClock;

pub fn part2(puzzle_input: &[String]) -> Any {
    let clock = SnailClock::from_input(puzzle_input);
    any(clock.find_first_alignment())
}
