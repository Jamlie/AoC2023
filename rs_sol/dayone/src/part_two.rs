use std::collections::HashMap;

pub fn first_and_last_digits(s: &str) -> Result<i32, &'static str> {
    let mut numbers: HashMap<String, char> = HashMap::new();
    numbers.insert("zero".to_string(), '0');
    numbers.insert("one".to_string(), '1');
    numbers.insert("two".to_string(), '2');
    numbers.insert("three".to_string(), '3');
    numbers.insert("four".to_string(), '4');
    numbers.insert("five".to_string(), '5');
    numbers.insert("six".to_string(), '6');
    numbers.insert("seven".to_string(), '7');
    numbers.insert("eight".to_string(), '8');
    numbers.insert("nine".to_string(), '9');
    numbers.insert("0".to_string(), '0');
    numbers.insert("1".to_string(), '1');
    numbers.insert("2".to_string(), '2');
    numbers.insert("3".to_string(), '3');
    numbers.insert("4".to_string(), '4');
    numbers.insert("5".to_string(), '5');
    numbers.insert("6".to_string(), '6');
    numbers.insert("7".to_string(), '7');
    numbers.insert("8".to_string(), '8');
    numbers.insert("9".to_string(), '9');
    
    let mut first_digit = String::new();
    let mut last_digit = String::new();

    let mut result = 0;

    for (i, ch) in s.chars().enumerate() {
        let mut char_value = String::new();

        if let Some(&value) = numbers.get(&ch.to_string()) {
            char_value.push(value);
        }

        if char_value.is_empty() {
            for (word, &value) in numbers.iter() {
                if s[i..].starts_with(word) {
                    char_value.push(value);
                }
            }
        }

        if first_digit.is_empty() && !char_value.is_empty() {
            first_digit = char_value.clone();
        }

        if !char_value.is_empty() {
            last_digit = char_value.clone();
        }
    }

    result += parse_and_sum(&first_digit, &last_digit);

    if result == 0 {
        return Err("No digits found");
    }

    return Ok(result);
}

fn parse_and_sum(first_digit: &str, last_digit: &str) -> i32 {
    let final_string = format!("{}{}", first_digit, last_digit);
    let final_int = final_string.parse::<i32>().unwrap_or(-1);
    return final_int;
}
