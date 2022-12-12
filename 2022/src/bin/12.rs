use std::collections::VecDeque;

use aoc::read_file_input;

fn main() {
    let input = read_file_input("12.txt".to_string()).to_string();

    let mut map: Vec<Vec<u8>> = Vec::new();
    let mut start: (usize, usize) = (0, 0);
    let mut end: (usize, usize) = (0, 0);

    for (i, line) in input.lines().enumerate() {
        let mut row: Vec<u8> = Vec::new();

        for (j, col) in line.chars().enumerate() {
            match col {
                'S' => {
                    row.push(0);
                    start = (i, j);
                }
                'E' => {
                    row.push(25);
                    end = (i, j);
                }
                _ => row.push(col as u8 - 97),
            }
        }

        map.push(row);
    }

    println!("{}", process(&map, start, end));

    let mut min = usize::MAX;

    for x in 0..map.len() {
        for y in 0..map[x].len() {
            if map[x][y] != 0 {
                continue;
            }

            min = min.min(process(&map, (x, y), end));
        }
    }

    println!("{}", min);
}

fn process(map: &Vec<Vec<u8>>, start: (usize, usize), end: (usize, usize)) -> usize {
    let mut visited = vec![vec![false; map[0].len()]; map.len()];
    let mut queue = VecDeque::new();

    queue.push_back((start, 0));

    while !queue.is_empty() {
        let (current, counter) = queue.pop_front().unwrap();
        if current == end {
            return counter;
        }

        let delta: Vec<(isize, isize)> = vec![(1, 0), (0, 1), (-1, 0), (0, -1)];

        for (delta_x, delta_y) in delta {
            let x = current.0 as isize + delta_x;
            let y = current.1 as isize + delta_y;

            if x < 0 || y < 0 || x as usize >= map.len() || y as usize >= map[0].len() {
                continue;
            }

            let x = x as usize;
            let y = y as usize;

            let height = map[x][y];
            if height > map[current.0][current.1] + 1 {
                continue;
            }

            if visited[x][y] {
                continue;
            }

            visited[x][y] = true;

            queue.push_back(((x, y), counter + 1));
        }
    }

    return usize::MAX;
}
