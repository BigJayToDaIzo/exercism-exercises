import gleam/int
import gleam/list

pub type Command {
  Wink
  DoubleBlink
  CloseYourEyes
  Jump
}

pub fn commands(encoded_message: Int) -> List(Command) {
  // reversed because we build with prepend (faster)
  let reversed_commands =
    [#(1, Wink), #(2, DoubleBlink), #(4, CloseYourEyes), #(8, Jump)]
    |> list.fold(from: [], with: fn(acc, command_trans) {
      case int.bitwise_and(encoded_message, command_trans.0) {
        0 -> acc
        _ -> [command_trans.1, ..acc]
      }
    })
  // since commands are already reversed, we need to unreverse if 16s bit is off
  case int.bitwise_and(encoded_message, 16) {
    0 -> list.reverse(reversed_commands)
    _ -> reversed_commands
  }
}
