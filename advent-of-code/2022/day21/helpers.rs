use std::collections::HashMap;

pub enum Expr {
    Num(i64),
    Op(String, char, String),
}

pub fn parse_monkeys(puzzle_input: &[String]) -> HashMap<String, Expr> {
    puzzle_input
        .iter()
        .filter(|line| !line.is_empty())
        .map(|line| {
            let (name, rest) = line.split_once(": ").unwrap();
            let expr = if let Ok(n) = rest.parse::<i64>() {
                Expr::Num(n)
            } else {
                let parts: Vec<&str> = rest.split_whitespace().collect();
                Expr::Op(
                    parts[0].to_string(),
                    parts[1].chars().next().unwrap(),
                    parts[2].to_string(),
                )
            };
            (name.to_string(), expr)
        })
        .collect()
}

pub fn eval(monkeys: &HashMap<String, Expr>, name: &str) -> i64 {
    match &monkeys[name] {
        Expr::Num(n) => *n,
        Expr::Op(l, op, r) => {
            let lv = eval(monkeys, l);
            let rv = eval(monkeys, r);
            match op {
                '+' => lv + rv,
                '-' => lv - rv,
                '*' => lv * rv,
                '/' => lv / rv,
                _ => panic!("unknown operator: {op}"),
            }
        }
    }
}

pub fn children(monkeys: &HashMap<String, Expr>, name: &str) -> (String, String) {
    match &monkeys[name] {
        Expr::Op(l, _, r) => (l.clone(), r.clone()),
        Expr::Num(_) => panic!("{name} has no children"),
    }
}

pub fn depends_on_humn(monkeys: &HashMap<String, Expr>, name: &str) -> bool {
    if name == "humn" {
        return true;
    }
    match &monkeys[name] {
        Expr::Num(_) => false,
        Expr::Op(l, _, r) => depends_on_humn(monkeys, l) || depends_on_humn(monkeys, r),
    }
}

pub fn solve(monkeys: &HashMap<String, Expr>, name: &str, target: i64) -> i64 {
    if name == "humn" {
        return target;
    }
    let Expr::Op(l, op, r) = &monkeys[name] else {
        panic!("{name} is a leaf, cannot solve")
    };

    if depends_on_humn(monkeys, l) {
        let rv = eval(monkeys, r);
        let new_target = match op {
            '+' => target - rv,
            '-' => target + rv,
            '*' => target / rv,
            '/' => target * rv,
            _ => panic!("unknown operator: {op}"),
        };
        solve(monkeys, l, new_target)
    } else {
        let lv = eval(monkeys, l);
        let new_target = match op {
            '+' => target - lv,
            '-' => lv - target,
            '*' => target / lv,
            '/' => lv / target,
            _ => panic!("unknown operator: {op}"),
        };
        solve(monkeys, r, new_target)
    }
}
