use std::fmt;

#[derive(Debug, PartialEq)]
pub struct Clock{
    hours: i32,
    minutes: i32,
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.hours, self.minutes) 
    }
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let mut hours = hours;
        let mut minutes = minutes;
        while minutes < 0 {
            minutes += 60;
            hours -= 1;
        }
        while hours < 0 {
            hours += 24;
        }
        hours = (hours + minutes / 60) % 24;
        minutes %= 60;
        Clock{hours, minutes}
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Clock::new(self.hours + minutes / 60, self.minutes + minutes % 60)
    }
}
