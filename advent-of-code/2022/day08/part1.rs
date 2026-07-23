use crate::helpers::{parse_grid, sightlines};
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let grid = parse_grid(puzzle_input);

    let count = grid
        .iter()
        .enumerate()
        .flat_map(|(row, line)| line.iter().enumerate().map(move |(col, &h)| (row, col, h)))
        .filter(|&(row, col, height)| {
            sightlines(&grid, row, col)
                .iter()
                .any(|line| line.iter().all(|&h| h < height))
        })
        .count();
    any(count)
}
