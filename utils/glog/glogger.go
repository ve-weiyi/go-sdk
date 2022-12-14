package glog

import (
	"fmt"
	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
	"gopkg.in/natefinch/lumberjack.v2"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// go使用zap + lumberjack重构项目的日志系统 https://blog.csdn.net/weixin_52000204/article/details/126651319
type Glogger struct {
	//rlog  *lumberjack.Logger
	rlog  interface{}
	log   *zap.Logger        //并重性能与易用性，支持结构化和 printf 风格的日志记录。
	sugar *zap.SugaredLogger // 非常强调性能，不提供 printf 风格的 api（减少了 interface{} 与 反射的性能损耗）

	path       string
	level      zapcore.Level
	rotateMu   *sync.Mutex
	rolling    bool
	lastRotate time.Time
}

func (mlog *Glogger) Logger() *zap.Logger {
	return mlog.log
}

func (mlog *Glogger) checkRotate() {
	if !mlog.rolling {
		return
	}
	jack, ok := mlog.rlog.(*lumberjack.Logger)
	if ok {
		n := time.Now()
		if mlog.differentDay(n) {
			mlog.rotateMu.Lock()
			defer mlog.rotateMu.Unlock()

			// 获得锁之后再次检查是否是不同日期
			// 避免上一次调用已经切割日志,
			if mlog.differentDay(n) {
				jack.Rotate()
				mlog.lastRotate = n
			}
		}
	}
}

// 判断是不是换天了，如果换天了就要重新调用rotate()
func (mlog *Glogger) differentDay(t time.Time) bool {
	y, m, d := mlog.lastRotate.Year(), mlog.lastRotate.Month(), mlog.lastRotate.Day()
	return y != t.Year() || m != t.Month() || d != t.Day()
}

func (mlog *Glogger) EnableDailyFile() {
	mlog.rolling = true
}

func (mlog *Glogger) Error(v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Error(v...)
}

func (mlog *Glogger) Warn(v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Warn(v...)
}

func (mlog *Glogger) Info(v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Info(v...)
}

func (mlog *Glogger) Debug(v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Debug(v...)
}

func (mlog *Glogger) Errorw(format string, v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Errorw(format, v...)
}

func (mlog *Glogger) Warnw(format string, v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Warnw(format, v...)
}

func (mlog *Glogger) Infow(format string, v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Infow(format, v...)
}

func (mlog *Glogger) Debugw(format string, v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Debugw(format, v...)
}

func (mlog *Glogger) Print(v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Info(v...)
}

func (mlog *Glogger) Printf(format string, v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Info(fmt.Sprintf(format, v...))
}

func (mlog *Glogger) GetUnderlyingLogger() *zap.Logger {
	return mlog.log
}

func (mlog *Glogger) Json(v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Info(jsonconv.ObjectToJson(v))
}
