package main

import (
	"fmt"
	"collectDataset/packages"
	"time"
)

func main() {
	base := "https://libraries.io/api/pypi"
	var resStrct []packages.Combined

	jdata := packages.LoadJson("jsons/PkgListSortedbySourcerank.json")

	for _, name := range jdata {
		// URLの設定
		url := fmt.Sprintf("%s/%s", base, name.PkgName)
		resdata := packages.CallApi(url)
		comb := packages.Combined {
			Pypilib: packages.Pypilib {
				SourceRank: name.SourceRank,
				PkgName: name.PkgName,
			},
			ApiRes: packages.ApiRes {
				Repository_url: resdata.Repository_url,
				Forks: resdata.Forks,
				Stars: resdata.Stars,
				Latest_release_published_at: resdata.Latest_release_published_at,
			},
		}
		fmt.Println(comb)
		resStrct = append(resStrct, comb)

		// sleep設定
		time.Sleep(1*time.Second)
	}

	// packages.WriteOutJson("./out.json", &resStrct)
	fmt.Println(resStrct)
}
