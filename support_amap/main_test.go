package support_amap

import (
	"bcd-util/util"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"math"
	"strconv"
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	curMileage := 150 * 1000.0
	maxMileage := 300 * 1000.0
	findMileage := 50 * 1000.0
	client := resty.New()
	get, err := client.R().
		SetQueryParam("key", "1cf19b21f25276e4d93f199b42304b03").
		SetQueryParam("origin", "114.404057,30.475014").
		SetQueryParam("destination", "116.379163,39.866435").
		SetQueryParam("origin_type", "0").
		SetQueryParam("cartype", "1").
		SetQueryParam("show_fields", "polyline").
		Get("https://restapi.amap.com/v5/direction/driving?parameters")
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	res := get.String()
	//util.Log.Infof("%s", res)
	parse := gjson.Parse(res)
	index := 0
	var arr [][2]float64
	for _, e1 := range parse.Get("route").Get("paths").Array() {
		for _, e2 := range e1.Get("steps").Array() {
			index++
			split1 := strings.Split(e2.Get("polyline").String(), ";")
			for _, e3 := range split1 {
				split2 := strings.Split(e3, ",")
				lng, err := strconv.ParseFloat(split2[0], 64)
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}
				lat, err := strconv.ParseFloat(split2[1], 64)
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}
				arr = append(arr, [2]float64{lng, lat})
			}
			//split1 := strings.Split(e2.Get("polyline").String(), ";")
			//prev := split1[0]
			//distance := 0.0
			//for i := 1; i < len(split1); i++ {
			//	e3 := split1[i]
			//	arr1 := strings.Split(prev, ",")
			//	arr2 := strings.Split(e3, ",")
			//	f1, _ := strconv.ParseFloat(arr1[0], 64)
			//	f2, _ := strconv.ParseFloat(arr1[1], 64)
			//	f3, _ := strconv.ParseFloat(arr2[0], 64)
			//	f4, _ := strconv.ParseFloat(arr2[1], 64)
			//	distance = distance + Distance(f1, f2, f3, f4)
			//	prev = e3
			//}
			//_, _ = fmt.Printf("%s %f\n", e2.Get("step_distance"), distance)
		}
	}

	prev := arr[0]
	temp := curMileage
	for i := 1; i < len(arr); i++ {
		cur := arr[i]
		temp = temp - Distance(prev[0], prev[1], cur[0], cur[1])
		if temp <= findMileage {
			response, err := client.R().
				SetQueryParam("key", "1cf19b21f25276e4d93f199b42304b03").
				SetQueryParam("types", "011100").
				SetQueryParam("location", fmt.Sprintf("%f,%f", arr[i][0], arr[i][1])).
				SetQueryParam("radius", "50000").Get("https://restapi.amap.com/v5/place/around?parameters")
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}
			fmt.Printf("%.2f\n", temp)
			result := gjson.Parse(response.String())
			pois := result.Get("pois").Array()
			for _, e := range pois[:3] {
				location := e.Get("location").String()
				distance := e.Get("distance").String()
				fmt.Printf("%s,%s\n", location, distance)
			}
			println()
			temp = maxMileage
		}
		prev = cur
	}

}

// R 地球半径，单位米
const R = 6367000

// Distance
// lonA, latA分别为A点的纬度和经度
// lonB, latB分别为B点的纬度和经度
// 返回的距离单位为米
func Distance(lngA, latA, lngB, latB float64) float64 {
	if lngA == lngB && latA == latB {
		return 0
	}
	c := math.Sin(latA)*math.Sin(latB)*math.Cos(lngA-lngB) + math.Cos(latA)*math.Cos(latB)
	return R * math.Acos(c) * math.Pi / 180
}
