use std::collections::HashMap;

use aoc::read_file_input;

fn main() {
    let input = read_file_input("16.txt".to_string());
    let scan = parse(input);

    println!("{:?}", scan);
}

#[derive(Debug)]
struct Output {
    rate: u8,
    tunnels: Vec<String>,
}

fn parse(input: String) -> HashMap<String, Output> {
    let mut scan: HashMap<String, Output> = HashMap::new();

    for line in input.lines() {
        let (valve, tunnels) = match line.split_once("; tunnels lead to valves ") {
            Some(v) => (v.0, v.1),
            None => line.split_once("; tunnel leads to valve ").unwrap(),
        };

        let (name, rate) = valve.split_once(" has flow rate=").unwrap();

        scan.insert(
            name.split_once("Valve ").unwrap().1.to_string(),
            Output {
                rate: rate.parse().unwrap(),
                tunnels: tunnels.split(", ").map(|x| x.clone().to_string()).collect(),
            },
        );
    }

    scan
}
