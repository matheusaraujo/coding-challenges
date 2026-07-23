use crate::helpers::{eval, parse_monkeys};
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let monkeys = parse_monkeys(puzzle_input);
    any(eval(&monkeys, "root"))
}
