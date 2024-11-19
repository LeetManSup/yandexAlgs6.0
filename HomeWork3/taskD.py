def read_data():
    expression = input().strip().split()
    return expression

def calc_prefix(expression):
    tmp_stack = []
    for element in expression:
        if element.isdigit():
            tmp_stack.append(int(element))
        else:
            op2 = tmp_stack.pop()
            op1 = tmp_stack.pop()
            tmp_stack.append(eval(f'{op1}{element}{op2}'))
    return tmp_stack.pop()


def main():
    expression = read_data()
    print(calc_prefix(expression))

if __name__ == '__main__':
    main()