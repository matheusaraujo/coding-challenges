use crate::helpers::{monkey_business, parse_monkeys};
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let mut monkeys = parse_monkeys(puzzle_input);
    let modulus: u64 = monkeys.iter().map(|monkey| monkey.divisor).product();
    any(monkey_business(&mut monkeys, 10_000, |worry| {
        worry % modulus
    }))
}
