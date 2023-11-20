package staticdata

import (
	"errors"
	"time"

	"github.com/junioryono/Riot-API-Golang/src/constants/patch"
	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

type Patches []patch.Patch

func GetPatches() (Patches, error) {
	var res Patches
	err := getJSON("https://ddragon.leagueoflegends.com/api/versions.json", &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type PatchesWithStartTime []patch.PatchWithStartTime

type patchesWithStartTimeHttpResponse struct {
	Patches []struct {
		Name  string `json:"name"`
		Start int64  `json:"start"`
	} `json:"patches"`
	Shifts map[region.Region]int64 `json:"shifts"`
}

func GetPatchesWithStartTime() (PatchesWithStartTime, error) {
	var res patchesWithStartTimeHttpResponse
	err := getJSON("https://junioryono.github.io/LoLPatches/github-pages/patches.json", &res)
	if err != nil {
		return nil, err
	}

	if len(res.Patches) == 0 {
		return nil, errors.New("no patches found")
	}

	var patches PatchesWithStartTime
	for _, p := range res.Patches {
		patches = append(patches, patch.PatchWithStartTime{
			Patch:     patch.ShortPatch(p.Name),
			StartTime: time.Unix(p.Start, 0),
			Shifts:    res.Shifts,
		})
	}

	return patches, nil
}

func (p Patches) CurrentPatch() patch.Patch {
	return p[0]
}

func (p PatchesWithStartTime) CurrentPatch() patch.PatchWithStartTime {
	return p[len(p)-1]
}
