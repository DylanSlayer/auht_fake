package routers

import (
	"auht_fake/common"
	"auht_fake/model"
	"auht_fake/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func commonRouters(r *gin.Engine) {
	r.GET("/", getMain)
}

func getMain(ctx *gin.Context) {
	ts := time.Now().AddDate(0, 0, -1)
	info := model.LeaveInfo{
		StudentName: utils.GetRandomItem(common.Settings.FakeDataInfo.StudentNames),
		StudentId:   strconv.Itoa(utils.GetNumberBetween(common.Settings.FakeDataInfo.IdBegin, common.Settings.FakeDataInfo.IdEnd)),
		Academy:     utils.GetRandomItem(common.Settings.FakeDataInfo.Academies),
		Telephone:   utils.GetRandomTel(),
		OutTime:     time.Now().Format("2006-01-02") + " 00:00:00",
		ReturnTime:  time.Now().Format("2006-01-02") + " 23:59:59",
		OutAddress:  utils.GetRandomItem(common.Settings.FakeDataInfo.Addresses),
		Reason:      utils.GetRandomItem(common.Settings.FakeDataInfo.Reason),
		ApplyTime:   utils.UnixToStr(time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, ts.Location()).Unix(), "2006-01-02") + " " + utils.GetRandomTime(1),
		Approver:    utils.GetRandomItem(common.Settings.FakeDataInfo.Approver),
		ApproveTime: utils.UnixToStr(time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, ts.Location()).Unix(), "2006-01-02") + " " + utils.GetRandomTime(2),
	}
	ctx.JSON(200, info)
}
