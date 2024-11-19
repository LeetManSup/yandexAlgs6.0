from collections import deque


def read_data():
    n, H = map(int, input().split())
    h = list(map(int, input().split()))
    w = list(map(int, input().split()))
    chairs = sorted(map(tuple, zip(h, w)), key=lambda x: (x[0], x[1]))
    return n, H, chairs

def get_min_inconvenience(n, H, chairs):
    if n == 1 or chairs[-1][1] >= H:
        return 0

    h_differences_deque = deque()
    result = deque()
    l = 0
    r = l
    bed_width = chairs[l][1]
    if bed_width >= H:
        return 0
    while l < n:
        while bed_width < H and r < n - 1:
            r += 1

            bed_width += chairs[r][1]
            if chairs[r][1] >= H:
                return 0
            h_difference = chairs[r][0] - chairs[r - 1][0]
            while len(h_differences_deque) != 0:
                max_difference = h_differences_deque.pop()
                if h_difference > max_difference:
                    continue
                else:
                    h_differences_deque.append(max_difference)
                    break
            h_differences_deque.append(h_difference)

        if len(result) != 0 and len(h_differences_deque) == 0:
            break
        if len(result) == 0 or result[0] > h_differences_deque[0] and bed_width >= H:
            result.appendleft(h_differences_deque[0])

        bed_width -= chairs[l][1]
        if abs(chairs[l+1][0] - chairs[l][0]) == h_differences_deque[0]:
            h_differences_deque.popleft()

        l += 1

    return result[0]

def main():
    n, H, chairs = read_data()
    print(get_min_inconvenience(n, H, chairs))

if __name__ == '__main__':
    main()