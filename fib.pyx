cdef extern from "go_fib.h":
    double go_fib(int)


def fib_with_go(n):
    """调用 Go 编写的斐波那契数列，以静态库形式存在"""
    return go_fib(n)
