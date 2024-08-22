def parse_input(input_file):
    with open(input_file, 'r') as file:
        return file.read().splitlines()

def day04_of_2022_part_a(pairs):
    fully_contained = 0
    for pair in pairs:
        ranges = pair.split(',')
        range1 = list(map(int, ranges[0].split('-')))
        range2 = list(map(int, ranges[1].split('-')))
        if (range1[0] <= range2[0] and range1[1] >= range2[1]) or (range2[0] <= range1[0] and range2[1] >= range1[1]):
            fully_contained += 1
    return fully_contained

def day04_of_2022_part_b(pairs):
    partially_contained = 0
    for pair in pairs:
        ranges = pair.split(',')
        range1 = list(map(int, ranges[0].split('-')))
        range2 = list(map(int, ranges[1].split('-')))
        if (range1[0] <= range2[0] <= range1[1]) or (range2[0] <= range1[0] <= range2[1]):
            partially_contained += 1
    return partially_contained

def main():
    try:
        pairs = parse_input("input.txt")
        fully_contained = day04_of_2022_part_a(pairs)
        partially_contained = day04_of_2022_part_b(pairs)
        print(fully_contained)
        print(partially_contained)
    except Exception as e:
        print(f"An error occurred: {e}")

if __name__ == "__main__":
    main()
