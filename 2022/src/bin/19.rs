use std::collections::HashSet;

use aoc::read_file_input;

fn main() {
    let input = read_file_input("19.txt".to_string());
    let blueprints = parse(input);

    let mut count = 0;

    for (i, bp) in blueprints.iter().enumerate() {
        count += max_geodes(&bp, 24) * (i + 1) as u16
    }

    println!("{}", count);

    let mut count = 1;

    for bp in blueprints.iter().take(3) {
        count *= max_geodes(&bp, 32)
    }

    println!("{}", count);
}

#[derive(Debug)]
struct Blueprint {
    ore: u16,
    clay: u16,
    obsidian: (u16, u16),
    geode: (u16, u16),
}

fn parse(input: String) -> Vec<Blueprint> {
    let mut blueprints: Vec<Blueprint> = vec![];

    for line in input.lines() {
        let mut blueprint = Blueprint {
            ore: 0,
            clay: 0,
            obsidian: (0, 0),
            geode: (0, 0),
        };

        let costs: Vec<&str> = line.split("costs ").skip(1).collect();
        for (i, cost) in costs.iter().enumerate() {
            if i == 0 {
                blueprint.ore = cost.split_once(" ").unwrap().0.parse::<u16>().unwrap();
            }
            if i == 1 {
                blueprint.clay = cost.split_once(" ").unwrap().0.parse::<u16>().unwrap();
            }
            if i == 2 {
                let c = cost.split_once("and ").unwrap();
                blueprint.obsidian = (
                    c.0.split_once(" ").unwrap().0.parse::<u16>().unwrap(),
                    c.1.split_once(" ").unwrap().0.parse::<u16>().unwrap(),
                );
            }
            if i == 3 {
                let c = cost.split_once("and ").unwrap();
                blueprint.geode = (
                    c.0.split_once(" ").unwrap().0.parse::<u16>().unwrap(),
                    c.1.split_once(" ").unwrap().0.parse::<u16>().unwrap(),
                );
            }
        }
        blueprints.push(blueprint);
    }

    return blueprints;
}

#[derive(Default, Clone, Copy, Hash, PartialEq, Eq)]
struct State {
    ore: u16,
    ore_robots: u16,
    clay: u16,
    clay_robots: u16,
    obsidian: u16,
    obsidian_robots: u16,
    geode: u16,
    geode_robots: u16,
}

fn collect(state: State) -> State {
    State {
        ore: state.ore + state.ore_robots,
        clay: state.clay + state.clay_robots,
        obsidian: state.obsidian + state.obsidian_robots,
        geode: state.geode + state.geode_robots,
        ..state
    }
}

fn max_geodes(blueprint: &Blueprint, time: u16) -> u16 {
    let max_ore_cost = *[
        blueprint.ore,
        blueprint.clay,
        blueprint.obsidian.0,
        blueprint.geode.0,
    ]
    .iter()
    .max()
    .unwrap();

    let mut seen = HashSet::with_capacity(10000);

    let mut states = vec![State {
        ore_robots: 1,
        ..State::default()
    }];

    let mut best_geode_robots = 0;

    for _ in 0..time {
        let mut new_states = vec![];

        for &state in &states {
            if state.geode_robots < best_geode_robots {
                continue;
            }

            if !seen.insert(state) {
                continue;
            }

            best_geode_robots = best_geode_robots.max(state.geode_robots);

            if state.ore >= blueprint.geode.0 && state.obsidian >= blueprint.geode.1 {
                let mut state = collect(state);

                state.ore -= blueprint.geode.0;
                state.obsidian -= blueprint.geode.1;
                state.geode_robots += 1;

                new_states.push(state);
            }

            if state.obsidian_robots < blueprint.geode.1
                && state.ore >= blueprint.obsidian.0
                && state.clay >= blueprint.obsidian.1
            {
                let mut state = collect(state);

                state.ore -= blueprint.obsidian.0;
                state.clay -= blueprint.obsidian.1;
                state.obsidian_robots += 1;

                new_states.push(state);
            }

            if state.ore_robots < max_ore_cost && state.ore >= blueprint.ore {
                let mut state = collect(state);

                state.ore -= blueprint.ore;
                state.ore_robots += 1;

                new_states.push(state);
            }

            if state.clay_robots < blueprint.obsidian.1 && state.ore >= blueprint.clay {
                let mut state = collect(state);

                state.ore -= blueprint.clay;
                state.clay_robots += 1;

                new_states.push(state);
            }

            new_states.push(collect(state));
        }
        states = new_states;
    }

    states.iter().map(|state| state.geode).max().unwrap()
}
