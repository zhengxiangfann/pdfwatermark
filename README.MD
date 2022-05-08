


###### 引用-pdfcpu

```pdfcpu/stamp.go
 972 行

    // ll := lowerLeftCorner(wm.vp.Width(), wm.vp.Height(), wm.bb.Width(), wm.bb.Height(), wm.Pos)
	//dx := ll.X + wm.bb.Width()/2 + float64(wm.Dx) + sin*(wm.bb.Height()/2+dy) - cos*wm.bb.Width()/2
	//dy = ll.Y + wm.bb.Height()/2 + float64(wm.Dy) - cos*(wm.bb.Height()/2+dy) - sin*wm.bb.Width()/2
        dx := +wm.vp.Width() * sin * sin
        dy =  -wm.vp.Width() * sin * cos
```

