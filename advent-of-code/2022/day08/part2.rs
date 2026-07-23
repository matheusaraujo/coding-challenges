use crate::helpers::{parse_grid, sightlines};
use crate::{Any, any};

fn viewing_distance(tree_height: u32, line: &[u32]) -> usize {
    let mut distance = 0;
    for &height in line {
        distance += 1;
        if height >= tree_height {
            break;
        }
    }
    distance
}

pub fn part2(puzzle_input: &[String]) -> Any {
    let grid = parse_grid(puzzle_input);

    let max_score = grid
        .iter()
        .enumerate()
        .flat_map(|(row, line)| line.iter().enumerate().map(move |(col, &h)| (row, col, h)))
        .map(|(row, col, height)| {
            sightlines(&grid, row, col)
                .iter()
                .map(|line| viewing_distance(height, line))
                .product::<usize>()
        })
        .max()
        .unwrap_or(0);
    any(max_score)
}
