"""
Identify 2 integer which sum is equal to given number
"""

def get_pairs(list_values, sum_val):
    """Return a list of tuple pair which sum is equal to given sum_val.

    Parameters
    ----------
    list_values : type
        `list_values` contains all the numbers.
    sum_val : type
        Two numbers from `list_values` should be equal to `sum_val`.

    Returns
    -------
    type
        List of tuple pairs.

    """
    set_values = set()
    ans = list()
    for val in list_values:
        diff = sum_val - val
        if diff in set_values:
            ans.append((diff, val))
        set_values.add(val)
    return ans


if __name__ == '__main__':
    filename = "test_cases.txt"
    with open(filename) as file:
        content = file.readline()
        while content:
            arr, k = content.split()
            k = int(k)
            arr = [int(a) for a in arr.split(',')]

            pairs = get_pairs(arr, k)
            if bool(pairs):
                print("Following are the pairs which sum is equal to {}: {}."
                      .format(k, pairs))
            else:
                print("No pair is present which sum is euqal to {}.".format(k))

            content = file.readline()
