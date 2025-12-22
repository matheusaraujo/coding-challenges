use super::helpers::TreeNode;
use super::helpers::solve;
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    any(solve(puzzle_input, Some(swap), true))
}

fn swap(a: &mut TreeNode, b: &mut TreeNode) {
    std::mem::swap(&mut a.rank, &mut b.rank);
    std::mem::swap(&mut a.symbol, &mut b.symbol);
}
