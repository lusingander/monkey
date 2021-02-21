#
# calc fibonacchi
#

let fib = fn(n) {
  if (n <= 1) {
    return 0
  }
  if (n == 2) {
    return 1
  }
  return fib(n - 1) + fib(n - 2) # TODO: memorize
};

let n = 30
let result = fib(n)

println(" in:", n)
println("out:", result)
