package main

import "sort"

var (
	personFatRage  = map[string]float64{}
)

func inputRecord(name string, fatRate float64) {
	personFatRage[name] = fatRate
}

func getRank(name string) (rank int,fatRate float64 ) {
	fatRate2PersonMap := map[float64]string{}
	rankArr := make([]float64,0,len(personFatRage))
	for nameItem, frItem := range personFatRage {
		fatRate2PersonMap[frItem] = nameItem
		rankArr = append(rankArr,frItem)
	}
	sort.Float64s(rankArr)
	for i,frItem := range rankArr {
		_name := fatRate2PersonMap[frItem]
		if _name == name {
			rank = i +1
			fatRate =frItem
			return
		}
	}
	return 0,0
}