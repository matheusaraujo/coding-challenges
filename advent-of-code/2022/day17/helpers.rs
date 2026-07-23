use std::collections::{HashMap, HashSet};

const SHAPES: [&[(i32, i32)]; 5] = [
    &[(0, 0), (1, 0), (2, 0), (3, 0)],
    &[(1, 0), (0, 1), (1, 1), (2, 1), (1, 2)],
    &[(0, 0), (1, 0), (2, 0), (2, 1), (2, 2)],
    &[(0, 0), (0, 1), (0, 2), (0, 3)],
    &[(0, 0), (1, 0), (0, 1), (1, 1)],
];

pub fn parse_jets(puzzle_input: &[String]) -> Vec<u8> {
    puzzle_input
        .iter()
        .find(|line| !line.is_empty())
        .expect("jet pattern line")
        .bytes()
        .collect()
}

pub fn simulate(jets: &[u8], total_rocks: u64) -> u64 {
    let mut occupied: HashSet<(i32, i32)> = HashSet::new();
    let mut col_top = [0i32; 7];
    let mut top: i32 = 0;
    let mut jet_idx = 0usize;
    let mut rock_count = 0u64;
    let mut extra_height = 0u64;
    let mut seen: HashMap<(usize, usize, [i32; 7]), (u64, i32)> = HashMap::new();
    let mut skipped = false;

    while rock_count < total_rocks {
        let rock_type = (rock_count % 5) as usize;
        let shape = SHAPES[rock_type];

        let mut x = 2i32;
        let mut y = top + 3;

        loop {
            let dx = if jets[jet_idx] == b'<' { -1 } else { 1 };
            jet_idx = (jet_idx + 1) % jets.len();
            let new_x = x + dx;
            if shape.iter().all(|&(sx, sy)| {
                let nx = new_x + sx;
                (0..7).contains(&nx) && !occupied.contains(&(nx, y + sy))
            }) {
                x = new_x;
            }

            let new_y = y - 1;
            if new_y >= 0
                && shape
                    .iter()
                    .all(|&(sx, sy)| !occupied.contains(&(x + sx, new_y + sy)))
            {
                y = new_y;
            } else {
                break;
            }
        }

        for &(sx, sy) in shape {
            let (px, py) = (x + sx, y + sy);
            occupied.insert((px, py));
            top = top.max(py + 1);
            col_top[px as usize] = col_top[px as usize].max(py + 1);
        }
        rock_count += 1;

        if !skipped {
            let mut profile = [0i32; 7];
            for (col, height) in col_top.iter().enumerate() {
                profile[col] = top - height;
            }
            let key = (rock_type, jet_idx, profile);
            if let Some(&(prev_count, prev_top)) = seen.get(&key) {
                let cycle_len = rock_count - prev_count;
                let cycle_height = (top - prev_top) as u64;
                let remaining = total_rocks - rock_count;
                let cycles = remaining / cycle_len;
                extra_height += cycles * cycle_height;
                rock_count += cycles * cycle_len;
                skipped = true;
            } else {
                seen.insert(key, (rock_count, top));
            }
        }
    }

    top as u64 + extra_height
}
