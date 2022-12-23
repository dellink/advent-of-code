use std::collections::HashSet;

use aoc::read_file_input;

fn main() {
    let input = read_file_input("15.txt".to_string()).to_string();
    let map = parse(input.clone());

    println!("{}", part_one(&map));
    println!("{}", part_two(&map));
}

#[derive(Debug, Clone)]
struct Unit {
    sensor: (i64, i64),
    beacon: (i64, i64),
    distance: i64,
}

fn parse(input: String) -> Vec<Unit> {
    let mut map: Vec<Unit> = vec![];

    for line in input.lines() {
        let (sensor, beacon) = line.split_once(": ").unwrap();

        let sensor = sensor
            .trim_start_matches("Sensor at ")
            .split_once(", ")
            .unwrap();

        let sensor_x: i64 = sensor.0.split_once("x=").unwrap().1.parse().unwrap();
        let sensor_y: i64 = sensor.1.split_once("y=").unwrap().1.parse().unwrap();

        let beacon = beacon
            .trim_start_matches("closest beacon is at ")
            .split_once(", ")
            .unwrap();

        let beacon_x: i64 = beacon.0.split_once("x=").unwrap().1.parse().unwrap();
        let beacon_y: i64 = beacon.1.split_once("y=").unwrap().1.parse().unwrap();

        map.push(Unit {
            sensor: (sensor_x, sensor_y),
            beacon: (beacon_x, beacon_y),
            distance: distance((sensor_x, sensor_y), (beacon_x, beacon_y)),
        });
    }

    map
}

fn distance(a: (i64, i64), b: (i64, i64)) -> i64 {
    (a.0 - b.0).abs() + (a.1 - b.1).abs()
}

fn part_one(map: &Vec<Unit>) -> usize {
    let mut cannot_contain_beacon = HashSet::new();
    let mut beacons_on_line = HashSet::new();
    let y: i64 = 2000000;

    for line in map.into_iter() {
        if line.beacon.1 == y {
            beacons_on_line.insert(line.beacon.0);
        }
        let min = distance(line.sensor, (line.sensor.0, y));
        if min <= line.distance {
            let distance_around_sensor = line.distance - min;

            for i in
                (line.sensor.0 - distance_around_sensor)..=(line.sensor.0 + distance_around_sensor)
            {
                cannot_contain_beacon.insert(i);
            }
        }
    }

    return cannot_contain_beacon.len() - beacons_on_line.len();
}

fn part_two(map: &Vec<Unit>) -> usize {
    let max_coordinate: i64 = 4000000;

    for y in 0..max_coordinate {
        let mut intervals = vec![];

        for line in map {
            let min_distance = distance(line.sensor, (line.sensor.0, y));
            if min_distance <= line.distance {
                let distance_around_sensor_x = line.distance - min_distance;

                let from = line.sensor.0 - distance_around_sensor_x;
                let to = line.sensor.0 + distance_around_sensor_x;

                intervals.push((from, to));
            }
        }

        intervals.sort_by(|a, b| a.0.cmp(&b.0));

        let mut merged_intervals: Vec<(i64, i64)> = vec![];

        for interval in intervals {
            let len = merged_intervals.len();

            if len > 0 && merged_intervals[len - 1].1 >= interval.0 {
                merged_intervals[len - 1] = (
                    merged_intervals[len - 1].0,
                    interval.1.max(merged_intervals[len - 1].1),
                );
            } else {
                merged_intervals.push(interval);
            }
        }

        let mut res = vec![];

        for interval in merged_intervals {
            if interval.0 > max_coordinate || 0 >= interval.1 {
                continue;
            }
            res.push((interval.0.max(0), interval.1.min(max_coordinate)))
        }

        if res.len() > 1 {
            let x = res[0].1 + 1;
            return (x * 4000000 + y).try_into().unwrap();
        }
    }

    unreachable!();
}
