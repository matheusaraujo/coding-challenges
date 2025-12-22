use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let mut floor = 1;

    for (i, c) in puzzle_input[0].chars().enumerate() {
        match c {
            '(' => floor += 1,
            ')' => floor -= 1,
            _ => {}
        }

        if floor == -1 {
            return any(i);
        }
    }

    any(-1)
}
