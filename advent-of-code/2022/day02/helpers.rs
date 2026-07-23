#[derive(Clone, Copy, PartialEq)]
pub enum Shape {
    Rock,
    Paper,
    Scissors,
}

impl Shape {
    pub fn score(self) -> u32 {
        match self {
            Shape::Rock => 1,
            Shape::Paper => 2,
            Shape::Scissors => 3,
        }
    }

    pub fn beats(self) -> Shape {
        match self {
            Shape::Rock => Shape::Scissors,
            Shape::Paper => Shape::Rock,
            Shape::Scissors => Shape::Paper,
        }
    }

    pub fn beaten_by(self) -> Shape {
        match self {
            Shape::Rock => Shape::Paper,
            Shape::Paper => Shape::Scissors,
            Shape::Scissors => Shape::Rock,
        }
    }

    pub fn from_opponent(c: char) -> Shape {
        match c {
            'A' => Shape::Rock,
            'B' => Shape::Paper,
            'C' => Shape::Scissors,
            _ => panic!("invalid opponent shape: {c}"),
        }
    }
}

pub fn outcome_score(opponent: Shape, you: Shape) -> u32 {
    if you == opponent {
        3
    } else if you.beats() == opponent {
        6
    } else {
        0
    }
}

pub fn round_columns(line: &str) -> (char, char) {
    let mut columns = line.split_whitespace();
    let opponent = columns.next().unwrap().chars().next().unwrap();
    let you = columns.next().unwrap().chars().next().unwrap();
    (opponent, you)
}
