use super::helpers::TreeNode;
use super::helpers::solve;
use crate::{Answer, answer};

pub fn part3(puzzle_input: &[String]) -> Answer {
    answer(solve(puzzle_input, Some(swap), false))
}

fn swap(a: &mut TreeNode, b: &mut TreeNode) {
    std::mem::swap(a, b);
}
