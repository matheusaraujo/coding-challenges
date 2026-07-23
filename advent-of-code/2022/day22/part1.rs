use crate::helpers::{parse_board_and_path, password, walk_flat};
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let (board, path) = parse_board_and_path(puzzle_input);
    let (row, col, facing) = walk_flat(&board, &path);
    any(password(row, col, facing))
}
