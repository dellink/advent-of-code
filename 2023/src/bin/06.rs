use std::str::FromStr;

use aoc::read_file_input;

fn main() {
    let input = read_file_input("06.txt".to_string());
    let (times, distances) = parse_input1(&input);

    let mut part1 = 1;

    for (index, time) in times.into_iter().enumerate() {
        let mut ways = 0;

        for i in 0..time {
            let max = i * (time - i);

            if max > distances[index] {
                ways += 1;
            }
        }

        part1 *= ways;
    }

    println!("{}", part1);

    let (time, distance) = parse_input2(&input);

    let mut part2 = 0;

    for i in 0..time {
        let max = i * (time - i);

        if max > distance {
            part2 += 1;
        }
    }

    println!("{}", part2);
}

fn parse_input1(input: &str) -> (Vec<i64>, Vec<i64>) {
    let mut lines = input.lines();
    let times = lines
        .next()
        .unwrap()
        .split_whitespace()
        .skip(1)
        .map(|s| i64::from_str(s).unwrap())
        .collect();
    let distances = lines
        .next()
        .unwrap()
        .split_whitespace()
        .skip(1)
        .map(|s| i64::from_str(s).unwrap())
        .collect();
    (times, distances)
}

fn parse_input2(input: &str) -> (i64, i64) {
    let mut lines = input.lines();
    let time = lines
        .next()
        .unwrap()
        .split_whitespace()
        .skip(1)
        .collect::<String>();
    let distance = lines
        .next()
        .unwrap()
        .split_whitespace()
        .skip(1)
        .collect::<String>();
    (
        i64::from_str(&time).unwrap(),
        i64::from_str(&distance).unwrap(),
    )
}
