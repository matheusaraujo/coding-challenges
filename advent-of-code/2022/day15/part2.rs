use crate::helpers::{find_distress_beacon, parse_sensors};
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let sensors = parse_sensors(puzzle_input);
    let (x, y) = find_distress_beacon(&sensors, 4_000_000).expect("distress beacon exists");
    any(x * 4_000_000 + y)
}
