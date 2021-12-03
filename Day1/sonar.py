import sys

if __name__ == '__main__':
	num = 0
	idx = 0
	for line in sys.stdin: 
		cur = int(line.rstrip())
		if idx > 0 and cur > last:
			num += 1
		last = cur
		idx += 1
	print(num)