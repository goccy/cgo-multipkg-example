build: test.o a.o b.o util.o
	gcc -o test test.o a.o b.o util.o

clean:
	rm -f *.o test

test.o:
	gcc -o test.o -c test.c

a.o:
	gcc -o a.o -c a.c

b.o:
	gcc -o b.o -c b.c

util.o:
	gcc -o util.o -c util.c
