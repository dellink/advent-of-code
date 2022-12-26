use std::collections::{HashMap, HashSet};

use aoc::read_file_input;

fn main() {
    let input = read_file_input("16.txt".to_string());
    let map = parse(input);
    let paths = get_path(&map);

    println!(
        "{}",
        best(&map, &paths, String::from("AA"), HashSet::new(), 30)
    )
}

#[derive(Debug)]
struct Valve {
    rate: isize,
    tunnels: Vec<String>,
}

fn parse(input: String) -> HashMap<String, Valve> {
    let mut scan: HashMap<String, Valve> = HashMap::new();

    for line in input.lines() {
        let (valve, tunnels) = match line.split_once("; tunnels lead to valves ") {
            Some(v) => (v.0, v.1),
            None => line.split_once("; tunnel leads to valve ").unwrap(),
        };

        let (name, rate) = valve.split_once(" has flow rate=").unwrap();

        scan.insert(
            name.split_once("Valve ").unwrap().1.to_string(),
            Valve {
                rate: rate.parse().unwrap(),
                tunnels: tunnels.split(", ").map(|x| x.clone().to_string()).collect(),
            },
        );
    }

    scan
}

fn distance(valves: &HashMap<String, Valve>, from: String, to: String) -> isize {
    let mut queue = vec![(&from, 0)];
    let mut visited = HashSet::from([&from]);

    while queue.len() > 0 {
        let current = queue.remove(0);
        if current.0.to_string() == to.to_string() {
            return current.1;
        }
        for valve in valves[current.0].tunnels.iter() {
            if !visited.contains(valve) {
                queue.push((valve, current.1 + 1));
                visited.insert(valve);
            }
        }
    }

    unreachable!()
}

fn get_path(map: &HashMap<String, Valve>) -> HashMap<String, HashMap<String, isize>> {
    let mut path: HashMap<String, HashMap<String, isize>> = HashMap::new();

    for from in map.keys() {
        let mut hash = HashMap::new();

        for to in map.keys() {
            if from != to && map[to].rate > 0 {
                let d = distance(map, from.to_string(), to.to_string());
                hash.insert(to.to_string(), d);
            }
        }

        path.insert(from.to_string(), hash);
    }

    path
}

fn best(
    valves: &HashMap<String, Valve>,
    paths: &HashMap<String, HashMap<String, isize>>,
    current: String,
    opened: HashSet<String>,
    time: isize,
) -> isize {
    if time == 0 {
        return 0;
    }
    let mut max: isize = 0;

    for to in paths[&current].keys() {
        if !opened.contains(to) && time >= paths[&current][to] as isize {
            let remaining: isize = time as isize - paths[&current][to] as isize - 1;

            let pressure = valves[to].rate * remaining;
            let mut set = HashSet::from(opened.clone());
            set.insert(to.to_string());

            let max_value = pressure + best(valves, paths, to.to_string(), set, remaining);

            max = max.max(max_value);
        }
    }

    return max;
}
