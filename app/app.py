# -*- coding:utf-8 -*-
import base64
import logging
import time

import orjson

logging.basicConfig(level=logging.DEBUG)
import gunicorn.app.base
from flask import Flask, send_file
# import orjson
from io import BytesIO
import wrapper_watermark as wm

# from watermarkso import WaterMark
# import gevent.monkey
# gevent.monkey.patch_all()

#AAA local
app = Flask(__name__)


class StandaloneApplication(gunicorn.app.base.BaseApplication):
    def __init__(self, app, options=None):
        self.options = options or {}
        self.application = app
        super(StandaloneApplication, self).__init__()


def load_config(self):
    config = {key: value for key, value in self.options.items()
              if key in self.cfg.settings and value is not None}
    for key, value in config.items():
        self.cfg.set(key.lower(), value)

    def load(self):
        return self.application


@app.route("/")
def index():
    return orjson.dumps({"code": 0, "msg": "ok", "data": {}})


@app.route('/add_water_mark', methods=['GET', 'POST'])
def route_api():
    params_file_path = "../libwatermark/abc.pdf".encode('utf-8')
    params_content = "斗破苍穹 水印 水印 斗罗大陆".encode('utf-8')

    ts = time.time()
    # with_watermark_file = WaterMark().add_water_mask(params_file_path, params_content)
    with_watermark_file = wm.add_water_mark(params_file_path, params_content)
    wm.release_memory()

    print("文件加水印耗时:", time.time() - ts)
    print('watermark_file', len(with_watermark_file))

    filename = "water_mark-abc.pdf"
    with open(filename, "wb") as fw:
        fw.write(with_watermark_file)

    response = send_file(
        BytesIO(base64.b64decode(with_watermark_file)),
        as_attachment=False,
        attachment_filename=filename,
        cache_timeout=1
    )
    response.headers['Content-Disposition'] += "inline; filename*=utf-8''{}".format(filename)
    response.headers['Content-Type'] = "application/pdf"
    return response


if __name__ == '__main__':
    # app.run(host="127.0.0.1", port=5000, debug=True)
    StandaloneApplication(app, {
        'bind': "127.0.0.1:5000",
        'workers': 1,
        'worker_class': 'sync',
        'timeout': 60,
        'loglevel': 'info',
        'logger_class': 'simple'
    }).run()
