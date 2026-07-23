use std::collections::HashSet;

pub fn parse_rocks(puzzle_input: &[String]) -> (HashSet<(i32, i32)>, i32) {
    let mut rocks = HashSet::new();
    let mut max_y = 0;

    for line in puzzle_input.iter().filter(|line| !line.is_empty()) {
        let points: Vec<(i32, i32)> = line
            .split(" -> ")
            .map(|point| {
                let mut coords = point.split(',');
                let x: i32 = coords.next().unwrap().parse().unwrap();
                let y: i32 = coords.next().unwrap().parse().unwrap();
                (x, y)
            })
            .collect();

        for pair in points.windows(2) {
            let (x1, y1) = pair[0];
            let (x2, y2) = pair[1];
            max_y = max_y.max(y1).max(y2);

            if x1 == x2 {
                for y in y1.min(y2)..=y1.max(y2) {
                    rocks.insert((x1, y));
                }
            } else {
                for x in x1.min(x2)..=x1.max(x2) {
                    rocks.insert((x, y1));
                }
            }
        }
    }

    (rocks, max_y)
}

pub enum Step {
    Open,
    Blocked,
    Void,
}

pub fn simulate_sand(
    rocks: &HashSet<(i32, i32)>,
    step: impl Fn(&HashSet<(i32, i32)>, i32, i32) -> Step,
) -> usize {
    let mut occupied = rocks.clone();
    let mut count = 0;

    loop {
        if matches!(step(&occupied, 500, 0), Step::Blocked) {
            break;
        }

        let (mut x, mut y) = (500, 0);
        loop {
            match step(&occupied, x, y + 1) {
                Step::Void => return count,
                Step::Open => y += 1,
                Step::Blocked => match step(&occupied, x - 1, y + 1) {
                    Step::Void => return count,
                    Step::Open => {
                        x -= 1;
                        y += 1;
                    }
                    Step::Blocked => match step(&occupied, x + 1, y + 1) {
                        Step::Void => return count,
                        Step::Open => {
                            x += 1;
                            y += 1;
                        }
                        Step::Blocked => {
                            occupied.insert((x, y));
                            count += 1;
                            break;
                        }
                    },
                },
            }
        }
    }

    count
}
