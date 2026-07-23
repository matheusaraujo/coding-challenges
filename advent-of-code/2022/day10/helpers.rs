use std::collections::HashMap;

fn ocr_alphabet() -> HashMap<&'static str, char> {
    HashMap::from([
        (".##.\n#..#\n#..#\n####\n#..#\n#..#", 'A'),
        ("###.\n#..#\n###.\n#..#\n#..#\n###.", 'B'),
        (".##.\n#..#\n#...\n#...\n#..#\n.##.", 'C'),
        ("####\n#...\n###.\n#...\n#...\n####", 'E'),
        ("####\n#...\n###.\n#...\n#...\n#...", 'F'),
        (".##.\n#..#\n#...\n#.##\n#..#\n.###", 'G'),
        ("#..#\n#..#\n####\n#..#\n#..#\n#..#", 'H'),
        (".###\n..#.\n..#.\n..#.\n..#.\n.###", 'I'),
        ("..##\n...#\n...#\n...#\n#..#\n.##.", 'J'),
        ("#..#\n#.#.\n##..\n#.#.\n#.#.\n#..#", 'K'),
        ("#...\n#...\n#...\n#...\n#...\n####", 'L'),
        (".##.\n#..#\n#..#\n#..#\n#..#\n.##.", 'O'),
        ("###.\n#..#\n#..#\n###.\n#...\n#...", 'P'),
        ("###.\n#..#\n#..#\n###.\n#.#.\n#..#", 'R'),
        (".###\n#...\n#...\n.##.\n...#\n###.", 'S'),
        ("####\n..#.\n..#.\n..#.\n..#.\n..#.", 'T'),
        ("#..#\n#..#\n#..#\n#..#\n#..#\n.##.", 'U'),
        ("####\n...#\n..#.\n.#..\n#...\n####", 'Z'),
    ])
}

pub fn parse_ocr(rows: &[String]) -> String {
    let alphabet = ocr_alphabet();
    let width = rows[0].len();

    let mut result = String::new();
    let mut col = 0;
    while col < width {
        let glyph = rows
            .iter()
            .map(|row| row.get(col..col + 4).unwrap_or(""))
            .collect::<Vec<_>>()
            .join("\n");
        if glyph.contains('#') {
            result.push(*alphabet.get(glyph.as_str()).unwrap_or(&'?'));
        }
        col += 5;
    }

    result
}

pub fn x_during_cycles(puzzle_input: &[String]) -> Vec<i32> {
    let mut x = 1;
    let mut history = Vec::new();

    for line in puzzle_input.iter().filter(|line| !line.is_empty()) {
        if line == "noop" {
            history.push(x);
        } else if let Some(operand) = line.strip_prefix("addx ") {
            let value: i32 = operand.parse().expect("addx operand");
            history.push(x);
            history.push(x);
            x += value;
        }
    }

    history
}
