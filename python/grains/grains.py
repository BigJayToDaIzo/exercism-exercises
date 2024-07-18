import math


def square(number):
    if number < 1 or number > 64:
        raise ValueError("square must be between 1 and 64")
    return int(math.pow(2, number - 1))


def total():
    sum = 0
    for i in range(1, 65):
        sum += square(i)
    return sum
