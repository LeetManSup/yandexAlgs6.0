def read_data():
    n = int(input())
    w = input()
    s = input()
    return n, w, s


def calc_min_brackets(n, w, s):
    brackets = {'(': ')', '[': ']'}
    result = s
    tmp_stack = []

    for c in s:
        if c in brackets:
            tmp_stack.append(c)
        else:
            tmp_stack.pop()

    def find_open_bracket():
        for b in w:
            if b in brackets:
                return b

    def find_matching_bracket():
        for b in w:
            if b == brackets.get(tmp_stack[-1], '') or b in brackets:
                return b

    def find_closing_bracket():
        for b in w:
            if b == brackets.get(tmp_stack[-1], ''):
                return b

    for _ in range(len(s), n):
        if not tmp_stack:
            new_bracket = find_open_bracket()
            tmp_stack.append(new_bracket)
            result += new_bracket
        elif n - len(result) > len(tmp_stack):
            new_bracket = find_matching_bracket()
            if new_bracket in brackets:
                tmp_stack.append(new_bracket)
            else:
                tmp_stack.pop()
            result += new_bracket
        else:
            result += find_closing_bracket()
            tmp_stack.pop()

    return result


def main():
    n, w, s = read_data()
    print(calc_min_brackets(n, w, s))

if __name__ == '__main__':
    main()