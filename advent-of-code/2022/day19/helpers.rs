pub struct Blueprint {
    pub id: i64,
    pub ore_cost: i64,
    pub clay_cost: i64,
    pub obsidian_cost: (i64, i64),
    pub geode_cost: (i64, i64),
    pub max_ore_cost: i64,
    pub max_clay_cost: i64,
    pub max_obsidian_cost: i64,
}

impl Blueprint {
    fn new(
        id: i64,
        ore_cost: i64,
        clay_cost: i64,
        obsidian_cost: (i64, i64),
        geode_cost: (i64, i64),
    ) -> Self {
        let max_ore_cost = ore_cost
            .max(clay_cost)
            .max(obsidian_cost.0)
            .max(geode_cost.0);
        Blueprint {
            id,
            ore_cost,
            clay_cost,
            obsidian_cost,
            geode_cost,
            max_ore_cost,
            max_clay_cost: obsidian_cost.1,
            max_obsidian_cost: geode_cost.1,
        }
    }
}

pub fn parse_blueprints(puzzle_input: &[String]) -> Vec<Blueprint> {
    puzzle_input
        .iter()
        .filter(|line| !line.is_empty())
        .map(|line| {
            let nums: Vec<i64> = line
                .split(|c: char| !c.is_ascii_digit())
                .filter(|s| !s.is_empty())
                .map(|s| s.parse().unwrap())
                .collect();
            Blueprint::new(
                nums[0],
                nums[1],
                nums[2],
                (nums[3], nums[4]),
                (nums[5], nums[6]),
            )
        })
        .collect()
}

#[derive(Clone, Copy)]
struct State {
    time_left: i64,
    ore: i64,
    clay: i64,
    obsidian: i64,
    geode: i64,
    ore_robots: i64,
    clay_robots: i64,
    obsidian_robots: i64,
    geode_robots: i64,
}

fn time_to_afford_one(have: i64, rate: i64, cost: i64) -> Option<i64> {
    if have >= cost {
        Some(0)
    } else if rate == 0 {
        None
    } else {
        Some((cost - have + rate - 1) / rate)
    }
}

fn time_to_afford(
    ore: i64,
    ore_rate: i64,
    ore_cost: i64,
    other: i64,
    other_rate: i64,
    other_cost: i64,
) -> Option<i64> {
    let ore_wait = time_to_afford_one(ore, ore_rate, ore_cost)?;
    let other_wait = time_to_afford_one(other, other_rate, other_cost)?;
    Some(ore_wait.max(other_wait))
}

fn dfs(bp: &Blueprint, s: State, best: &mut i64) {
    let projected = s.geode + s.geode_robots * s.time_left + s.time_left * (s.time_left - 1) / 2;
    if projected <= *best {
        return;
    }
    *best = (*best).max(s.geode + s.geode_robots * s.time_left);

    if s.ore_robots > 0
        && s.obsidian_robots > 0
        && let Some(wait) = time_to_afford(
            s.ore,
            s.ore_robots,
            bp.geode_cost.0,
            s.obsidian,
            s.obsidian_robots,
            bp.geode_cost.1,
        )
        && s.time_left - wait - 1 > 0
    {
        let elapsed = wait + 1;
        dfs(
            bp,
            State {
                time_left: s.time_left - elapsed,
                ore: s.ore + s.ore_robots * elapsed - bp.geode_cost.0,
                clay: s.clay + s.clay_robots * elapsed,
                obsidian: s.obsidian + s.obsidian_robots * elapsed - bp.geode_cost.1,
                geode: s.geode + s.geode_robots * elapsed,
                geode_robots: s.geode_robots + 1,
                ..s
            },
            best,
        );
    }

    if s.obsidian_robots < bp.max_obsidian_cost
        && s.clay_robots > 0
        && s.ore_robots > 0
        && let Some(wait) = time_to_afford(
            s.ore,
            s.ore_robots,
            bp.obsidian_cost.0,
            s.clay,
            s.clay_robots,
            bp.obsidian_cost.1,
        )
        && s.time_left - wait - 1 > 0
    {
        let elapsed = wait + 1;
        dfs(
            bp,
            State {
                time_left: s.time_left - elapsed,
                ore: s.ore + s.ore_robots * elapsed - bp.obsidian_cost.0,
                clay: s.clay + s.clay_robots * elapsed - bp.obsidian_cost.1,
                obsidian: s.obsidian + s.obsidian_robots * elapsed,
                geode: s.geode + s.geode_robots * elapsed,
                obsidian_robots: s.obsidian_robots + 1,
                ..s
            },
            best,
        );
    }

    if s.clay_robots < bp.max_clay_cost
        && s.ore_robots > 0
        && let Some(wait) = time_to_afford_one(s.ore, s.ore_robots, bp.clay_cost)
        && s.time_left - wait - 1 > 0
    {
        let elapsed = wait + 1;
        dfs(
            bp,
            State {
                time_left: s.time_left - elapsed,
                ore: s.ore + s.ore_robots * elapsed - bp.clay_cost,
                clay: s.clay + s.clay_robots * elapsed,
                obsidian: s.obsidian + s.obsidian_robots * elapsed,
                geode: s.geode + s.geode_robots * elapsed,
                clay_robots: s.clay_robots + 1,
                ..s
            },
            best,
        );
    }

    if s.ore_robots < bp.max_ore_cost
        && let Some(wait) = time_to_afford_one(s.ore, s.ore_robots, bp.ore_cost)
        && s.time_left - wait - 1 > 0
    {
        let elapsed = wait + 1;
        dfs(
            bp,
            State {
                time_left: s.time_left - elapsed,
                ore: s.ore + s.ore_robots * elapsed - bp.ore_cost,
                clay: s.clay + s.clay_robots * elapsed,
                obsidian: s.obsidian + s.obsidian_robots * elapsed,
                geode: s.geode + s.geode_robots * elapsed,
                ore_robots: s.ore_robots + 1,
                ..s
            },
            best,
        );
    }
}

pub fn max_geodes(bp: &Blueprint, time_limit: i64) -> i64 {
    let mut best = 0;
    let state = State {
        time_left: time_limit,
        ore: 0,
        clay: 0,
        obsidian: 0,
        geode: 0,
        ore_robots: 1,
        clay_robots: 0,
        obsidian_robots: 0,
        geode_robots: 0,
    };
    dfs(bp, state, &mut best);
    best
}
