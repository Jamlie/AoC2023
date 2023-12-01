mod part_one;
mod part_two;

fn main() {
    let mut i = part_one::first_and_last_digits("1abc2");
    match i {
        Ok(i) => println!("{}", i),
        Err(e) => println!("{}", e),
    }

    i = part_one::first_and_last_digits("pqr3stu8vwx");
    match i {
        Ok(i) => println!("{}", i),
        Err(e) => println!("{}", e),
    }

    i = part_one::first_and_last_digits("a1b2c3d4e5f");
    match i {
        Ok(i) => println!("{}", i),
        Err(e) => println!("{}", e),
    }

    i = part_one::first_and_last_digits("treb7uchet");
    match i {
        Ok(i) => println!("{}", i),
        Err(e) => println!("{}", e),
    }

    // {"two1nine", 29},
    // {"eightwothree", 83},
    // {"abcone2threexyz", 13},
    // {"xtwone3four", 24},
    // {"4nineeightseven2", 42},
    // {"zoneight234", 14},
    // {"7pqrstsixteen", 76},

    let mut j = part_two::first_and_last_digits("two1nine");
    match j {
        Ok(j) => println!("{}", j),
        Err(e) => println!("{}", e),
    }

    j = part_two::first_and_last_digits("eightwothree");
    match j {
        Ok(j) => println!("{}", j),
        Err(e) => println!("{}", e),
    }

    j = part_two::first_and_last_digits("abcone2threexyz");
    match j {
        Ok(j) => println!("{}", j),
        Err(e) => println!("{}", e),
    }

    j = part_two::first_and_last_digits("xtwone3four");
    match j {
        Ok(j) => println!("{}", j),
        Err(e) => println!("{}", e),
    }

    j = part_two::first_and_last_digits("4nineeightseven2");
    match j {
        Ok(j) => println!("{}", j),
        Err(e) => println!("{}", e),
    }

    j = part_two::first_and_last_digits("zoneight234");
    match j {
        Ok(j) => println!("{}", j),
        Err(e) => println!("{}", e),
    }

    j = part_two::first_and_last_digits("7pqrstsixteen");
    match j {
        Ok(j) => println!("{}", j),
        Err(e) => println!("{}", e),
    }
}
