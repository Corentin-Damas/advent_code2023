def count(puzzle, nums):
    if puzzle == "":
        return 1 if nums == () else 0
    if nums == ():
        return 0 if "#" in puzzle else 1

    result = 0

    if puzzle[0] in ".?":
        result += count(puzzle[1:], nums)

    if puzzle[0] in "#\?":
        if nums[0] <= len(puzzle) and "." not in puzzle[:nums[0]] and (nums[0] == len(puzzle) or puzzle[nums[0]] != "#"):
            result += count(puzzle[nums[0] + 1:], nums[1:])
        else:
            result += 0
    return result


total = 0

for line in open("Day12/data.txt"):
    puzzle, nums = line.split()
    print(line)
    nums = tuple(map(int, nums.split(",")))
    total += count(puzzle, nums)

print(total)
