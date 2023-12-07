use itertools::Itertools;
use std::collections::HashMap;

use aoc::read_file_input;

fn main() {
    let input = read_file_input("07.txt".to_string());

    let mut hands = Vec::new();
    let mut bids: HashMap<String, i32> = HashMap::new();

    for line in input.lines() {
        let parts: Vec<&str> = line.split(' ').collect();
        let hand: &str = parts[0];

        bids.insert(hand.to_string(), parts[1].parse().unwrap());
        hands.push(hand);
    }

    let card_strengths: HashMap<char, i32> = [
        ('2', 1),
        ('3', 2),
        ('4', 3),
        ('5', 4),
        ('6', 5),
        ('7', 6),
        ('8', 7),
        ('9', 8),
        ('T', 9),
        ('J', 10),
        ('Q', 11),
        ('K', 12),
        ('A', 13),
    ]
    .iter()
    .cloned()
    .collect();

    println!(
        "{}",
        calculate(&hands, card_strengths, &bids, determine_hand_type1)
    );

    let card_strengths: HashMap<char, i32> = [
        ('J', 1),
        ('2', 2),
        ('3', 3),
        ('4', 4),
        ('5', 5),
        ('6', 6),
        ('7', 7),
        ('8', 8),
        ('9', 9),
        ('T', 10),
        ('Q', 11),
        ('K', 12),
        ('A', 13),
    ]
    .iter()
    .cloned()
    .collect();

    println!(
        "{}",
        calculate(&hands, card_strengths, &bids, determine_hand_type2)
    );
}

fn calculate(
    hands: &Vec<&str>,
    card_strengths: HashMap<char, i32>,
    bids: &HashMap<String, i32>,
    determine_hand_type: fn(&str) -> i32,
) -> i32 {
    let mut ranked_hands: Vec<(&str, i32)> = hands
        .iter()
        .map(|&hand| {
            let hand_strength = determine_hand_type(hand);
            (hand, hand_strength)
        })
        .collect();

    ranked_hands.sort_by(|a, b| {
        if a.1 == b.1 {
            let cards_a: Vec<i32> = a.0.chars().map(|c| card_strengths[&c]).collect();
            let cards_b: Vec<i32> = b.0.chars().map(|c| card_strengths[&c]).collect();

            for (i, a) in cards_a.iter().enumerate() {
                if cards_b[i] == *a {
                    continue;
                } else {
                    return cards_b[i].cmp(&a);
                }
            }

            return std::cmp::Ordering::Equal;
        } else {
            return b.1.cmp(&a.1);
        }
    });

    let mut output = 0;

    for (i, hand) in ranked_hands.iter().enumerate() {
        output += bids[hand.0] * (ranked_hands.len() - i) as i32;
    }

    return output;
}

fn determine_hand_type1(hand: &str) -> i32 {
    let mut counts = [0; 13];
    for card in hand.chars() {
        match card {
            '2' => counts[0] += 1,
            '3' => counts[1] += 1,
            '4' => counts[2] += 1,
            '5' => counts[3] += 1,
            '6' => counts[4] += 1,
            '7' => counts[5] += 1,
            '8' => counts[6] += 1,
            '9' => counts[7] += 1,
            'T' => counts[8] += 1,
            'J' => counts[9] += 1,
            'Q' => counts[10] += 1,
            'K' => counts[11] += 1,
            'A' => counts[12] += 1,
            _ => {}
        }
    }
    match counts.iter().max().unwrap() {
        5 => 7,
        4 => 6,
        3 if counts.iter().filter(|&&x| x == 2).count() == 1 => 5,
        3 => 4,
        2 if counts.iter().filter(|&&x| x == 2).count() == 2 => 3,
        2 => 2,
        _ => 1,
    }
}

fn determine_hand_type2(hand: &str) -> i32 {
    let mut counts = [0; 12];
    for card in hand.chars() {
        match card {
            '2' => counts[0] += 1,
            '3' => counts[1] += 1,
            '4' => counts[2] += 1,
            '5' => counts[3] += 1,
            '6' => counts[4] += 1,
            '7' => counts[5] += 1,
            '8' => counts[6] += 1,
            '9' => counts[7] += 1,
            'T' => counts[8] += 1,
            'Q' => counts[9] += 1,
            'K' => counts[10] += 1,
            'A' => counts[11] += 1,
            _ => {}
        }
    }

    let jokers = hand.matches('J').count();
    let ranks = counts.iter().sorted().rev().cloned().collect::<Vec<_>>();

    if jokers == 5 || ranks[0] + jokers >= 5 {
        return 7;
    } else if ranks[0] + jokers >= 4 {
        return 6;
    } else if ranks[0] + ranks[1] + jokers >= 5 {
        return 5;
    } else if ranks[0] + jokers >= 3 {
        return 4;
    } else if ranks[0] + ranks[1] + jokers >= 4 {
        return 3;
    } else if ranks[0] + jokers >= 2 {
        return 2;
    } else {
        return 1;
    }
}
