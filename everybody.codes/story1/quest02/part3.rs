use super::helpers::TreeNode;
use super::helpers::solve;
use crate::{Any, any};

pub fn part3(puzzle_input: &[String]) -> Any {
    any(solve(puzzle_input, Some(swap), false))
}

fn swap(a: &mut TreeNode, b: &mut TreeNode) {
    std::mem::swap(a, b);
}
