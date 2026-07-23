use crate::helpers::{parse_elves, simulate_round};
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let mut elves = parse_elves(puzzle_input);
    let mut round = 0;
    loop {
        let (new_elves, moved) = simulate_round(&elves, round % 4);
        round += 1;
        if !moved {
            break;
        }
        elves = new_elves;
    }
    any(round)
}
