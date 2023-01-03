use std::collections::HashMap;

use aoc::read_file_input;

fn main() {
    let input = read_file_input("17.txt".to_string());

    println!("{}", part_one(&input));
    println!("{}", part_two(&input));
}

#[derive(Clone)]
struct Rock {
    shape: [[usize; 4]; 4],
    x: usize,
    y: usize,
    width: usize,
    height: usize,
}

const ROCKS: [Rock; 5] = [
    Rock {
        shape: [[1; 4], [0; 4], [0; 4], [0; 4]],
        x: 2,
        y: 0,
        width: 4,
        height: 1,
    },
    Rock {
        shape: [[0, 1, 0, 0], [1, 1, 1, 0], [0, 1, 0, 0], [0; 4]],
        x: 2,
        y: 0,
        width: 3,
        height: 3,
    },
    Rock {
        shape: [[0, 0, 1, 0], [0, 0, 1, 0], [1, 1, 1, 0], [0; 4]],
        x: 2,
        y: 0,
        width: 3,
        height: 3,
    },
    Rock {
        shape: [[1, 0, 0, 0]; 4],
        x: 2,
        y: 0,
        width: 1,
        height: 4,
    },
    Rock {
        shape: [[1, 1, 0, 0], [1, 1, 0, 0], [0; 4], [0; 4]],
        x: 2,
        y: 0,
        width: 2,
        height: 2,
    },
];

fn collide(chamber: &[[bool; 7]], rock: &Rock, dx: isize, dy: isize) -> bool {
    for x in 0..rock.width {
        for y in 0..rock.height {
            if rock.shape[y][x] == 0 {
                continue;
            }
            if let Some(row) = chamber.get((rock.y as isize + dy) as usize - y) {
                if *row
                    .get((rock.x as isize + dx) as usize + x)
                    .unwrap_or(&false)
                {
                    return true;
                }
            }
        }
    }
    false
}

fn drop(
    count: usize,
    highest: &mut usize,
    chamber: &mut Vec<[bool; 7]>,
    pattern: &str,
    jet: &mut usize,
) {
    let mut rock = ROCKS[count % ROCKS.len()].clone();
    rock.y = *highest + rock.height + 3;

    if rock.y >= chamber.len() {
        chamber.extend(vec![[false; 7]; rock.y - chamber.len()]);
    }

    loop {
        match pattern.as_bytes()[*jet % pattern.len()] {
            b'<' => {
                if rock.x != 0 && !collide(&*chamber, &rock, -1, 0) {
                    rock.x -= 1;
                }
            }
            b'>' => {
                if rock.x + rock.width < 7 && !collide(&*chamber, &rock, 1, 0) {
                    rock.x += 1;
                }
            }
            _ => {}
        }

        *jet += 1;

        if rock.y == rock.height || collide(&*chamber, &rock, 0, -1) {
            break;
        }

        rock.y -= 1;
    }

    if rock.y > *highest {
        *highest = rock.y;
    }

    for x in 0..rock.width {
        for y in 0..rock.height {
            chamber[rock.y - y][rock.x + x] |= rock.shape[y][x] == 1;
        }
    }
}

pub fn part_one(pattern: &str) -> usize {
    let mut chamber = vec![[false; 7]];
    let mut highest = 0;
    let mut jet = 0;

    for i in 0..2022 {
        drop(i, &mut highest, &mut chamber, pattern, &mut jet);
    }

    highest
}

fn columns(chamber: &[[bool; 7]], highest: usize) -> [usize; 7] {
    let mut heights = [0; 7];

    for (i, height) in heights.iter_mut().enumerate() {
        *height = (0..highest)
            .find(|&x| chamber[highest - x][i])
            .unwrap_or(usize::MAX);
    }

    heights
}

pub fn part_two(pattern: &str) -> usize {
    let mut chamber = vec![[false; 7]];
    let mut highest = 0;
    let mut total_height = 0;
    let mut jet = 0;
    let mut count = 0usize;
    let mut cache = HashMap::new();

    while count <= 1000000000000 {
        drop(count, &mut highest, &mut chamber, pattern, &mut jet);

        let key = (
            count % ROCKS.len(),
            jet % pattern.len(),
            columns(&chamber, highest),
        );

        if let Some((c, height)) = cache.get(&key) {
            let repeat = (1000000000000 - c) / (count - c) - 1;
            count += (count - c) * repeat;
            total_height += (highest - height) * repeat;
        } else {
            cache.insert(key, (count, highest));
        }

        count += 1;
    }

    total_height + highest - 1
}
