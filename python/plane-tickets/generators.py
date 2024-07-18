"""Functions to automate Conda airlines ticketing system."""


def generate_seat_letters(number):
    """Generate a series of letters for airline seats.

    :param number: int - total number of seat letters to be generated.
    :return: generator - generator that yields seat letters.

    Seat letters are generated from A to D.
    After D it should start again with A.

    Example: A, B, C, D

    """
    seat_letters = ["A", "B", "C", "D"]
    yield from (seat_letters[i % 4] for i in range(number))


def generate_seats(number):
    """Generate a series of identifiers for airline seats.

    :param number: int - total number of seats to be generated.
    :return: generator - generator that yields seat numbers.

    A seat number consists of the row number and the seat letter.

    There is no row 13.
    Each row has 4 seats.

    Seats should be sorted from low to high.

    Example: 3C, 3D, 4A, 4B

    """
    while number > 0:
        for row in range(1, 22):  # Problem shows 21 rows in example
            if row == 13:
                continue
            if number >= 4:
                for seat in generate_seat_letters(4):
                    yield f"{row}{seat}"
                    number -= 1
            else:
                for seat in generate_seat_letters(number):
                    yield f"{row}{seat}"
                    number -= 1


def assign_seats(passengers):
    """Assign seats to passengers.

    :param passengers: list[str] - a list of strings containing names of passengers.
    :return: dict - with the names of the passengers as keys and seat numbers as values.

    Example output: {"Adele": "1A", "Björk": "1B"}

    """
    seat_assignments = {}
    for seat, passenger in zip(generate_seats(len(passengers)), passengers):
        seat_assignments[passenger] = seat
    return seat_assignments


def generate_codes(seat_numbers, flight_id):
    """Generate codes for a ticket.

    :param seat_numbers: list[str] - list of seat numbers.
    :param flight_id: str - string containing the flight identifier.
    :return: generator - generator that yields 12 character long ticket codes.

    """
    for seat in seat_numbers:
        s = f"{seat}{flight_id}"
        yield s.ljust(12, "0")
