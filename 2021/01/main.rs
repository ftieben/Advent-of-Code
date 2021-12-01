use std::fs::File;
use std::io::{prelude::*, BufReader};

fn main() { 
    part_one(&read_list("input.txt"));
    
    part_two(&read_list("input.txt"));
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
    let mut last: i32 = 0;
    let mut first: bool = true;
    let mut increased: i32 = 0;

    for item in liste {
        if !first {
            if item > &last {
                increased = increased +1 ;
            }
            last = *item;
        } else {
            first = false;
        }
       
    }
    println!("{}", increased);
}
 
fn part_two(liste: &Vec<i32>) {
    let mut one_ago: i32 = 0;
    let mut two_ago: i32 = 0;
    
    let mut cur_window_value: i32 = 0;
    let mut last_window_value: i32 = 0;


    let mut first: bool = true;
    let mut second: bool = true;
    let mut third: bool = true;

    let mut status: String ="".to_string();
    let mut position: i32 = liste.len() as i32 % 3 -1;
    println!("started at position {}", position);
    let mut increased: i32 = 0;
    

    for item in liste {
        if position == liste.len() as i32 {
            println!("stopped at position {}", position);
            //break;
        }
        if !first {
            if !second {
                cur_window_value = one_ago + two_ago + item;
                if !third {
                    if cur_window_value > last_window_value {
                        increased = increased +1 ;
                        status = "increased".to_string();
                    }else if cur_window_value == last_window_value {
                        status = "equal".to_string();
                    }else{
                        status = "decreased".to_string();
                    }
                }else {
                    third = false;
                }
                
            } else {
                two_ago = *item;
                second = false;
            }
        } else {
            one_ago = *item;
            first = false;
        }
       position = position +1;

       println!("pos:{} value:{} lastWindow:{} curWindow:{} status:{}", position,item, last_window_value, cur_window_value, status);
       last_window_value = cur_window_value;
    }
    println!("{}", increased);
}