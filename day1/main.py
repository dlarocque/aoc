

def main(input):
    cals = []
    curr_cals = 0
    for line in input:
        if line == '\n':
            # next elf
            cals.append(curr_cals)
            curr_cals = 0
        else:
            curr_cals += int(line)

    return sum(sorted(cals)[-3:])


if __name__ == '__main__':
    with open('input.txt') as f:
        lines = f.readlines()
        print(lines)
        print('answer: ', main(lines))