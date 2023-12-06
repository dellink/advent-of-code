use regex::Regex;
use std::{i64::MAX, str::FromStr};

use aoc::read_file_input;

type Range = (i64, i64, i64);

fn main() {
    let input = read_file_input("05.txt".to_string());

    let mut sections = input.split("\n\n");
    let mut seeds = parse_seeds(sections.next().unwrap());
    let seeds2 = seeds.clone();
    let maps = sections
        .map(|s| parse_maps(s))
        .collect::<Vec<Vec<Vec<i64>>>>();
    let ranges = maps.iter().map(|m| create_ranges(m)).collect::<Vec<_>>();

    for range in &ranges {
        for seed in &mut seeds {
            *seed = map_value(*seed, range);
        }
    }

    println!("{}", seeds.iter().min().unwrap());

    let mut min = MAX;

    for seed in seeds2.chunks(2).into_iter() {
        for s in seed[0]..seed[0] + seed[1] {
            let mut output = s;
            for range in &ranges {
                output = map_value(output, range);
            }
            min = min.min(output);
        }
    }

    println!("{}", min);
}

fn parse_seeds(seeds: &str) -> Vec<i64> {
    seeds
        .split_whitespace()
        .filter_map(|s| i64::from_str(s).ok())
        .collect()
}

fn parse_maps(maps: &str) -> Vec<Vec<i64>> {
    let re = Regex::new(r"(\d+) (\d+) (\d+)").unwrap();
    re.captures_iter(maps)
        .filter_map(|cap| {
            let nums = (1..=3)
                .filter_map(|i| cap.get(i).and_then(|m| i64::from_str(m.as_str()).ok()))
                .collect::<Vec<_>>();
            if nums.len() == 3 {
                Some(nums)
            } else {
                None
            }
        })
        .collect()
}

fn create_ranges(data: &[Vec<i64>]) -> Vec<Range> {
    data.iter().map(|r| (r[0], r[1], r[2])).collect()
}

fn map_value(value: i64, ranges: &[Range]) -> i64 {
    for &(dest_start, src_start, length) in ranges {
        if value >= src_start && value < src_start + length {
            return dest_start + (value - src_start);
        }
    }
    value
}
