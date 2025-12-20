use super::helpers::TreeNode;
use super::helpers::solve;

pub fn part2(puzzle_input: &[String]) -> String {
    solve(puzzle_input, Some(swap), true)
}

fn swap(a: &mut TreeNode, b: &mut TreeNode) {
    std::mem::swap(&mut a.rank, &mut b.rank);
    std::mem::swap(&mut a.symbol, &mut b.symbol);
}
