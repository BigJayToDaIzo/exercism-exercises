import math


def is_armstrong_number(number):
    numDigits = len(str(number))
    armNumCandidate = 0
    for i in range(numDigits):
        armNumCandidate += math.pow(int(str(number)[i]), numDigits)
    return armNumCandidate == number
