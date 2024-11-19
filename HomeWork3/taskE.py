import re

def read_data():
    expression = input()
    allowed_chars = set('0123456789+-*() ')
    if not all(char in allowed_chars for char in expression):
        return 'WRONG'

    brackets_stack = []
    for char in expression:
        if char == '(':
            brackets_stack.append(char)
        elif char == ')':
            if not brackets_stack:
                return 'WRONG'
            brackets_stack.pop()
    if brackets_stack:
        return 'WRONG'

    if re.search(r'\d\s+\d', expression):
        return 'WRONG'

    expression = expression.replace(' ', '')

    return expression

def infix_to_prefix(infix):
    priority = {
        '+': 0,
        '-': 0,
        '*': 1
    }
    result = []
    operators = []

    i = 0
    while i < len(infix):
        if infix[i].isdigit():
            num = ''
            while i < len(infix) and infix[i].isdigit():
                num += infix[i]
                i += 1
            result.append(int(num))
            continue
        elif infix[i] in priority:
            while operators and operators[-1] != '(' and priority[operators[-1]] >= priority[infix[i]]:
                result.append(operators.pop())
            operators.append(infix[i])
        elif infix[i] == '(':
            operators.append(infix[i])
        elif infix[i] == ')':
            while operators and operators[-1] != '(':
                result.append(operators.pop())
            operators.pop()
        i += 1

    while operators:
        result.append(operators.pop())

    return result

def calc_prefix(expression):
    try:
        tmp_stack = []
        for element in expression:
            if isinstance(element, int):
                tmp_stack.append(element)
            else:
                op2 = tmp_stack.pop()
                op1 = tmp_stack.pop()
                match element:
                    case '+':
                        tmp_stack.append(op1 + op2)
                    case '-':
                        tmp_stack.append(op1 - op2)
                    case '*':
                        tmp_stack.append(op1 * op2)
        return tmp_stack.pop()
    except:
        return 'WRONG'



def main():
    expression = read_data()
    if expression == 'WRONG':
        print('WRONG')
    else:
        expression = infix_to_prefix(expression)
        print(calc_prefix(expression))



if __name__ == '__main__':
    main()