"""
A Product Array Puzzle without division operator
"""

def product_array(arr):
    n = len(arr)
    prod_arr = [1] * n
    temp = 1
    for i in range(n):
        prod_arr[i] = temp
        temp *= arr[i]

    temp = 1
    for i in range(n-1, -1, -1):
        prod_arr[i] *= temp
        temp *= arr[i]

    return prod_arr


if __name__ == '__main__':
    filename = "test_cases.txt"
    with open(filename) as file:
        content = file.readline()
        while content:

            arr = [int(a) for a in content.split(',')]
            print("Product array: {}".format(product_array(arr)))

            content = file.readline()
