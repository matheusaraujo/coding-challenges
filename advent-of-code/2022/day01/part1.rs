use crate::helpers::elf_calorie_totals;
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let max = elf_calorie_totals(puzzle_input)
        .into_iter()
        .max()
        .unwrap_or(0);
    any(max)
}
