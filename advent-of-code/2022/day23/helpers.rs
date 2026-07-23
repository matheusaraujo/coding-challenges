use std::collections::{HashMap, HashSet};

pub type Pos = (i32, i32);

pub fn parse_elves(puzzle_input: &[String]) -> HashSet<Pos> {
    puzzle_input
        .iter()
        .filter(|line| !line.is_empty())
        .enumerate()
        .flat_map(|(r, line)| {
            line.chars()
                .enumerate()
                .filter(move |&(_, ch)| ch == '#')
                .map(move |(c, _)| (r as i32, c as i32))
        })
        .collect()
}

const DIRS8: [Pos; 8] = [
    (-1, -1),
    (-1, 0),
    (-1, 1),
    (0, -1),
    (0, 1),
    (1, -1),
    (1, 0),
    (1, 1),
];

const CHECKS: [[Pos; 3]; 4] = [
    [(-1, -1), (-1, 0), (-1, 1)],
    [(1, -1), (1, 0), (1, 1)],
    [(-1, -1), (0, -1), (1, -1)],
    [(-1, 1), (0, 1), (1, 1)],
];

const MOVE_DELTA: [Pos; 4] = [(-1, 0), (1, 0), (0, -1), (0, 1)];

pub fn simulate_round(elves: &HashSet<Pos>, start_dir: usize) -> (HashSet<Pos>, bool) {
    let mut proposals: HashMap<Pos, Vec<Pos>> = HashMap::new();

    for &(r, c) in elves {
        let has_neighbor = DIRS8
            .iter()
            .any(|&(dr, dc)| elves.contains(&(r + dr, c + dc)));
        if !has_neighbor {
            continue;
        }

        for i in 0..4 {
            let dir_idx = (start_dir + i) % 4;
            if CHECKS[dir_idx]
                .iter()
                .all(|&(dr, dc)| !elves.contains(&(r + dr, c + dc)))
            {
                let (mdr, mdc) = MOVE_DELTA[dir_idx];
                proposals
                    .entry((r + mdr, c + mdc))
                    .or_default()
                    .push((r, c));
                break;
            }
        }
    }

    let mut moved = false;
    let mut new_elves = elves.clone();
    for (dest, sources) in proposals {
        if sources.len() == 1 {
            new_elves.remove(&sources[0]);
            new_elves.insert(dest);
            moved = true;
        }
    }

    (new_elves, moved)
}

pub fn count_empty_ground(elves: &HashSet<Pos>) -> i64 {
    let min_r = elves.iter().map(|p| p.0).min().unwrap();
    let max_r = elves.iter().map(|p| p.0).max().unwrap();
    let min_c = elves.iter().map(|p| p.1).min().unwrap();
    let max_c = elves.iter().map(|p| p.1).max().unwrap();
    let area = (max_r - min_r + 1) as i64 * (max_c - min_c + 1) as i64;
    area - elves.len() as i64
}
