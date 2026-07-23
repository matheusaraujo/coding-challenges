use std::collections::VecDeque;

pub struct Grid {
    pub heights: Vec<Vec<i32>>,
    pub start: (usize, usize),
    pub end: (usize, usize),
}

pub fn parse_grid(puzzle_input: &[String]) -> Grid {
    let mut start = (0, 0);
    let mut end = (0, 0);

    let heights = puzzle_input
        .iter()
        .filter(|line| !line.is_empty())
        .enumerate()
        .map(|(row, line)| {
            line.chars()
                .enumerate()
                .map(|(col, c)| match c {
                    'S' => {
                        start = (row, col);
                        0
                    }
                    'E' => {
                        end = (row, col);
                        25
                    }
                    c => c as i32 - 'a' as i32,
                })
                .collect()
        })
        .collect();

    Grid {
        heights,
        start,
        end,
    }
}

fn neighbors(r: usize, c: usize, rows: usize, cols: usize) -> Vec<(usize, usize)> {
    let mut result = Vec::new();
    if r > 0 {
        result.push((r - 1, c));
    }
    if r + 1 < rows {
        result.push((r + 1, c));
    }
    if c > 0 {
        result.push((r, c - 1));
    }
    if c + 1 < cols {
        result.push((r, c + 1));
    }
    result
}

pub fn distances_from_end(grid: &Grid) -> Vec<Vec<i32>> {
    let rows = grid.heights.len();
    let cols = grid.heights[0].len();
    let mut dist = vec![vec![-1; cols]; rows];
    let mut queue = VecDeque::new();

    dist[grid.end.0][grid.end.1] = 0;
    queue.push_back(grid.end);

    while let Some((r, c)) = queue.pop_front() {
        let current_height = grid.heights[r][c];
        for (nr, nc) in neighbors(r, c, rows, cols) {
            if dist[nr][nc] == -1 && grid.heights[nr][nc] >= current_height - 1 {
                dist[nr][nc] = dist[r][c] + 1;
                queue.push_back((nr, nc));
            }
        }
    }

    dist
}
