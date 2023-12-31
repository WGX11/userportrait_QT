package label_gene

import (
	"backend/biz/entity/event_data"
	"backend/biz/entity/rule"
	"backend/biz/util"
	"backend/consumer/common"
	"context"
	"fmt"
	"github.com/bytedance/gopkg/util/logger"
	"math"
	"strconv"
	"strings"
)

/*
processEventCntLabel函数作用
 1. 读取用户的事件数据
 2. 根据labelId，处理事件数据，得到对应的label
 3. 返回结果
*/
func processEventCntLabel(ctx context.Context, appId int64, labelId int64) map[int64]string {
	res := make(map[int64]string)
	// 数据文件路径
	userEventPath := common.GetUserEventPath(ctx, appId)
	if len(userEventPath) == 0 {
		return res
	}

	// 处理数据
	userIds := make([]int64, 0, len(userEventPath))
	plMap := make(map[int64]string)         // 编程语言 u_id -> program language
	codeSpeedMap := make(map[int64]float64) // 打字速度 u_id -> code speed
	shortcutCntMap := make(map[int64]int64) // 快捷键次数 u_id -> cnt
	visitCntMap := make(map[int64]string)
	editMap := make(map[int64]string) // 是否偏好编辑

	// git 操作次数 u_id -> cnt
	for userId, paths := range userEventPath {
		userIds = append(userIds, userId)
		cCnt, cppCnt := int64(0), int64(0)
		keyClickCnt, keyClickDuration := int64(0), int64(0)
		shortcutCnt := int64(0)
		preferEdit := bool(false)
		visitCnt := make(map[string]int64)
		for _, path := range paths {
			events, err := common.OpenFile(path)
			if err != nil {
				logger.Error("open file failed. err=", err.Error())
				continue
			}
			switch labelId {
			case ProgramLanguage:
				c, cpp := processProgramLanguage(events)
				cCnt = cppCnt + c
				cppCnt = cppCnt + cpp
			case CodeSpeed:
				cnt, duration := processCodeSpeed(events)
				keyClickCnt = keyClickCnt + cnt
				keyClickDuration = keyClickDuration + duration
			case ShortcutFre:
				shortcutCnt = shortcutCnt + processShortcutCnt(events)
			case MostVisit:
				processVisitCnt(events, visitCnt)
			case UsePerfer:
				if processUserPerfer(events) {
					preferEdit = true
				}
			default:
			}
		}
		switch labelId {
		case ProgramLanguage:
			if cCnt > cppCnt {
				plMap[userId] = "1"
			} else {
				plMap[userId] = "2"
			}
		case UsePerfer:
			if preferEdit {
				editMap[userId] = "1"
			} else {
				editMap[userId] = "2"
			}
		case CodeSpeed:
			codeSpeedMap[userId] = float64(keyClickCnt) / float64(keyClickDuration)
		case ShortcutFre:
			shortcutCntMap[userId] = shortcutCnt
		case MostVisit:
			mostName := ""
			maxValue := int64(math.MinInt64)
			for key, value := range visitCnt {
				if value > maxValue {
					maxValue = value
					mostName = key
				}
			}
			visitCntMap[userId] = mostName
		default:

		}
	}

	// 结果
	switch labelId {
	case ProgramLanguage:
		res = plMap
	case UsePerfer:
		res = editMap
	case CodeSpeed:
		gradeMap := util.GradeByPercent(codeSpeedMap, []float64{0.3, 0.7})
		for userId, grade := range gradeMap {
			res[userId] = fmt.Sprintf("%d", grade)
		}
	case ShortcutFre:
		gradeMap := util.GradeByPercent(util.ConvertIntMap2Float(shortcutCntMap), []float64{0.3, 0.7})
		for userId, grade := range gradeMap {
			res[userId] = fmt.Sprintf("%d", grade)
		}
	case MostVisit:
		res = visitCntMap
	default:
	}

	return res
}

func processGitCnt(events [][]string) int64 {
	cnt := int64(0)
	// git操作
	value := []string{
		"(3|1|1|||MainWindow.menubar.menuGit)",
		"(3|1|1|||MainWindow.<class_name=QMenu>.actionGit_Create_Repository)",
		"(3|1|1|||MainWindow.<class_name=QMenu,1>.actionGit)",
	}
	for _, event := range events {
		if rule.MatchEvent(event, value) {
			cnt++
		}
	}
	return cnt
}

func processVisitCnt(events [][]string, visitCnt map[string]int64) {
	for _, event := range events {
		if event_data.ComponentNameIndex > len(event)-1 {
			continue
		}

		extra := event[event_data.ComponentNameIndex]
		componentNames := strings.Split(extra, ".")
		lastPart := ""
		var lenth = len(componentNames)
		if lenth > 1 {
			lastPart += componentNames[lenth-2] + "."
		}
		lastPart += componentNames[len(componentNames)-1]
		if len(lastPart) > 0 {
			visitCnt[lastPart]++
		}

	}
}

func processShortcutCnt(events [][]string) int64 {
	cnt := int64(0)
	for _, event := range events {
		if event_data.EventTypeIndex >= len(events) {
			continue
		}
		if event[event_data.EventTypeIndex] == string(event_data.Shortcut) {
			cnt++
		}
	}

	return cnt
}

func processCodeSpeed(events [][]string) (keyClickCnt int64, duration int64) {
	keyClickCnt = 0
	duration = 0
	// 代码区键盘输入
	value := []string{
		"(5|0|0|1||MainWindow.centralwidget.EditorPanel.splitterEditorPanel.EditorTabs)",
		"(5|0|0|2||MainWindow.centralwidget.EditorPanel.splitterEditorPanel.EditorTabs)",
	}

	match := false // 上次事件是否为键盘输入
	lastTime := int64(0)
	for i := 1; i < len(events); i++ {
		event := events[i]
		timeStr := event[event_data.EventTimeIndex]
		timeStamp, err := strconv.ParseInt(timeStr, 10, 64)
		if err != nil {
			logger.Error("event time parse failed. err=", err.Error())
			continue
		}
		if rule.MatchEvent(event, value) {
			if match == true { // 计算次数 & 时间
				keyClickCnt++
				duration = duration + timeStamp - lastTime
			}
			match = true
		} else {
			match = false
		}

		lastTime = timeStamp
	}

	return keyClickCnt, duration
}

func processProgramLanguage(events [][]string) (cCnt int64, cppCnt int64) {
	cMap := make(map[string]bool)
	cppMap := make(map[string]bool)
	for _, event := range events {
		if event_data.ComponentExtraIndex > len(event)-1 {
			continue
		}

		extra := event[event_data.ComponentExtraIndex]
		if strings.HasSuffix(extra, ".c") {
			cMap[extra] = true
		} else if strings.HasSuffix(extra, ".cpp") {
			cppMap[extra] = true
		}
	}

	return int64(len(cMap)), int64(len(cppMap))
}

func processUserPerfer(events [][]string) bool {
	cnt := int64(0)
	for _, event := range events {
		if event_data.KeyClickTypeIndex > len(event)-1 {
			continue
		}

		extra := event[event_data.KeyClickTypeIndex]
		if len(extra) > 0 {
			cnt++
		}
	}
	return cnt > 20
}
