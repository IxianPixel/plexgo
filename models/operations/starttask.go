// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TaskName - the name of the task to be started.
type TaskName string

const (
	TaskNameBackupDatabase            TaskName = "BackupDatabase"
	TaskNameBuildGracenoteCollections TaskName = "BuildGracenoteCollections"
	TaskNameCheckForUpdates           TaskName = "CheckForUpdates"
	TaskNameCleanOldBundles           TaskName = "CleanOldBundles"
	TaskNameCleanOldCacheFiles        TaskName = "CleanOldCacheFiles"
	TaskNameDeepMediaAnalysis         TaskName = "DeepMediaAnalysis"
	TaskNameGenerateAutoTags          TaskName = "GenerateAutoTags"
	TaskNameGenerateChapterThumbs     TaskName = "GenerateChapterThumbs"
	TaskNameGenerateMediaIndexFiles   TaskName = "GenerateMediaIndexFiles"
	TaskNameOptimizeDatabase          TaskName = "OptimizeDatabase"
	TaskNameRefreshLibraries          TaskName = "RefreshLibraries"
	TaskNameRefreshLocalMedia         TaskName = "RefreshLocalMedia"
	TaskNameRefreshPeriodicMetadata   TaskName = "RefreshPeriodicMetadata"
	TaskNameUpgradeMediaAnalysis      TaskName = "UpgradeMediaAnalysis"
)

func (e TaskName) ToPointer() *TaskName {
	return &e
}

func (e *TaskName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BackupDatabase":
		fallthrough
	case "BuildGracenoteCollections":
		fallthrough
	case "CheckForUpdates":
		fallthrough
	case "CleanOldBundles":
		fallthrough
	case "CleanOldCacheFiles":
		fallthrough
	case "DeepMediaAnalysis":
		fallthrough
	case "GenerateAutoTags":
		fallthrough
	case "GenerateChapterThumbs":
		fallthrough
	case "GenerateMediaIndexFiles":
		fallthrough
	case "OptimizeDatabase":
		fallthrough
	case "RefreshLibraries":
		fallthrough
	case "RefreshLocalMedia":
		fallthrough
	case "RefreshPeriodicMetadata":
		fallthrough
	case "UpgradeMediaAnalysis":
		*e = TaskName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaskName: %v", v)
	}
}

type StartTaskRequest struct {
	// the name of the task to be started.
	TaskName TaskName `pathParam:"style=simple,explode=false,name=taskName"`
}

func (o *StartTaskRequest) GetTaskName() TaskName {
	if o == nil {
		return TaskName("")
	}
	return o.TaskName
}

type StartTaskResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *StartTaskResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *StartTaskResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *StartTaskResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}