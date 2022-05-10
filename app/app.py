# -*- coding:utf-8 -*-
# CreateDate: 2022/5/9 11:37
import logging
logging.basicConfig(level=logging.DEBUG)
import gunicorn.app.base
from flask import Flask
import orjson
import watermark as wm

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


@app.route('/add_water_mark', methods=['GET', 'POST'])
def route_api():
    water_file = wm.add_water_mask("../abc.pdf", "随意 水印 水印 索引")
    wm.release_memory()
    with open("out-water.pdf", "wb") as fw:
        fw.write(water_file)
    return orjson.dumps({"code":0, "msg":"ok"})
    
if __name__ == '__main__':
    # app.run(host="127.0.0.1",port=5000,debug=True)
    StandaloneApplication(app, {
        'bind': "127.0.0.1:5000",
        'workers': 1,
        'worker_class':'sync',
        'timeout': 60,
        'loglevel': 'info',
        'logger_class': 'simple'
    }).run()
