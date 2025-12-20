use super::helpers::TreeNode;
use super::helpers::solve;

pub fn part3(puzzle_input: &[String]) -> String {
    solve(puzzle_input, Some(swap), false)
}

fn swap(a: &mut TreeNode, b: &mut TreeNode) {
    std::mem::swap(a, b);
}
