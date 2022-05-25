##### 1. python调用go语言编写的动态库来加快速度



```
libwatermark(go语言加水印)
依赖 pdfcpu

pdfcpu源码修改了
pdfcpu/stamp.go 大概是 972行

//ll := lowerLeftCorner(wm.vp.Width(), wm.vp.Height(), wm.bb.Width(), wm.bb.Height(), wm.Pos)
//dx := ll.X + wm.bb.Width()/2 + float64(wm.Dx) + sin*(wm.bb.Height()/2+dy) - cos*wm.bb.Width()/2
//dy = ll.Y + wm.bb.Height()/2 + float64(wm.Dy) - cos*(wm.bb.Height()/2+dy) - sin*wm.bb.Width()/2
dx := +wm.vp.Width() * sin * sin
dy =  -wm.vp.Width() * sin * cos

```

###### 2.编译为动态库

```
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s"  -buildmode=c-shared -o watermark.so watermark.go
chmod 0777 watermark.so
cp watermark.so app/

通过 app/watermarkso.py 加载动态库

```

##### 3. 或者编译为静态库并且打包成python 的扩展包
```
linux规范 必须lib开头

CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s"  -buildmode=c-archive -o libwatermark.a watermark.go

编写 setup.py 和 watermark.pyx 

python3.8 setup.py build

将生成的 wrapper_watermark.cpython-38-darwin.so 拷贝到 app/

import wrapper_watermark


```



