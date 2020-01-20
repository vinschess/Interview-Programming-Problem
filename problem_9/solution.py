def find_filepath_len(main_stack, file_name):
    length = len(file_name)
    for val in main_stack:
        length += len(val) + 1
    return length


def new_line_sep(str_dir, main_stack, tree_level):
    values = str_dir.split('\\n', 1)
    length = 0
    if '.' in values[0]:
        length = find_filepath_len(main_stack, values[0])
    else:
        while len(main_stack) != tree_level:
            main_stack.pop()
        main_stack.append(values[0])
    return values[1], length


def tab_sep(str_dir, main_stack):
    values = str_dir.split('\\t', 1)
    return values[1]


def longest_absolute_filepath(str_dir):
    str_dir += '\\n'
    main_stack = []
    max_abs_filepath = 0
    tree_level = 0
    length = 0
    while(len(str_dir) > 0):
        length = 0
        if str_dir.find('\\n') != -1 and str_dir.find('\\t') != -1:
            if str_dir.find('\\n') < str_dir.find('\\t'):
                str_dir, length = new_line_sep(str_dir, main_stack, tree_level)
                tree_level = 0
            else:
                str_dir = tab_sep(str_dir, main_stack)
                tree_level += 1
        else:
            if str_dir.find('\\n') != -1:
                str_dir, length = new_line_sep(str_dir, main_stack, tree_level)
                tree_level = 0
            elif str_dir.find('\\t') != -1:
                str_dir = tab_sep(str_dir, main_stack)
                tree_level += 1
        max_abs_filepath = max(max_abs_filepath, length)
    return max_abs_filepath


if __name__ == '__main__':
    filename = "test_cases.txt"
    lines = list()
    with open(filename) as file:
        lines = file.readlines()

    for line in lines:
        string, ans = line.split()
        if longest_absolute_filepath(string) == int(ans):
            print("Passed")
        else:
            print("Failed")
