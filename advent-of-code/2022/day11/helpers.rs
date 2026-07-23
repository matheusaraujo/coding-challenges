use std::collections::VecDeque;

#[derive(Clone)]
pub enum Operation {
    Add(u64),
    Mul(u64),
    Square,
}

impl Operation {
    pub fn apply(&self, old: u64) -> u64 {
        match self {
            Operation::Add(n) => old + n,
            Operation::Mul(n) => old * n,
            Operation::Square => old * old,
        }
    }
}

#[derive(Clone)]
pub struct Monkey {
    pub items: VecDeque<u64>,
    pub operation: Operation,
    pub divisor: u64,
    pub if_true: usize,
    pub if_false: usize,
}

pub fn parse_monkeys(puzzle_input: &[String]) -> Vec<Monkey> {
    puzzle_input
        .split(|line| line.is_empty())
        .filter(|block| !block.is_empty())
        .map(|block| {
            let items = block[1]
                .split(':')
                .nth(1)
                .unwrap()
                .split(',')
                .map(|n| n.trim().parse().unwrap())
                .collect();

            let op_parts: Vec<&str> = block[2].split_whitespace().collect();
            let operator = op_parts[op_parts.len() - 2];
            let operand = op_parts[op_parts.len() - 1];
            let operation = if operand == "old" {
                Operation::Square
            } else {
                let n: u64 = operand.parse().unwrap();
                if operator == "*" {
                    Operation::Mul(n)
                } else {
                    Operation::Add(n)
                }
            };

            let divisor = block[3].split_whitespace().last().unwrap().parse().unwrap();
            let if_true = block[4].split_whitespace().last().unwrap().parse().unwrap();
            let if_false = block[5].split_whitespace().last().unwrap().parse().unwrap();

            Monkey {
                items,
                operation,
                divisor,
                if_true,
                if_false,
            }
        })
        .collect()
}

pub fn monkey_business(monkeys: &mut [Monkey], rounds: usize, relief: impl Fn(u64) -> u64) -> u64 {
    let mut inspections = vec![0u64; monkeys.len()];

    for _ in 0..rounds {
        for i in 0..monkeys.len() {
            let items: Vec<u64> = monkeys[i].items.drain(..).collect();
            for item in items {
                inspections[i] += 1;
                let worry = relief(monkeys[i].operation.apply(item));
                let target = if worry.is_multiple_of(monkeys[i].divisor) {
                    monkeys[i].if_true
                } else {
                    monkeys[i].if_false
                };
                monkeys[target].items.push_back(worry);
            }
        }
    }

    inspections.sort_unstable_by(|a, b| b.cmp(a));
    inspections[0] * inspections[1]
}
