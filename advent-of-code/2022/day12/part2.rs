use crate::helpers::{distances_from_end, parse_grid};
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let grid = parse_grid(puzzle_input);
    let dist = distances_from_end(&grid);

    let best = grid
        .heights
        .iter()
        .enumerate()
        .flat_map(|(r, row)| row.iter().enumerate().map(move |(c, &h)| (r, c, h)))
        .filter(|&(_, _, h)| h == 0)
        .filter_map(|(r, c, _)| {
            let d = dist[r][c];
            (d >= 0).then_some(d)
        })
        .min()
        .expect("a reachable lowest point exists");
    any(best)
}
