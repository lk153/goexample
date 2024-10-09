package relationship

import (
	"strings"
)

func GetMinStep(relationships []string, n1, n2 string) int {
	ppls := make(map[string][]string, 0)
	for _, rels := range relationships {
		tmp := strings.Split(rels, ":")
		if tmp[0] == tmp[1] {
			continue
		}

		ppls[tmp[0]] = append(ppls[tmp[0]], tmp[1])
		ppls[tmp[1]] = append(ppls[tmp[1]], tmp[0])
	}

	//Begin flow
	mapChecked := map[string]bool{}
	checked := 0
	finalPath := ""

	minPath, _ := findPath(n1, ppls, n2, mapChecked, checked, finalPath)
	return minPath
}

func findPath(target string, ppls map[string][]string, finalName string, mapChecked map[string]bool, checked int, finalPath string) (int, string) {
	rels, ok := ppls[target]
	if !ok {
		return 0, finalPath
	}

	mapChecked[target] = true
	checked++
	minPath := 0
	tmpPath := ""
	for _, rel := range rels {
		if mapChecked[rel] {
			continue
		}

		if rel == finalName {
			return checked, "-" + target + "-" + finalName
		}

		v, path := findPath(rel, ppls, finalName, mapChecked, checked, finalPath)
		if v == 0 {
			continue
		}

		if minPath == 0 || v < minPath {
			minPath = v
			tmpPath = "-" + target + path
		}
	}

	return minPath, finalPath + tmpPath
}
