package docx_writer

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gingfrederik/docx"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type DocxWriter struct {
	fileRelPath  string
	fileFullPath string
	lineDuration time.Duration
	originInput  time.Time //input side

	f           *docx.File
	lineIndex   int
	lastStamp   time.Time
	lastCaption string
}

// CaptionResult 一段 格式
type CaptionResult struct {
	Text  string        `json:"text"`
	Start time.Duration `json:"sent_start_ts"`
	End   time.Duration `json:"sent_end_ts"`
}

func NewDocxWriter(dir string, uid string, originInput time.Time, lineDuration int) (*DocxWriter, error) {
	sw := &DocxWriter{}
	sw.originInput = originInput
	sw.lineDuration = time.Duration(lineDuration) * time.Millisecond
	pathParts := strings.Split(dir, "|")
	if len(pathParts) < 2 {
		subDir := originInput.Format("01.2006")
		sw.fileRelPath = fmt.Sprintf("%s%s%s_%s_%v.docx", strings.Trim(subDir, `"'`), string(os.PathSeparator), uid, originInput.Format("02_150405.000"), time.Now().Unix())
	} else {
		dir = pathParts[0]
		sw.fileRelPath = pathParts[1]
	}
	sw.fileFullPath = filepath.Join(strings.Trim(dir, `"'`), sw.fileRelPath)
	err := os.MkdirAll(filepath.Dir(sw.fileFullPath), os.ModePerm)
	if err != nil {
		return nil, err
	}

	sw.f = docx.NewFile()
	logs.Debug("caption will be saved at %s", sw.fileFullPath)
	sw.lineIndex = 1
	return sw, err
}
func GenRelPath(dir string, uid string, originInput time.Time, filetype string) (fileRelPath string) {
	// todo: 这个 path part 可能有多个？
	pathParts := strings.Split(dir, "|")
	if len(pathParts) < 2 {
		subDir := originInput.Format("01.2006")
		fileRelPath = fmt.Sprintf("%s%s%s_%s_%v.%s", subDir, string(os.PathSeparator), uid, originInput.Format("02_150405.000"), time.Now().Unix(), filetype)
	} else {
		dir = pathParts[0]
		fileRelPath = pathParts[1]
	}
	return fileRelPath
}
func (sw *DocxWriter) GetRelPath() string {
	return sw.fileRelPath
}

func (sw *DocxWriter) getTimeString(stamp time.Time) string {
	return sw.getDurString(stamp.Sub(sw.originInput))
}

func (sw *DocxWriter) getDurString(dur time.Duration) string {
	d := dur.Nanoseconds() / 1e6
	milli := d % 1000
	d = d / 1000
	sec := d % 60
	d = d / 60
	min := d % 60
	hour := d / 60
	return fmt.Sprintf("%02d:%02d:%02d,%03d", hour, min, sec, milli)
}

func (sw *DocxWriter) Append(ts time.Time, caption string) (err error) {
	if len(sw.lastCaption) > 0 {
		endts := sw.lastStamp.Add(sw.lineDuration)
		if ts.Before(endts) {
			endts = ts
		}
		sw.f.AddParagraph().AddText(fmt.Sprintf("%d", sw.lineIndex))
		sw.f.AddParagraph().AddText(fmt.Sprintf("%s--> %s", sw.getTimeString(sw.lastStamp), sw.getTimeString(endts)))
		sw.f.AddParagraph().AddText(fmt.Sprintf("%s", sw.lastCaption))
		sw.f.AddParagraph().AddText("")

		sw.lineIndex += 1
	}
	sw.f.Save(sw.fileFullPath)
	sw.lastStamp = ts
	sw.lastCaption = caption
	return
}

func (sw *DocxWriter) AppendCaptions(captions []*CaptionResult) (err error) {
	for _, caption := range captions {
		sw.f.AddParagraph().AddText(fmt.Sprintf("%d", sw.lineIndex))
		sw.f.AddParagraph().AddText(fmt.Sprintf("%s--> %s", sw.getDurString(caption.Start), sw.getDurString(caption.End)))
		sw.f.AddParagraph().AddText(fmt.Sprintf("%s", caption.Text))
		sw.f.AddParagraph().AddText("")
		sw.lineIndex += 1
	}
	return
}

func (sw *DocxWriter) Dispose() {
	sw.Append(sw.lastStamp.Add(sw.lineDuration), "")
	logs.Debug("caption file written complete!")
}
