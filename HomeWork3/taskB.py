def check_new_city(cities):
    tmp_stack = []
    for index, city in enumerate(cities):
        if len(tmp_stack) == 0:
            tmp_stack.append((index, city))
        else:
            while len(tmp_stack) != 0:
                last = tmp_stack.pop()
                if last[1] > city:
                    cities[last[0]] = index
                    continue
                else:
                    tmp_stack.append(last)
                    break
            tmp_stack.append((index, city))
    for index, city in tmp_stack:
        cities[index] = -1
    return cities

def read_data():
    n = int(input())
    cities = list(map(int, input().split()))
    return cities

def main():
    cities = read_data()
    print(*check_new_city(cities))

if __name__ == '__main__':
    main()