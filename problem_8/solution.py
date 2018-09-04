def isValid(count_arr, k):
    val = 0
    for i in count_arr:
        if i > 0:
            val += 1

    return (k >= val)


def validate_string(s, k):
    """
    Check if unique chars in string s is greater than or euqal to k or not.
    """
    count_arr = [0] * 26
    unique_char = 0
    for i in range(len(s)):
        if count_arr[ord(s[i]) - ord('a')] == 0:
            unique_char += 1
        count_arr[ord(s[i]) - ord('a')] += 1

    return (unique_char >= k)


def longest_substr_k_uniq(s, k):
    if not validate_string(s, k):
        return -1

    count_arr = [0] * 26
    count_arr[ord(s[0]) - ord('a')] = 1

    curr_start = 0
    curr_end = 0

    max_window_size = 1
    max_window_start = 0

    for i in range(1, len(s)):
        count_arr[ord(s[i]) - ord('a')] += 1
        curr_end += 1

        while not isValid(count_arr, k):
            count_arr[ord(s[curr_start]) - ord('a')] -= 1
            curr_start += 1

        if curr_end - curr_start + 1 > max_window_size:
            max_window_size = curr_end - curr_start + 1
            max_window_start = curr_start

    # print("Max substring: {}".format(s[max_window_start: max_window_start + max_window_size]))
    return max_window_size


if __name__ == '__main__':
    filename = "test_cases.txt"
    with open(filename) as file:
        for line in file:
            string, k, ans = line.split()
            k = int(k)
            ans = int(ans)

            if longest_substr_k_uniq(string, k) == ans:
                print("Passed")
            else:
                print("Failed")
