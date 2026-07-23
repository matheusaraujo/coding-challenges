use crate::helpers::{parse_sensors, row_intervals};
use crate::{Any, any};
use std::collections::HashSet;

pub fn part1(puzzle_input: &[String]) -> Any {
    let sensors = parse_sensors(puzzle_input);
    let row = 2_000_000;
    let intervals = row_intervals(&sensors, row);

    let covered: i64 = intervals.iter().map(|(start, end)| end - start + 1).sum();
    let beacons_on_row: HashSet<i64> = sensors
        .iter()
        .filter(|sensor| sensor.beacon_y == row)
        .map(|sensor| sensor.beacon_x)
        .collect();
    let beacons_in_range = beacons_on_row
        .iter()
        .filter(|&&bx| {
            intervals
                .iter()
                .any(|&(start, end)| bx >= start && bx <= end)
        })
        .count() as i64;

    any(covered - beacons_in_range)
}
