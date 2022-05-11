# -*- coding:utf-8 -*-


from distutils.core import setup, Extension
from Cython.Build import cythonize
ext = Extension(
name="wrapper_watermark",
                sources=["watermark.pyx"],
                library_dirs=["."],
                libraries=["watermark"]
                )

setup(ext_modules=cythonize(ext, language_level=3))
