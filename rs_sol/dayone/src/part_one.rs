use std::num::ParseIntError;

pub fn first_and_last_digits(s: &str) -> Result<i32, &'static str> {
    let mut first_digit: Option<char> = None;
    let mut last_digit: Option<char> = None;

    for ch in s.chars() {
        if ch.is_digit(10) {
            if first_digit.is_none() {
                first_digit = Some(ch);
            }
            last_digit = Some(ch);
        }
    }

    if let (Some(first), Some(last)) = (first_digit, last_digit) {
        let final_string = format!("{}{}", first, last);
        let final_int = final_string.parse::<i32>().unwrap();
        return Ok(final_int);
    }

    return Err("No digits found");
}
