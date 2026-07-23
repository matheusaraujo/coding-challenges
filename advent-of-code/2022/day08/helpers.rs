pub fn parse_grid(puzzle_input: &[String]) -> Vec<Vec<u32>> {
    puzzle_input
        .iter()
        .filter(|line| !line.is_empty())
        .map(|line| line.chars().filter_map(|c| c.to_digit(10)).collect())
        .collect()
}

pub fn sightlines(grid: &[Vec<u32>], row: usize, col: usize) -> [Vec<u32>; 4] {
    let height = grid.len();
    let width = grid[0].len();

    let up = (0..row).rev().map(|r| grid[r][col]).collect();
    let down = (row + 1..height).map(|r| grid[r][col]).collect();
    let left = (0..col).rev().map(|c| grid[row][c]).collect();
    let right = (col + 1..width).map(|c| grid[row][c]).collect();

    [up, down, left, right]
}
