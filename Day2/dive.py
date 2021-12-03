import sys

if __name__ == '__main__':
	x = 0
	y = 0
	for line in sys.stdin: 
		words = line.rstrip().split(' ')
		m = int(words[1])
		if words[0] == 'forward':
			x += m
		elif words[0] == 'down':
			y += m
		else:
			y -= m

	print(x*y)