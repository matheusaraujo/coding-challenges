use crate::helpers::elf_calorie_totals;
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let mut totals = elf_calorie_totals(puzzle_input);
    totals.sort_unstable_by(|a, b| b.cmp(a));
    let top_three: u32 = totals.into_iter().take(3).sum();
    any(top_three)
}
