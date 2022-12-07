use std::{fs, collections::VecDeque};
use regex::{Regex, Captures};

const FILE: &str = "../input.txt";

fn parse_stacks<'a>(input:&'a String, retain_order:bool){
    //Identify stack count
    let re = Regex::new(r"(?m)(?P<s>\d)(?:[ ]$)").unwrap();
    let c = re.captures(input).unwrap();
    let v = &c["s"];
    let stack_count: i32 = v.parse().unwrap();

    //Create stacks
    let mut stacks = vec![VecDeque::<&str>::new()];
    for _n in 0..stack_count -1 {
        let _ = &stacks.push(VecDeque::<&str>::new());
    }

    //Add crates to stacks
    let items: Regex = Regex::new(r"(?m)\[\w\]").unwrap();
    for i in items.find_iter(input)
    {
        let stack_num = (i.start() / 4) % stack_count as usize;
        let m = i.as_str().trim_start_matches('[').trim_end_matches(']');
        let _ = &stacks[stack_num].push_back(m);
    }

    //Parse instructions
    let re = Regex::new(r"(?m)^move.*\n").unwrap();
    let parse_moves = Regex::new(r"(?:\s)(\d*)(?:\s)").unwrap();
    
    for i in re.find_iter(input){
        //Parse instruction values
        let move_string = i.as_str();
        let inst : Vec<Captures> = parse_moves.captures_iter(move_string).collect();
        let c : i32 = extract_int(inst.get(0));
        let from = extract_int(inst.get(1)) as usize;
        let to = extract_int(inst.get(2)) as usize;

        //Work out which crates need moving, add them to a temporary stack
        let mut to_move = VecDeque::<&str>::new();
        for _x in 0..c{
            let p = stacks[from - 1].pop_front().unwrap();

            if retain_order
            {
                to_move.push_front(p);
            } else{
                to_move.push_back(p);
            }
        }

        //Move em!
        for _x in to_move{
            let _ = stacks[to - 1].push_front(_x);
        }
    }

    println!("Top crates");
    for s in stacks{
        let t = s.get(0).unwrap();
        print!("{t}");
    }
    println!("");
}

fn extract_int(c:Option<&Captures>) -> i32{
    return c.unwrap().get(0).unwrap().as_str().trim().parse().unwrap();
}

fn main() {
    let contents = fs::read_to_string(FILE)
        .expect("Should have been able to read the file");

    println!("Part 1: One at a time");
    parse_stacks(&contents, false);
    println!("");
    println!("Part 2: Multi-move");
    parse_stacks(&contents, true);
}
