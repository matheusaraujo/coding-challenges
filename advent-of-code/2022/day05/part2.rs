use crate::helpers::{apply_moves, parse_input, top_crates};
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let (mut stacks, instructions) = parse_input(puzzle_input);
    apply_moves(&mut stacks, &instructions, |from, to, count| {
        let split_point = from.len() - count;
        to.extend(from.split_off(split_point));
    });
    any(top_crates(&stacks))
}
