import sys

if __name__ == '__main__':
	num = 0
	ints = []
	for line in sys.stdin: 
		ints.append(int(line.rstrip()))
	tot = len(ints)
	sums = [0 for i in range(tot)]
	for (i, x) in enumerate(ints):
		low = max(0, i-2)
		high = i if (i+2 < tot) else i-1
		t = low
		while t <= high:
			sums[t] += x
			t += 1
	for (i, s) in enumerate(sums):
		if i > 0 and s > last:
			num += 1
		last = s

	print(num)