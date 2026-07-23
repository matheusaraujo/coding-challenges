use std::collections::HashSet;

type Pos = (i32, i32);

fn parse_moves(puzzle_input: &[String]) -> Vec<(char, i32)> {
    puzzle_input
        .iter()
        .filter(|line| !line.is_empty())
        .map(|line| {
            let mut parts = line.split_whitespace();
            let dir = parts.next().unwrap().chars().next().unwrap();
            let steps = parts.next().unwrap().parse().unwrap();
            (dir, steps)
        })
        .collect()
}

fn step_head((x, y): Pos, dir: char) -> Pos {
    match dir {
        'U' => (x, y + 1),
        'D' => (x, y - 1),
        'L' => (x - 1, y),
        'R' => (x + 1, y),
        _ => panic!("invalid direction: {dir}"),
    }
}

fn follow(head: Pos, tail: Pos) -> Pos {
    let dx = head.0 - tail.0;
    let dy = head.1 - tail.1;
    if dx.abs() <= 1 && dy.abs() <= 1 {
        tail
    } else {
        (tail.0 + dx.signum(), tail.1 + dy.signum())
    }
}

pub fn tail_visits(puzzle_input: &[String], num_knots: usize) -> usize {
    let moves = parse_moves(puzzle_input);
    let mut knots = vec![(0, 0); num_knots];
    let mut visited: HashSet<Pos> = HashSet::new();
    visited.insert(knots[num_knots - 1]);

    for (dir, steps) in moves {
        for _ in 0..steps {
            knots[0] = step_head(knots[0], dir);
            for i in 1..num_knots {
                knots[i] = follow(knots[i - 1], knots[i]);
            }
            visited.insert(knots[num_knots - 1]);
        }
    }

    visited.len()
}
