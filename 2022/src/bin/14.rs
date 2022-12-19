use aoc::read_file_input;

fn main() {
    let input = read_file_input("14.txt".to_string()).to_string();

    let mut map = vec![vec![false; 1000]; 1000];
    let mut max_y = 0;

    for line in input.lines() {
        let points: Vec<(usize, usize)> = line
            .split(" -> ")
            .map(|point| {
                let p = point.split_once(",").unwrap();
                let x = p.0.parse::<usize>().unwrap();
                let y = p.1.parse::<usize>().unwrap();

                max_y = max_y.max(y);

                return (x, y);
            })
            .collect();

        let mut prev = points[0];

        for point in points {
            let mut x = prev.0;
            let mut y = prev.1;

            map[x][y] = true;

            while x != point.0 || y != point.1 {
                if x < point.0 {
                    x += 1;
                } else if x > point.0 {
                    x -= 1;
                }
                if y < point.1 {
                    y += 1;
                } else if y > point.1 {
                    y -= 1;
                }

                map[x][y] = true;
            }

            prev = point;
        }
    }

    let mut units = 0;
    let mut map_clone = map.clone();

    for _ in 0.. {
        let mut x = 500;
        let mut y = 0;

        while y <= max_y {
            if map_clone[x][y + 1] != true {
                y += 1;
            } else if map_clone[x - 1][y + 1] != true {
                x -= 1;
                y += 1;
            } else if map_clone[x + 1][y + 1] != true {
                x += 1;
                y += 1;
            } else {
                map_clone[x][y] = true;
                units += 1;
                break;
            }
        }

        if y >= max_y {
            break;
        }
    }

    println!("{}", units);

    let mut units = 0;

    for _ in 0.. {
        let mut x = 500;
        let mut y = 0;

        loop {
            if y == max_y + 1 {
                map[x][y] = true;
                units += 1;
                break;
            } else if map[x][y + 1] != true {
                y += 1;
            } else if map[x - 1][y + 1] != true {
                x -= 1;
                y += 1;
            } else if map[x + 1][y + 1] != true {
                x += 1;
                y += 1;
            } else {
                map[x][y] = true;
                units += 1;
                break;
            }
        }

        if y == 0 {
            break;
        }
    }

    println!("{}", units);
}
