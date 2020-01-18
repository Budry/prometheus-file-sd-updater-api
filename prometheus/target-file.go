package prometheus

import (
	"encoding/json"
	"io/ioutil"
)

type Item struct {
	Targets []string `json:"targets"`
}

type TargetFile struct {
	Path string
}

func NewTargetFile(path string) *TargetFile {
	targetFile := &TargetFile{Path: path}

	return targetFile
}

func (targetFile *TargetFile) Append(hostname string) {
	data, err := ioutil.ReadFile(targetFile.Path)
	if err != nil {
		panic(err)
	}

	var items []Item
	err = json.Unmarshal(data, &items)
	if err != nil {
		panic(err)
	}

	if !contains(items[0].Targets, hostname) {
		items[0].Targets = append(items[0].Targets, hostname)
		newContent, _ := json.MarshalIndent(items, "", "   ")
		err = ioutil.WriteFile(targetFile.Path, newContent, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func (targetFile *TargetFile) Remove(hostname string) {
	data, err := ioutil.ReadFile(targetFile.Path)
	if err != nil {
		panic(err)
	}

	var items []Item
	err = json.Unmarshal(data, &items)
	if err != nil {
		panic(err)
	}

	needUpdate := false
	for i, v := range items[0].Targets {
		if v == hostname {
			items[0].Targets = append(items[0].Targets[:i], items[0].Targets[i+1:]...)
			needUpdate = true
			break
		}
	}

	if needUpdate {
		newContent, _ := json.MarshalIndent(items, "", "  ")
		err = ioutil.WriteFile(targetFile.Path, newContent, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func contains(slice []string, element string) bool {
	for _, actual := range slice {
		if actual == element {
			return true
		}
	}
	return false
}