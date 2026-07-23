use crate::helpers::{monkey_business, parse_monkeys};
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let mut monkeys = parse_monkeys(puzzle_input);
    any(monkey_business(&mut monkeys, 20, |worry| worry / 3))
}
