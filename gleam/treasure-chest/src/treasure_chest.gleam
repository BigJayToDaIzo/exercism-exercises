pub type TreasureChest(a) {
  TreasureChest(password: String, treasure: a)
}

pub type UnlockResult(a) {
  Unlocked(a)
  WrongPassword
}

pub fn get_treasure(
  chest: TreasureChest(treasure),
  password: String,
) -> UnlockResult(treasure) {
  case password {
    p if p == chest.password -> Unlocked(chest.treasure)
    _ -> WrongPassword
  }
}
