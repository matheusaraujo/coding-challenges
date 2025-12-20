use super::helpers::SnailClock;

pub fn part1(puzzle_input: &[String]) -> String {
    let clock = SnailClock::from_input(puzzle_input);
    let sum: i128 = clock
        .snails
        .iter()
        .map(|s| {
            let s = s.clone();
            let s_size = s.disc_size();
            let nx = ((s.x - 1 + 100) % s_size) + 1;
            let ny = s_size - (nx - 1);
            nx + (100 * ny)
        })
        .sum();
    sum.to_string()
}
