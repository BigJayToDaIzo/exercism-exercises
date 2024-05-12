// Please define the expected_minutes_in_oven function
pub fn expected_minutes_in_oven() -> Int { 40 }

// remaining_minutes_in_oven function returns the remaining minutes in oven
pub fn remaining_minutes_in_oven(elapsed_minutes: Int) -> Int {
    expected_minutes_in_oven() - elapsed_minutes
}

// preparation_time_in_minutes function returns the time in minutes
pub fn preparation_time_in_minutes(layers: Int) -> Int { 2*layers }

// total_time_in_minutes function total times in minutes
pub fn total_time_in_minutes(layers: Int, elapsed_minutes: Int) -> Int {
    preparation_time_in_minutes(layers) + elapsed_minutes
}

// alarm function Ding!s
pub fn alarm() -> String { "Ding!" }
