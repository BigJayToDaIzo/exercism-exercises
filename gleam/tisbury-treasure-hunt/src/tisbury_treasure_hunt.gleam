pub fn place_location_to_treasure_location(
  place_location: #(String, Int),
) -> #(Int, String) {
  #(place_location.1, place_location.0)
}

pub fn treasure_location_matches_place_location(
  place_location: #(String, Int),
  treasure_location: #(Int, String),
) -> Bool {
  case place_location {
    #(place_name, place_distance) -> {
      case treasure_location {
        #(treasure_distance, treasure_name) -> {
          place_distance == treasure_distance && place_name == treasure_name
        }
      }
    }
  }
}

pub fn count_place_treasures(
  place: #(String, #(String, Int)),
  treasures: List(#(String, #(Int, String))),
) -> Int {
  let count = 0
  // for each treasure in treasures
  case treasures {
    // if the treasure_location_matches_place_location
    [treasure, ..rest] -> {
      case treasure {
        #(treasure_name, treasure_location) -> {
          case treasure_location_matches_place_location(place, treasure_location) {
            let count = count + 1
          }
          let count = count_place_treasures(place, rest)
        }
      }
    }
    [] -> count
  }
}

pub fn special_case_swap_possible(
  found_treasure: #(String, #(Int, String)),
  place: #(String, #(String, Int)),
  desired_treasure: #(String, #(Int, String)),
) -> Bool {
  todo
}
