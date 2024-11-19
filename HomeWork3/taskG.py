def read_data():
    n, k = map(int, input().split())
    nums = list(map(int, input().split()))
    return n, k, nums

def calc_total_time(n, k, nums):
    waiting = 0
    total_time = 0
    for i in range(n):
        total_time += nums[i] + waiting
        waiting = max(0, nums[i] + waiting - k)
    total_time += waiting
    return total_time


def main():
    n, k, nums = read_data()
    total_time = calc_total_time(n, k, nums)
    print(total_time)

if __name__ == '__main__':
    main()