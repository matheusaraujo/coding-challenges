use std::collections::{HashSet, VecDeque};

type Cube = (i32, i32, i32);

const NEIGHBORS: [Cube; 6] = [
    (1, 0, 0),
    (-1, 0, 0),
    (0, 1, 0),
    (0, -1, 0),
    (0, 0, 1),
    (0, 0, -1),
];

pub fn parse_cubes(puzzle_input: &[String]) -> HashSet<Cube> {
    puzzle_input
        .iter()
        .filter(|line| !line.is_empty())
        .map(|line| {
            let coords: Vec<i32> = line.split(',').map(|n| n.parse().unwrap()).collect();
            (coords[0], coords[1], coords[2])
        })
        .collect()
}

pub fn surface_area(cubes: &HashSet<Cube>) -> usize {
    cubes
        .iter()
        .map(|&(x, y, z)| {
            NEIGHBORS
                .iter()
                .filter(|&&(dx, dy, dz)| !cubes.contains(&(x + dx, y + dy, z + dz)))
                .count()
        })
        .sum()
}

fn bounds(cubes: &HashSet<Cube>) -> (Cube, Cube) {
    let min_x = cubes.iter().map(|c| c.0).min().unwrap() - 1;
    let max_x = cubes.iter().map(|c| c.0).max().unwrap() + 1;
    let min_y = cubes.iter().map(|c| c.1).min().unwrap() - 1;
    let max_y = cubes.iter().map(|c| c.1).max().unwrap() + 1;
    let min_z = cubes.iter().map(|c| c.2).min().unwrap() - 1;
    let max_z = cubes.iter().map(|c| c.2).max().unwrap() + 1;
    ((min_x, min_y, min_z), (max_x, max_y, max_z))
}

pub fn exterior_surface_area(cubes: &HashSet<Cube>) -> usize {
    let (min, max) = bounds(cubes);
    let mut outside: HashSet<Cube> = HashSet::new();
    let mut queue = VecDeque::new();

    outside.insert(min);
    queue.push_back(min);

    while let Some((x, y, z)) = queue.pop_front() {
        for &(dx, dy, dz) in &NEIGHBORS {
            let next = (x + dx, y + dy, z + dz);
            let in_bounds = (min.0..=max.0).contains(&next.0)
                && (min.1..=max.1).contains(&next.1)
                && (min.2..=max.2).contains(&next.2);
            if in_bounds && !cubes.contains(&next) && outside.insert(next) {
                queue.push_back(next);
            }
        }
    }

    cubes
        .iter()
        .map(|&(x, y, z)| {
            NEIGHBORS
                .iter()
                .filter(|&&(dx, dy, dz)| outside.contains(&(x + dx, y + dy, z + dz)))
                .count()
        })
        .sum()
}
