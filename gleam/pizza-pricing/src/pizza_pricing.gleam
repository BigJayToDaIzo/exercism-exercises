pub type Pizza {
  Margherita
  Caprese
  Formaggio
  ExtraSauce(Pizza)
  ExtraToppings(Pizza)
}

pub fn pizza_price(pizza: Pizza) -> Int {
  case pizza {
    Margherita -> 7
    Caprese -> 9
    Formaggio -> 10
    ExtraSauce(p) -> pizza_price(p) + 1
    ExtraToppings(p) -> pizza_price(p) + 2
  }
}

pub fn order_price(order: List(Pizza)) -> Int {
  todo
}
