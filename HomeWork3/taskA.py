def check_brackets(brackets):
    stack = []
    dict = {
        '(': ')',
        '[': ']',
        '{': '}'
    }

    for bracket in brackets:
        if bracket in dict:
            stack.append(bracket)
        else:
            if len(stack) == 0 or dict[stack.pop()] != bracket:
                return 'no'

    if len(stack) == 0:
        return 'yes'
    else:
        return 'no'


def main():
    brackets = input()
    print(check_brackets(brackets))

if __name__ == '__main__':
    main()