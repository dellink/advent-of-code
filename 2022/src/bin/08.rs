use aoc::read_file_input;

fn main() {
    let input = read_file_input("08.txt".to_string());

    let mut trees = Vec::new();

    for line in input.lines() {
        let mut tree: Vec<usize> = Vec::new();
        for char in line.split("").filter(|x| !x.is_empty()) {
            tree.push(char.parse().unwrap());
        }
        trees.push(tree);
    }

    let mut count = 0;

    for row in 0..trees.len() {
        for col in 0..trees[row].len() {
            let height = trees[row][col];
            if trees[..row].iter().all(|x| x[col] < height)
                || trees[row][..col].iter().all(|x| x < &height)
                || trees[row + 1..].iter().all(|x| x[col] < height)
                || trees[row][col + 1..].iter().all(|x| x < &height)
            {
                count += 1;
            }
        }
    }

    println!("{:?}", count);

    let mut count = 0;

    for row in 0..trees.len() {
        for col in 0..trees[row].len() {
            let height = trees[row][col];
            let mut best = 1;

            calculate(&mut best, height, trees[..row].iter().map(|x| x[col]).rev());
            calculate(&mut best, height, trees[row][..col].iter().rev().copied());
            calculate(&mut best, height, trees[row + 1..].iter().map(|x| x[col]));
            calculate(&mut best, height, trees[row][col + 1..].iter().copied());

            count = count.max(best);
        }
    }

    println!("{:?}", count);
}

fn calculate(best: &mut usize, height: usize, trees: impl Iterator<Item = usize>) {
    let mut count = 0;
    for tree in trees {
        count += 1;
        if tree >= height {
            break;
        }
    }
    *best *= count;
}
