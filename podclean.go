package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/annidy/podclean/utils"
)

var podPaths []string
var totalCleanSize float64

func init() {
	podPaths = append(podPaths, utils.Expand("~/Library/Caches/CocoaPods/v1/Pods/Release"))
}

func main() {
	for _, path := range podPaths {
		files, _ := os.ReadDir(path)
		for _, file := range files {
			if !file.IsDir() {
				continue
			}
			locks, _ := filepath.Glob(filepath.Join(path, file.Name(), "*.lock"))
			if len(locks) > 0 {
				var poddirs []*utils.Poddir
				for _, lock := range locks {
					poddir := utils.NewPoddir(strings.TrimSuffix(lock, ".lock"))
					if exist, _ := utils.Exists(poddir.Path); exist {
						poddirs = append(poddirs, poddir)
					}
				}

				sort.Slice(poddirs, func(i, j int) bool {
					return poddirs[i].Version.LessThan(poddirs[j].Version)
				})

				doClean(poddirs)
			}
		}
	}
	fmt.Printf("总共清理了 %.2f MB\n", totalCleanSize)
}

func doClean(poddirs []*utils.Poddir) {
	if len(poddirs) > 1 {
		for _, poddir := range poddirs[0 : len(poddirs)-1] {
			fmt.Printf("删除 %v\n", poddir)
			err := os.RemoveAll(poddir.Path)
			if err != nil {
				log.Fatal(err)
			} else {
				totalCleanSize += poddir.Size()
			}
		}
	}
}
