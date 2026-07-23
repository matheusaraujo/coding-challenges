pub struct Sensor {
    pub x: i64,
    pub y: i64,
    pub beacon_x: i64,
    pub beacon_y: i64,
    pub radius: i64,
}

pub fn parse_sensors(puzzle_input: &[String]) -> Vec<Sensor> {
    puzzle_input
        .iter()
        .filter(|line| !line.is_empty())
        .map(|line| {
            let nums: Vec<i64> = line
                .split(|c: char| !c.is_ascii_digit() && c != '-')
                .filter(|s| !s.is_empty())
                .map(|s| s.parse().unwrap())
                .collect();
            let (x, y, beacon_x, beacon_y) = (nums[0], nums[1], nums[2], nums[3]);
            let radius = (x - beacon_x).abs() + (y - beacon_y).abs();
            Sensor {
                x,
                y,
                beacon_x,
                beacon_y,
                radius,
            }
        })
        .collect()
}

fn merge_intervals(mut intervals: Vec<(i64, i64)>) -> Vec<(i64, i64)> {
    intervals.sort_unstable();
    let mut merged: Vec<(i64, i64)> = Vec::new();
    for (start, end) in intervals {
        if let Some(last) = merged.last_mut()
            && start <= last.1 + 1
        {
            last.1 = last.1.max(end);
            continue;
        }
        merged.push((start, end));
    }
    merged
}

pub fn row_intervals(sensors: &[Sensor], row: i64) -> Vec<(i64, i64)> {
    let intervals = sensors
        .iter()
        .filter_map(|sensor| {
            let spread = sensor.radius - (sensor.y - row).abs();
            (spread >= 0).then_some((sensor.x - spread, sensor.x + spread))
        })
        .collect();
    merge_intervals(intervals)
}

pub fn find_distress_beacon(sensors: &[Sensor], bound: i64) -> Option<(i64, i64)> {
    for y in 0..=bound {
        let intervals = row_intervals(sensors, y);
        let mut x = 0i64;
        for &(start, end) in &intervals {
            if start > x {
                return Some((x, y));
            }
            x = x.max(end + 1);
            if x > bound {
                break;
            }
        }
        if x <= bound {
            return Some((x, y));
        }
    }
    None
}
