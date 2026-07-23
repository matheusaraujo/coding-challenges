use crate::helpers::{count_empty_ground, parse_elves, simulate_round};
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let mut elves = parse_elves(puzzle_input);
    for round in 0..10 {
        let (new_elves, _) = simulate_round(&elves, round % 4);
        elves = new_elves;
    }
    any(count_empty_ground(&elves))
}
