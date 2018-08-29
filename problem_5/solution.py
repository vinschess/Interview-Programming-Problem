def decode_count(str_digit):
    count_2nd_last = 1
    count_last = 1
    count_curr = 1
    n = len(str_digit)
    i = 2
    while i <= n:
        count_curr = 0

        if str_digit[i-1] > '0':
            count_curr = count_last

        # check for valid character, that is between 1 to 26
        if (str_digit[i-2] == '1' or (str_digit[i-2] == '2' and str_digit[i-1] < '7')):
            count_curr += count_2nd_last

        count_2nd_last, count_last = count_last, count_curr

        i += 1

    return count_curr


if __name__ == '__main__':
    filename = "test_cases.txt"
    with open(filename) as file:
        content = file.readline()
        while content:
            str_digit, ans = content.split()
            ans = int(ans)

            if(decode_count(str_digit) == ans):
                print("Passed")
            else:
                print("Failed")
            content = file.readline()
