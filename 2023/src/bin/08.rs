use num::integer::lcm;
use std::collections::HashMap;

use aoc::read_file_input;

fn main() {
    let input = read_file_input("08.txt".to_string());

    let mut sections = input.split("\n\n");

    let instructions: Vec<_> = sections.next().unwrap().chars().collect();

    let mut map = HashMap::new();

    for line in sections.next().unwrap().lines() {
        if line.trim().is_empty() {
            continue;
        }

        let parts: Vec<&str> = line.split('=').collect();

        let key = parts[0].trim();
        let value = parts[1]
            .trim()
            .trim_start_matches('(')
            .trim_end_matches(')');

        map.insert(key, value.split(',').map(|s| s.trim()).collect::<Vec<_>>());
    }

    let mut start = "AAA";
    let mut steps = 0;

    while start != "ZZZ" {
        start = get_next(start, &map, &instructions, steps);

        steps += 1;
    }

    println!("{}", steps);

    let mut starts: Vec<&str> = map
        .keys()
        .filter(|&key| key.ends_with('A'))
        .cloned()
        .collect();

    let mut step = 0;
    let mut steps = vec![];

    loop {
        for start in &mut starts {
            if start.ends_with('Z') {
                continue;
            }

            *start = get_next(start, &map, &instructions, step);

            if start.ends_with('Z') {
                steps.push(step + 1);
            }
        }

        step += 1;

        if steps.len() == starts.len() {
            break;
        }
    }

    let lcm = steps.into_iter().reduce(lcm);

    println!("{}", lcm.unwrap());
}

fn get_next<'a>(
    start: &'a str,
    map: &'a HashMap<&'a str, Vec<&'a str>>,
    instructions: &'a Vec<char>,
    steps: usize,
) -> &'a str {
    let next = map.get(start).unwrap();

    let mut i = steps;

    if steps >= instructions.len() {
        i = steps % instructions.len();
    }

    if instructions[i] == 'L' {
        next[0]
    } else {
        next[1]
    }
}
