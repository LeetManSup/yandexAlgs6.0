def read_data():
    n = int(input())
    operations = []
    for i in range(n):
        operations.append(input())
    return operations

def stack_with_sum(operations):
    stack_values = []
    stack_prefix_sums = []
    result = []

    for op in operations:
        if op[0] == '+':
            x = int(op[1:])
            stack_values.append(x)
            if stack_prefix_sums:
                stack_prefix_sums.append(stack_prefix_sums[-1] + x)
            else:
                stack_prefix_sums.append(x)

        elif op[0] == '-':
            removed_value = stack_values.pop()
            stack_prefix_sums.pop()
            result.append(str(removed_value))

        elif op[0] == '?':
            k = int(op[1:])
            if k == len(stack_prefix_sums):
                result.append(str(stack_prefix_sums[-1]))
            else:
                sum_k = stack_prefix_sums[-1] - stack_prefix_sums[-k - 1]
                result.append(str(sum_k))

    return "\n".join(result)

def main():
    operations = read_data()
    print(stack_with_sum(operations))

if __name__ == '__main__':
    main()