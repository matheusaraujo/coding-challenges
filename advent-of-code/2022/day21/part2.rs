use crate::helpers::{children, depends_on_humn, eval, parse_monkeys, solve};
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let monkeys = parse_monkeys(puzzle_input);
    let (left, right) = children(&monkeys, "root");

    let target = if depends_on_humn(&monkeys, &left) {
        solve(&monkeys, &left, eval(&monkeys, &right))
    } else {
        solve(&monkeys, &right, eval(&monkeys, &left))
    };
    any(target)
}
