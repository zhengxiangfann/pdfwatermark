# -*- coding:utf-8 -*-
import base64
import ctypes


class WaterMark(object):
    so_path = ""
    so = None

    def __init__(self, pdf_path=None, watermark_text=None):
        self.pdf_path = pdf_path
        self.watermark_text = watermark_text

    @classmethod
    def register_dynamic_link_library(cls, so_path=None):
        if so_path:
            cls.so_path = so_path
        else:
            cls.so_path = "./watermark.so"
        try:
            cls.so = ctypes.CDLL(cls.so_path)
        except FileExistsError as ex:
            print(f"{ex}")
            raise ex

    def _to_byte(self, u_str):
        if isinstance(u_str, str):
            return u_str.encode("utf-8")
        else:
            return u_str

    def add_water_mask(self, pdf_path=None, watermark_text=None):
        self.pdf_path = pdf_path or self.pdf_path
        self.watermark_text = watermark_text or self.watermark_text

        AddWaterMark = self.so.AddWaterMark
        AddWaterMark.argtypes = [ctypes.c_char_p, ctypes.c_char_p]
        AddWaterMark.restype = ctypes.c_char_p
        return base64.b64decode(AddWaterMark(
            self._to_byte(self.pdf_path),
            self._to_byte(self.watermark_text)
        ))


WaterMark.register_dynamic_link_library()

if __name__ == '__main__':
    import time


    t1 = ts =  time.time()
    wm = WaterMark()
    with open("abc-watermark.pdf", "wb") as fw:
        buf = wm.add_water_mask("abc.pdf", " 水印 水印 水印 ")
        fw.write(buf)
    print(time.time() - t1)


    