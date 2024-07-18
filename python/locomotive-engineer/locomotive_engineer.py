"""Functions which helps the locomotive engineer to keep track of the train."""


def get_list_of_wagons(*wagons):
    """Return a list of wagons.

    :param: arbitrary number of wagons.
    :return: list - list of wagons.
    """
    return list(wagons)


def fix_list_of_wagons(each_wagons_id, missing_wagons):
    """Fix the list of wagons.

    :param each_wagons_id: list - the list of wagons.
    :param missing_wagons: list - the list of missing wagons.
    :return: list - list of wagons.
    """
    m1, m2, *fixed = each_wagons_id
    fixed.append(m1)
    fixed.append(m2)
    i = 1
    for wagon in missing_wagons:
        fixed.insert(i, wagon)
        i += 1
    return list(fixed)


def add_missing_stops(route, **stops):
    """Add missing stops to route dict.

    :param route: dict - the dict of routing information.
    :param: arbitrary number of stops.
    :return: dict - updated route dictionary.
    """
    route.update({"stops": list(stops.values())})
    return route


def extend_route_information(route, more_route_information):
    """Extend route information with more_route_information.

    :param route: dict - the route information.
    :param more_route_information: dict -  extra route information.
    :return: dict - extended route information.
    """
    new_route = {**route, **more_route_information}
    return new_route


def fix_wagon_depot(wagons_rows):
    """Fix the list of rows of wagons.

    :param wagons_rows: list[list[tuple]] - the list of rows of wagons.
    :return: list[list[tuple]] - list of rows of wagons.
    """
    wagons_rows[0][1], wagons_rows[1][0] = wagons_rows[1][0], wagons_rows[0][1]
    wagons_rows[1][2], wagons_rows[2][1] = wagons_rows[2][1], wagons_rows[1][2]
    wagons_rows[2][0], wagons_rows[0][2] = wagons_rows[0][2], wagons_rows[2][0]
    return wagons_rows
