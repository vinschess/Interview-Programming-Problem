"""
Write custom square root function
"""

def sqrt(val):
    """Custom square root function.

    Parameters
    ----------
    val : type
        Identify the square root of `val`.

    Returns
    -------
    type
        Sqaure root value as float.

    """
    sqrt_val = val/2
    prev = 0.0
    while sqrt_val-prev >= 0.001 or prev-sqrt_val >= 0.001:
        prev = sqrt_val
        sqrt_val -= (sqrt_val*sqrt_val - val) / (2*sqrt_val)
    return round(sqrt_val, 2)


if __name__ == '__main__':
    print(sqrt(1))
    print(sqrt(16))
    print(sqrt(25))
    print(sqrt(100))
    print(sqrt(101))
    print(sqrt(176))
