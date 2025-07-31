import gleam/list

pub fn place_location_to_treasure_location(
  place_location: #(String, Int),
) -> #(Int, String) {
  #(place_location.1, place_location.0)
}

pub fn treasure_location_matches_place_location(
  place_location: #(String, Int),
  treasure_location: #(Int, String),
) -> Bool {
  case place_location, treasure_location {
    p, t if p.0 == t.1 && p.1 == t.0 -> True
    _, _ -> False
  }
}

pub fn count_place_treasures(
  place: #(String, #(String, Int)),
  treasures: List(#(String, #(Int, String))),
) -> Int {
  list.filter(treasures, fn(t) { t.1.1 == place.1.0 })
  |> list.length
}

pub fn special_case_swap_possible(
  found_treasure: #(String, #(Int, String)),
  place: #(String, #(String, Int)),
  desired_treasure: #(String, #(Int, String)),
) -> Bool {
  case found_treasure, place, desired_treasure {
    ft, place, _
      if ft.0 == "Brass Spyglass" && place.0 == "Abandoned Lighthouse"
    -> True
    ft, place, dt
      if ft.0 == "Amethyst Octopus"
      && place.0 == "Stormy Breakwater"
      && { dt.0 == "Crystal Crab" || dt.0 == "Glass Starfish" }
    -> True
    ft, place, dt
      if ft.0 == "Vintage Pirate Hat"
      && place.0 == "Harbor Managers Office"
      && {
        dt.0 == "Model Ship in Large Bottle"
        || dt.0 == "Antique Glass Fishnet Float"
      }
    -> True
    _, _, _ -> False
  }
}
