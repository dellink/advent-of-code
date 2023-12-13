use regex::Regex;

use aoc::read_file_input;

fn main() {
    let input = read_file_input("09.txt".to_string());
    let re = Regex::new(r"(-?\d+)").unwrap();

    let mut part1 = 0;
    let mut part2 = 0;

    for line in input.lines() {
        let numbers: Vec<_> = re
            .captures_iter(line)
            .map(|cap| cap[1].parse::<i32>().unwrap())
            .collect();

        let mut datasets: Vec<Vec<i32>> = Vec::new();
        datasets.push(numbers);

        while !datasets.last().unwrap().iter().all(|&num| num == 0) {
            let last = datasets.last().unwrap();
            let mut new = Vec::new();

            for i in 1..last.len() {
                new.push(last[i] - last[i - 1]);
            }

            datasets.push(new);
        }

        let mut prev_right = 0;
        let mut prev_left = 0;

        for dataset in datasets.into_iter().rev() {
            prev_right += dataset.last().unwrap();
            prev_left = dataset.first().unwrap() - prev_left;
        }

        part1 += prev_right;
        part2 += prev_left;
    }

    println!("{:?}", part1);
    println!("{:?}", part2);
}
