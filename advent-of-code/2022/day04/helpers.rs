pub type Range = (u32, u32);

pub fn parse_pair(line: &str) -> (Range, Range) {
    let mut sections = line.split(',').map(parse_range);
    (sections.next().unwrap(), sections.next().unwrap())
}

fn parse_range(s: &str) -> Range {
    let mut bounds = s.split('-').map(|n| n.parse::<u32>().unwrap());
    (bounds.next().unwrap(), bounds.next().unwrap())
}

pub fn fully_contains((a_start, a_end): Range, (b_start, b_end): Range) -> bool {
    (a_start <= b_start && a_end >= b_end) || (b_start <= a_start && b_end >= a_end)
}

pub fn overlaps((a_start, a_end): Range, (b_start, b_end): Range) -> bool {
    a_start <= b_end && b_start <= a_end
}
