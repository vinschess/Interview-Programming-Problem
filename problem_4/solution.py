def segregate(arr):
    """
    Put negative and 0 value at left size of the array.
    """
    j = 0
    for i, val in enumerate(arr):
        if val <= 0:
            arr[i], arr[j] = arr[j], arr[i] # shift to the left of array
            j += 1
    return j


def find_missing_pos(arr):
    size = len(arr)
    neg_size = segregate(arr)
    pos_size = size - neg_size
    i = neg_size
    while i < size:
        index = abs(arr[i]) - 1 + neg_size
        if (index < size and arr[index] > 0):
            arr[index] = -arr[index]
        i += 1

    i = neg_size
    while i < size:
        if arr[i] > 0:
            return i - neg_size + 1
        i += 1

    return pos_size + 1


if __name__ == '__main__':
    filename = "test_cases.txt"
    with open(filename) as file:
        content = file.readline()
        while content:
            arr, ans = content.split()
            ans = int(ans)
            arr = [int(a) for a in arr.split(',')]
            if(find_missing_pos(arr) == ans):
                print("Passed")
            else:
                print("Failed")
            content = file.readline()
