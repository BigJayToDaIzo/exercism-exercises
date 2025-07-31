pub fn is_armstrong_number(num: u32) -> bool {
    let pow: u32 = num.to_string().len().try_into().unwrap();
    let mut sum_of_powered_digits = 0;
    let mut digits = num;
    while digits > 0 {
        let digit = digits % 10;
        sum_of_powered_digits += u32::pow(digit, pow);
        digits /= 10;
    }
    sum_of_powered_digits == num
}
