use aoc::read_file_input;
use serde_json::Value;
use std::cmp::Ordering;

fn main() {
    let input = read_file_input("13.txt".to_string()).to_string();

    let mut packets: Vec<Value> = vec![];
    let mut score = 0;

    for (i, pair) in input.split("\n\n").enumerate() {
        let (left, right) = pair.split_once('\n').unwrap();

        let left = serde_json::from_str::<Value>(left).unwrap();
        let right = serde_json::from_str::<Value>(right).unwrap();

        if compare(&left, &right) == Ordering::Less {
            score += i + 1;
        }

        packets.push(left);
        packets.push(right);
    }

    println!("{}", score);

    let dividers = [
        serde_json::from_str::<Value>("[[2]]").unwrap(),
        serde_json::from_str::<Value>("[[6]]").unwrap(),
    ];

    for divider in dividers.iter() {
        packets.push(divider.clone())
    }

    packets.sort_by(compare);

    let mut key = 1;

    for (i, packet) in packets.iter().enumerate() {
        if dividers.contains(packet) {
            key *= i + 1;
        }
    }

    println!("{}", key);
}

fn compare(a: &Value, b: &Value) -> Ordering {
    match (a, b) {
        (Value::Array(a), Value::Array(b)) => {
            for i in 0..a.len().max(b.len()) {
                match (a.get(i), b.get(i)) {
                    (None, _) => return Ordering::Less,
                    (_, None) => return Ordering::Greater,
                    (Some(x), Some(y)) => match compare(x, y) {
                        Ordering::Equal => {}
                        o => return o,
                    },
                }
            }
            Ordering::Equal
        }
        (Value::Array(_), Value::Number(_)) => compare(a, &Value::Array(vec![b.clone()])),
        (Value::Number(x), Value::Number(y)) => x.as_u64().unwrap().cmp(&y.as_u64().unwrap()),
        (Value::Number(_), Value::Array(_)) => compare(&Value::Array(vec![a.clone()]), b),
        _ => unreachable!(),
    }
}
