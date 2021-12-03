import sys

def str_to_intarray(s):
    return [int(c) for c in s]

def add_two_intarray(f, s):
    return list(map(lambda x,y: x+y, f, s))

if __name__ == '__main__':
    gamma, epsilon = 0, 0
    sums = []
    nr, nd = 0, 0
    for line in sys.stdin:
        inp = str_to_intarray(line.rstrip())
        if nr == 0:
            sums = inp
            nd = len(sums)
        else:
            sums = add_two_intarray(inp, sums)
        nr += 1
    for i, x in enumerate(sums):
        if x < nr/2:
            epsilon += 2**(nd - i - 1)
        else:
            gamma += 2**(nd - i - 1)
    print(gamma*epsilon)