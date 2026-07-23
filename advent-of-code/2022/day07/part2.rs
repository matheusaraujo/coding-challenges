use crate::helpers::directory_sizes;
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let sizes = directory_sizes(puzzle_input);
    let used = sizes[&vec!["/".to_string()]];
    let unused = 70_000_000 - used;
    let needed = 30_000_000 - unused;

    let smallest = sizes
        .values()
        .filter(|&&size| size >= needed)
        .min()
        .expect("a directory large enough exists");
    any(*smallest)
}
