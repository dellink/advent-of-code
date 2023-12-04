use regex::Regex;
use std::collections::{HashMap, HashSet};

use aoc::read_file_input;

fn main() {
    let input = read_file_input("04.txt".to_string());

    let mut part1 = 0;

    for line in input.lines() {
        let intersection = get_intersection(line);

        if intersection.len() > 0 {
            part1 += 2_i32.pow((intersection.len() - 1).try_into().unwrap());
        }
    }
    println!("{}", part1);

    let mut part2 = 0;
    let mut copies: HashMap<usize, usize> = HashMap::new();

    for (index, line) in input.lines().enumerate() {
        let intersection = get_intersection(line);

        let game_index = index + 1;
        let copy_count = copies.get(&game_index).unwrap_or(&0) + 1;

        part2 += copy_count;

        if intersection.len() > 0 {
            for num in game_index + 1..game_index + intersection.len() + 1 {
                copies.insert(num, copies.get(&num).unwrap_or(&0) + copy_count);
            }
        }
    }

    println!("{}", part2);
}

fn get_intersection(line: &str) -> Vec<i32> {
    let re = Regex::new(r"\d+").unwrap();

    let parts: Vec<&str> = line.splitn(2, ':').collect();

    if parts.len() == 2 {
        let halves: Vec<&str> = parts[1].splitn(2, '|').collect();
        if halves.len() == 2 {
            let numbers_before: HashSet<_> = re
                .find_iter(halves[0])
                .map(|mat| mat.as_str().parse::<i32>().unwrap())
                .collect();
            let numbers_after: HashSet<_> = re
                .find_iter(halves[1])
                .map(|mat| mat.as_str().parse::<i32>().unwrap())
                .collect();

            return numbers_before
                .intersection(&numbers_after)
                .cloned()
                .collect();
        }
    }

    return vec![];
}
