pub type Approval {
  Yes
  No
  Maybe
}

pub type Cuisine {
  Korean
  Turkish
}

pub type Genre {
  Crime
  Horror
  Romance
  Thriller
}

pub type Activity {
  BoardGame
  Chill
  Movie(Genre)
  Restaurant(Cuisine)
  Walk(Int)
}

pub fn rate_activity(activity: Activity) -> Approval {
  case activity {
    BoardGame | Chill -> No
    Movie(g) -> case g {
      Romance -> Yes
      _ -> No
    }
    Restaurant(c) -> case c {
      Korean -> Yes
      Turkish -> Maybe
    }
    Walk(m) -> case m {
      m if m > 11 -> Yes
      m if m > 6 -> Maybe
      _ -> No
    }
  }
}
