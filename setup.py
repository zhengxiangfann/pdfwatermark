# -*- coding:utf-8 -*-
# CreateDate: 2022/5/9 14:17

from distutils.core import setup, Extension
from Cython.Build import cythonize

# 这里我们不能在 sources 里面写上 ["fib.pyx", "libfib.a"]，这是不合法的，因为 sources 里面需要放入源文件
# 静态库和动态库需要通过 library_dirs 和 libraries 指定
ext = Extension(name="wrapper_watermark",
                sources=["watermark.pyx"],
                # 相当于 gcc 的 -L 参数，路径可以指定多个
                library_dirs=["."],
                # 相当于 gcc 的 -l 参数，链接的库可以指定多个
                # 注意：不能写 libfib.a，直接写 fib 就行，所以静态命名需要遵循规范，要以 lib 开头、.a 结尾
                # 动态库同理，lib 开头、.so 结尾
                libraries=["watermark"]
                # 如果还需要头文件的话，那么通过 include_dirs 指定
                # 只不过由于头文件就在当前目录中，所以我们不需要指定
                )

setup(ext_modules=cythonize(ext, language_level=3))
