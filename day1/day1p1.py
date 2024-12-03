import functools
with open("input.txt", "r") as file:
    left_list = []
    right_list = []
    for line in file.readlines():
        nums = line.split("   ")
        left_list.append(int(nums[0]))
        right_list.append(int(nums[1]))

    left_list.sort()
    right_list.sort()

    pairs = zip(left_list, right_list)
    total_distance = [abs(pair[1]-pair[0]) for pair in pairs]
    print(total_distance)
    print(sum(total_distance))

