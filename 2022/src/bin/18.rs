use std::collections::HashSet;

use aoc::read_file_input;
use itertools::Itertools;

fn main() {
    let input = read_file_input("18.txt".to_string());
    let mut cubes = HashSet::new();

    for line in input.lines() {
        if let Some(cube) = line
            .split(',')
            .map(|x| x.parse::<isize>().unwrap())
            .collect_tuple::<(isize, isize, isize)>()
        {
            cubes.insert(cube);
        }
    }

    let mut res = 0;

    for cube in cubes.clone().into_iter() {
        res += count(cube, &cubes);
    }

    println!("{}", res);

    let max = cubes
        .clone()
        .into_iter()
        .flat_map(|(x, y, z)| [x, y, z])
        .max()
        .unwrap();

    let min = cubes
        .clone()
        .into_iter()
        .flat_map(|(x, y, z)| [x, y, z])
        .min()
        .unwrap();

    let mut res = 0;

    let mut visited = HashSet::new();
    let mut queue = vec![(0, 0, 0)];

    while let Some((x, y, z)) = queue.pop() {
        if visited.contains(&(x, y, z)) || cubes.contains(&(x, y, z)) {
            continue;
        }

        if x < min - 1 || y < min - 1 || z < min - 1 || x > max + 1 || y > max + 1 || z > max + 1 {
            continue;
        }

        visited.insert((x, y, z));

        res += count_exposed((x, y, z), &cubes);

        queue.push((x + 1, y, z));
        queue.push((x - 1, y, z));
        queue.push((x, y + 1, z));
        queue.push((x, y - 1, z));
        queue.push((x, y, z + 1));
        queue.push((x, y, z - 1));
    }

    println!("{}", res);
}

fn count(cube: (isize, isize, isize), cubes: &HashSet<(isize, isize, isize)>) -> isize {
    let mut count = 0;

    if !cubes.contains(&(cube.0 + 1, cube.1, cube.2)) {
        count += 1;
    }
    if !cubes.contains(&(cube.0 - 1, cube.1, cube.2)) {
        count += 1;
    }
    if !cubes.contains(&(cube.0, cube.1 + 1, cube.2)) {
        count += 1;
    }
    if !cubes.contains(&(cube.0, cube.1 - 1, cube.2)) {
        count += 1;
    }
    if !cubes.contains(&(cube.0, cube.1, cube.2 + 1)) {
        count += 1;
    }
    if !cubes.contains(&(cube.0, cube.1, cube.2 - 1)) {
        count += 1;
    }

    return count;
}

fn count_exposed(cube: (isize, isize, isize), cubes: &HashSet<(isize, isize, isize)>) -> isize {
    let mut count = 0;

    if cubes.contains(&(cube.0 + 1, cube.1, cube.2)) {
        count += 1;
    }
    if cubes.contains(&(cube.0 - 1, cube.1, cube.2)) {
        count += 1;
    }
    if cubes.contains(&(cube.0, cube.1 + 1, cube.2)) {
        count += 1;
    }
    if cubes.contains(&(cube.0, cube.1 - 1, cube.2)) {
        count += 1;
    }
    if cubes.contains(&(cube.0, cube.1, cube.2 + 1)) {
        count += 1;
    }
    if cubes.contains(&(cube.0, cube.1, cube.2 - 1)) {
        count += 1;
    }

    return count;
}
