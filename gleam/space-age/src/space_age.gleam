const seconds_per_year = 31_556_600.0

pub type Planet {
  Mercury
  Venus
  Earth
  Mars
  Jupiter
  Saturn
  Uranus
  Neptune
}

fn earth_to_planet_period(planet: Planet) -> Float {
  case planet {
    Mercury -> 0.2408467
    Venus -> 0.61519726
    Earth -> 1.0
    Mars -> 1.8808158
    Jupiter -> 11.862615
    Saturn -> 29.447498
    Uranus -> 84.016846
    Neptune -> 164.79132
  }
}

pub fn age(planet: Planet, seconds: Float) -> Float {
  seconds /. seconds_per_year /. earth_to_planet_period(planet)
}
