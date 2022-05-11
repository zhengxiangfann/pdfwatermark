
cdef extern from "go_watermark.h":
    char* AddWaterMark(char*, char*)
    void ReleaseMemory()


def add_water_mark(f, text):
    return AddWaterMark(f, text)

def release_memory():
    ReleaseMemory()

