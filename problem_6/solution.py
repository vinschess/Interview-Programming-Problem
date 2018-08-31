def largest_non_adj_sum(arr):
    first = 0
    next = 0
    for i in arr:
        temp = max(first, next)

        first = next + i
        next = temp

    return max(first, next)


if __name__ == '__main__':
    filename = "test_cases.txt"
    with open(filename) as file:
        content = file.readline()
        while content:
            arr, ans = content.split()
            ans = int(ans)
            arr = [int(a) for a in arr.split(',')]

            if(largest_non_adj_sum(arr) == ans):
                print("Passed")
            else:
                print("Failed")
            content = file.readline()
