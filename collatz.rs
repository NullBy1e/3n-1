use std::env;

fn collatz(number: i32) {
  print!("{},", number);
  if number == 1 || number == 2 {
    print!("\n");
  } else {
    if number % 2 == 0 {
      let final_number = number/2;
      collatz(final_number);
    } else {
      let final_number = number*3+1;
      collatz(final_number); 
    }
  }
}

fn main(){
  let argv: Vec<String> = env::args().collect();
  // it's 32bit so max value can be 2147483647
  let mut n: i32 = 0; 
  if argv.len() < 2 {
    println!("pass some arguments to argv");
  } else if argv.len() >= 2 {
    // parse string to 32bit int
    n = argv[1].parse::<i32>().unwrap();
  }

  for i in 1..=n {
    collatz(i);
  }
}
