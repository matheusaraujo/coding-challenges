use std::collections::HashSet;

pub struct Blizzards {
    height: i32,
    width: i32,
    pub start: (i32, i32),
    pub end: (i32, i32),
    list: Vec<(i32, i32, u8)>,
}

pub fn parse_blizzards(puzzle_input: &[String]) -> Blizzards {
    let lines: Vec<&str> = puzzle_input
        .iter()
        .filter(|line| !line.is_empty())
        .map(|s| s.as_str())
        .collect();
    let rows = lines.len();
    let cols = lines[0].len();
    let height = (rows - 2) as i32;
    let width = (cols - 2) as i32;

    let start_col = lines[0].chars().position(|c| c == '.').unwrap() as i32;
    let end_col = lines[rows - 1].chars().position(|c| c == '.').unwrap() as i32;

    let mut list = Vec::new();
    for (r, line) in lines.iter().enumerate() {
        for (c, ch) in line.chars().enumerate() {
            if matches!(ch, '^' | 'v' | '<' | '>') {
                list.push((r as i32, c as i32, ch as u8));
            }
        }
    }

    Blizzards {
        height,
        width,
        start: (0, start_col),
        end: ((rows - 1) as i32, end_col),
        list,
    }
}

impl Blizzards {
    fn occupied_at(&self, time: i32) -> HashSet<(i32, i32)> {
        self.list
            .iter()
            .map(|&(r, c, dir)| match dir {
                b'^' => (((r - 1 - time).rem_euclid(self.height)) + 1, c),
                b'v' => (((r - 1 + time).rem_euclid(self.height)) + 1, c),
                b'<' => (r, ((c - 1 - time).rem_euclid(self.width)) + 1),
                b'>' => (r, ((c - 1 + time).rem_euclid(self.width)) + 1),
                _ => unreachable!(),
            })
            .collect()
    }
}

const MOVES: [(i32, i32); 5] = [(0, 0), (-1, 0), (1, 0), (0, -1), (0, 1)];

pub fn min_time(bl: &Blizzards, from: (i32, i32), to: (i32, i32), start_time: i32) -> i32 {
    let mut time = start_time;
    let mut positions: HashSet<(i32, i32)> = HashSet::from([from]);

    loop {
        time += 1;
        let occupied = bl.occupied_at(time);
        let mut next = HashSet::new();

        for &(r, c) in &positions {
            for &(dr, dc) in &MOVES {
                let candidate = (r + dr, c + dc);
                if candidate == to {
                    return time;
                }
                if candidate == from {
                    next.insert(candidate);
                    continue;
                }
                let (nr, nc) = candidate;
                if nr < 1 || nr > bl.height || nc < 1 || nc > bl.width {
                    continue;
                }
                if !occupied.contains(&candidate) {
                    next.insert(candidate);
                }
            }
        }

        positions = next;
    }
}
