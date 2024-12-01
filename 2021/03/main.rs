#![allow(unused)]
use std::fs::File;
use std::io::{prelude::*, BufReader};

fn main() {
    use std::env;
    let args: Vec<String> = env::args().collect();
    if args.len() != 2 {
        println!("Usage: d02 <input filename>");
    } else {
        let input = read_file(&args[1]);
        let gamma : i32 = calc_values(&input, "gamma".to_string());
        let epsilon : i32 = calc_values(&input, "epsilon".to_string());
        println!("gamma: {}, epsilon: {}, answer part1: {}", gamma, epsilon, gamma*epsilon);
    }    
}

fn read_file(filename: &str) -> Vec<String> {
    let file = File::open(filename).expect("oopse");
    let reader = BufReader::new(file);
    let mut intlist: Vec<String> = Vec::new();

    for line in reader.lines() {
        let load: String = line.expect("ops").trim().parse().unwrap();
        intlist.extend(vec![load]);
    }
    return intlist;
}

fn calc_values(liste: &Vec<String>, operator: String) -> i32{
    
    let mut len = 0;

    for item in liste { //Todo: Find a nicer way to do this
        len = item.len();
        break;
    }

    let zero_vec = vec![0; len];
    let mut entries: i32 = 0;
    let mut entries_at_len: i32 = 0;
    let mut vec_zeros = vec![0; len];;
    let mut vec_ones = vec![0; len];;

    for item in liste {
        for (i, c) in item.chars().enumerate() {
            if c == '0' {
                vec_zeros[i] += 1;
            } else if c == '1' {
                vec_ones[i] += 1;
            } else {
                panic!("Invalid character");
            }
            if i == len {
                entries_at_len += 1;
            }
        }
        entries += 1;
        if entries != entries_at_len {
            println!("Invalid entry: {} at line {}", item, entries);
        }
    }

    let mut vec_result = vec![0; len];;
    for i in 0..5 {
        if operator == "gamma".to_string(){
            if vec_zeros[i] > vec_ones[i] {
                vec_result[i] = 0;
            } else {
                vec_result[i] = 1;
            }
        } else if operator == "epsilon".to_string() {
            if vec_zeros[i] < vec_ones[i] {
                vec_result[i] = 0;
            } else {
                vec_result[i] = 1;
            }
        }
    }
    
    let result = vec_result.into_iter().map(|i| i.to_string()).collect::<String>();
    
    //return result.parse::<i32>().unwrap();
    return i32::from_str_radix(&result, 2).unwrap();
}


 