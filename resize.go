// Title：图片缩放
//
// Description:
//
// Author:black
//
// Createtime:2013-10-18 14:16
//
// Version:1.0
//
// 修改历史:版本号 修改日期 修改人 修改说明
//
// 1.0 2013-10-18 14:16 black 创建文档
package commonlib

import (
	"github.com/gosexy/canvas"
)

func Resample(sourceImgPath,distImgPaht string, w, h int) error {

	img := canvas.New()
	err := img.Open(sourceImgPath)

	if err != nil {
		return err
	}

	img.AutoOrientate()
	img.SetQuality(95)
	img.Thumbnail(uint(w), uint(h))
	img.Write(distImgPaht)

	return nil
}
