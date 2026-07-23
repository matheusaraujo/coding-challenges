use crate::helpers::{min_time, parse_blizzards};
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let bl = parse_blizzards(puzzle_input);
    let there = min_time(&bl, bl.start, bl.end, 0);
    let back = min_time(&bl, bl.end, bl.start, there);
    let there_again = min_time(&bl, bl.start, bl.end, back);
    any(there_again)
}
