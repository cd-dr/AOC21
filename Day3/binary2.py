import sys

def str_to_intarray(s):
    return [int(c) for c in s]

class Node:
    def __init__(self, digit):
        self.digit = digit
        self.count = (0, 0)
        self.zero = None
        self.one = None

def update_tree(root, digits):
    cur = root
    for d in digits:
        if d == 0:
            if cur.zero == None:
                cur.zero = Node(0)
            cur.count = (cur.count[0] + 1, cur.count[1])
            cur = cur.zero
        else:
            if cur.one == None:
                cur.one = Node(1)
            cur.count = (cur.count[0], cur.count[1] + 1)
            cur = cur.one

def get_oxygen(root):
    cur = root
    oxy = []
    while cur.count != (0, 0):
        if cur.count[0] <= cur.count[1]:
            oxy.append(1)
            cur = cur.one
        else:
            oxy.append(0)
            cur = cur.zero
    nd = len(oxy)
    oxynum = 0
    for i, d in enumerate(oxy):
        oxynum += (d * 2 ** (nd - i - 1))
    return oxynum

def get_co2(root):
    cur = root
    co2 = []
    while cur.count != (0, 0):
        if cur.count[0] == 0:
            co2.append(1)
            cur = cur.one
        elif cur.count[1] == 0:
            co2.append(0)
            cur = cur.zero
        else: 
            if cur.count[0] <= cur.count[1]:
                co2.append(0)
                cur = cur.zero
            else:
                co2.append(1)
                cur = cur.one
    nd = len(co2)
    co2num = 0
    for i, d in enumerate(co2):
        co2num += (d * 2 ** (nd - i - 1))
    return co2num


if __name__ == '__main__':
    root = Node(-1)
    for line in sys.stdin:
        digits = str_to_intarray(line.rstrip())
        update_tree(root, digits)
    oxy = get_oxygen(root)
    co2 = get_co2(root)

    # print(oxy)
    # print(co2)

    print(oxy*co2)
    