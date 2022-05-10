<<<<<<< HEAD
cdef extern from "go_watermark.h":
    char* AddWaterMark(char*, char*)
    void ReleaseMemory()


def add_water_mark(f, text):
    return AddWaterMark(f, text)

def release_memory():
    ReleaseMemory()
=======
# -*- coding:utf-8 -*-
# CreateDate: 2022/5/9 15:27

cdef extern from "go_watermark.h":
    char * AddWaterMark(char *, char *)
    void ReleaseMemory()

def add_water_mark(file, text):
    return AddWaterMark(file, text)

def release_memory():
    ReleaseMemory()
>>>>>>> 78f1458 (test)
