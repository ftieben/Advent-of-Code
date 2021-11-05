use std::fs::File;
use std::io::{prelude::*, BufReader};

fn main() { 
    part_one(&read_list("../input.txt"));
    part_two(&read_list("../input.txt"));
}

fn read_list(filename: &str) -> Vec<i32> {
    let file = File::open(filename).expect("oopse");
    let reader = BufReader::new(file);
    let mut intlist: Vec<i32> = Vec::new();

    for line in reader.lines() {
        let load: i32 = line.expect("ops").trim().parse().unwrap();
        intlist.extend(vec![load]) ;
    }
    return intlist
}

fn part_one(liste: &Vec<i32>) {
    let mut breaker: bool;
    breaker = false;

    for item1 in liste{
        for item2 in liste{
            if item1 + item2 == 2020 {
                let sum = item1 * item2;
                println!("{}",sum);
                breaker = true;
                break;
            }
        }
        if breaker{
            break
        }
    }    
}

fn part_two(liste: &Vec<i32>) {
    let mut breaker: bool;
    breaker = false;

    for item1 in liste{
        for item2 in liste{
            for item3 in liste{
                if item1 + item2 + item3 == 2020 {
                    let sum = item1 * item2 * item3 ;
                    println!("{}",sum);
                    breaker = true;
                    break;
                }
            }
            if breaker{break}
        }
        if breaker{break}
    }    
}