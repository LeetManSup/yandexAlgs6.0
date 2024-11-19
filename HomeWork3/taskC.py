from collections import deque


def read_data():
    n, k = map(int, input().split())
    nums = list(map(int, input().split()))
    return n, k, nums

def add_another_num(num, my_deque):
    while len(my_deque) != 0:
        deque_max = my_deque.pop()
        if num < deque_max:
            continue
        else:
            my_deque.append(deque_max)
            break
    my_deque.append(num)

def print_win_min(k, nums):
    tmp_deque = deque()
    for i in range(k):
        add_another_num(nums[i], tmp_deque)
    print(tmp_deque[0])
    for i in range(k, len(nums)):
        if nums[i-k] == tmp_deque[0]:
            tmp_deque.popleft()
        add_another_num(nums[i], tmp_deque)
        print(tmp_deque[0])

def main():
    n, k, nums = read_data()
    print_win_min(k, nums)

if __name__ == '__main__':
    main()