// can eat_ghost plz
pub fn eat_ghost(power_pellet_active: Bool, touching_ghost: Bool) -> Bool {
  power_pellet_active && touching_ghost
}
// can eat_dot or pellet plz
pub fn score(touching_power_pellet: Bool, touching_dot: Bool) -> Bool {
  touching_power_pellet || touching_dot
}
// can not die plz
pub fn lose(power_pellet_active: Bool, touching_ghost: Bool) -> Bool {
  !power_pellet_active && touching_ghost
}
// can win plz
pub fn win(
  has_eaten_all_dots: Bool,
  power_pellet_active: Bool,
  touching_ghost: Bool,
) -> Bool {
  has_eaten_all_dots && power_pellet_active || !touching_ghost
}
