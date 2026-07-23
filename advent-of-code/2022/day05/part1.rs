use crate::helpers::{apply_moves, parse_input, top_crates};
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let (mut stacks, instructions) = parse_input(puzzle_input);
    apply_moves(&mut stacks, &instructions, |from, to, count| {
        for _ in 0..count {
            if let Some(c) = from.pop() {
                to.push(c);
            }
        }
    });
    any(top_crates(&stacks))
}
