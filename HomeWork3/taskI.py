def read_data():
    n = int(input())
    a, b = map(int, input().split())
    rovers = []
    for _ in range(n):
        d, t = map(int, input().split())
        rover = {
            'direction': d,
            'arrival_time': t,
            'road_type': 'main' if d == a or d == b else 'secondary',
            'passed': False,
            'cross_time': None
        }
        rovers.append(rover)
    return a, b, rovers

def print_rovers_time(a, b, rovers):
    # ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ:
    # Получить сторону справа от стороны d
    def get_right_side(d):
        return 4 if d == 1 else d - 1

    # Проверить, должен ли первый ровер уступить второму
    def must_give_way(rover_1, rover_2):
        if rover_1['road_type'] == 'secondary' and rover_2['road_type'] == 'main':
            return True
        elif rover_1['road_type'] == 'main' and rover_2['road_type'] == 'secondary':
            return False
        elif rover_1['road_type'] == rover_2['road_type']:
            return rover_2['direction'] == get_right_side(rover_1['direction'])
        return False


    # ПОДГОТОВКА ОСНОВНЫХ ДАННЫХ:
    n = len(rovers)                 # Количество роверов
    crossroad = {                   # Перекрёсток с очередями 1-4
        1: [],
        2: [],
        3: [],
        4: []
    }

    # Заполнение очередей на перекрёстке парами (время_прибытия, индекс_ровера)
    for index, rover in enumerate(rovers):
        crossroad[rover['direction']].append((rover['arrival_time'], index))

    # Сортировка каждой очереди на перекрёстке по времени прибытия
    for queue in crossroad.values():
        queue.sort()

    time = 1                # Текущее время
    rovers_remaining = n    # Сколько роверов не проехало


    # МОДЕЛЬ ДВИЖЕНИЯ НА ПЕРЕКРЁСТКЕ:
    # Пока есть не проехавшие роверы
    while rovers_remaining > 0:
        ready_rovers = []   # Индексы роверов, которые готовы проехать

        # Определение роверов, которые готовы проехать
        for side in range(1, 5):
            queue = crossroad[side] # Текущая рассматриваемая очередь

            # Пока очередь не пуста и ближайший в текущей очереди ровер уже прошёл перекрёсток
            while queue and rovers[queue[0][1]]['passed']:
                queue.pop(0)    # Убираем ближайший в текущей очереди ровер из этой очереди

            # Если очередь не пуста и ближайший в ней ровер успел прибыть на перекрёсток к текущему времени
            if queue and queue[0][0] <= time:
                index = queue[0][1]         # Индекс ближайшего в очереди ровера, который успел прибыть
                ready_rovers.append(index)  # Делаем такой ровер готовым проехать перекрёсток

        can_go = [True] * len(ready_rovers) # Значения возможности проехать для каждого из готовых роверов

        # Проверяем, какие роверы могут проехать. Для каждого готового ровера
        for i in range(len(ready_rovers)):
            first_rover = rovers[ready_rovers[i]]   # Ровер, для которого узнаём, должен ли он уступить другому роверу
            # Перебираем остальные готовые роверы
            for j in range(len(ready_rovers)):
                # Не сочетаем готовый ровер с самим собой
                if i == j:
                    continue

                second_rover = rovers[ready_rovers[j]]  # Ровер, с которым узнаём, должен ли ему уступить первый ровер

                # Если первый ровер должен уступить второму
                if must_give_way(first_rover, second_rover):
                    can_go[i] = False   # Указываем, что первый ровер не может сейчас проехать
                    break # Завершаем проверку для текущего i-го ровера

        # Обновим статусы роверов. Для каждого индекса среди готовых роверов
        for i, index in enumerate(ready_rovers):
            # Если текущий рассматриваемый ровер может проехать
            if can_go[i]:
                rovers[index]['passed'] = True      # Он проезжает и мы указываем по его индексу, что он проехал
                rovers[index]['cross_time'] = time  # Для проехавшего ровера запоминаем текущее время как время проезда
                rovers_remaining -= 1               # Уменьшаем количество не проехавших роверов на 1

        # Увеличиваем счётчик текущего времени
        time += 1


    # ВЫВОД ВРЕМЕНИ ПРОЕЗДА ДЛЯ КАЖДОГО РОВЕРА
    for rover in rovers:
        print(rover['cross_time'])


def main():
    a, b, rovers = read_data()
    print_rovers_time(a, b, rovers)


if __name__ == '__main__':
    main()