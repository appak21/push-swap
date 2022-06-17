# push-swap

### [Full description](https://github.com/appak21/01-edu/tree/master/subjects/push-swap)

Algorithm: [Radix Sort](https://en.wikipedia.org/wiki/Radix_sort)

### Usage: push-swap
> As an argument, enter numbers or use -rand flag to use randomly generated numbers

> Each result will be saved in `push-swap-result.txt` file. So if you want to test big numbers using `-rand` flag, you can see what numbers _random number generator_ gave and what instructions were used to sort them. It helps you easily copy/paste for **checker** program.
```console
$ go run . "2 1 3 6 5 8"
pb
pb
ra
sa
rrr
pa
pa
$ go run . "0 one 2 3"
Error
$ go run . -rand
Numbers will be randomly generated in the range [min,max).
Enter the minimum, maximum values and the length of numbers in order.
min max length
-100 100 5
ra
pb
ra
pa
$
```
### Usage: checker
```console
$ echo -e "rra\npb\nsa\n" | go run ./checker "3 2 one 0"
Error
$ echo -e "rra\npb\nsa\nrra\npa"
rra
pb
sa
rra
pa
$ echo -e "rra\npb\nsa\nrra\npa" | go run ./checker "3 2 1 0"
OK
$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | go run ./checker "$ARG"
OK
```
[Algorithm explanation](https://medium.com/nerd-for-tech/push-swap-tutorial-fa746e6aba1e)
