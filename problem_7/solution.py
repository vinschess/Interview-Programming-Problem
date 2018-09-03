# Where m represents total ways of steps at a time
def steps_count(n, m=2):
    arr = [0 for a in range(n+1)]
    arr[0] = arr[1] = 1
    for i in range(2, n+1):
        j = 1
        while j <= m and j <= i:
            arr[i] += arr[i - j]
            j += 1
    return arr[n]


if __name__ == '__main__':
    filename = "test_cases.txt"
    with open(filename) as file:
        for line in file:
            steps, ans = line.split()
            steps = int(steps)
            ans = int(ans)

            if steps_count(steps) == ans :
                print("Passed")
            else:
                print("Failed")
