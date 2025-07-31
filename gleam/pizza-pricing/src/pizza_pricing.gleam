import gleam/list

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
  case order {
    [] -> 0
    [p] -> {
      3 + pizza_price(p)
    }
    [p1, p2] -> {
      2 + pizza_price(p1) + pizza_price(p2)
    }
    _ -> {
      let acc = 0
      list.fold(order, acc, fn(acc, p) {acc + pizza_price(p)})
    }
  }
}

