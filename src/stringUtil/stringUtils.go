package stringUtil

import (
	"strings"
	"unicode"
)
const period string ="."
const space string =""
//判断字符串是否由数字组成
func IsNum(src string) (ok bool) {
	if strings.Contains(period, src) {
		if strings.Count(period, src) != 1 {
			return false
		}
	} else {
		for _, c := range src {
			if c != 46 {
				if !unicode.IsDigit(c) {
					return false
				}
			}
		}
	}
	return true
}

func Split(s,sep string) (dist []string){
	  tmp:= strings.Split(s, sep)
	  dist=make([]string,len(tmp))
	  i:=0
	  for _,s :=range tmp{
	  ts:=strings.TrimSpace(s)
	  	if ts!=space{
	  		dist[i]=s
	  		i++
	  	}
	  }
	  return dist
}

