package main

import (
    "fmt"
    "flag"
    "strconv"
    "math"
	)

func main() {
    // 変数設定
    var num int //判定数字
    var notPrimeFlg bool //素数でないフラグ
    //コマンドライン引数
    flag.Parse()
    line := flag.Args()

    for _,str := range line{
       //素数でないフラグをfalseで初期化
       notPrimeFlg = false
        num, _ = strconv.Atoi(str)
        //カウントアップする最大値を取得
        halfNum := int(math.Ceil(float64(num)/float64(2.0)))
        for i:=2; i<=halfNum;i++{
            if num%i==0{
                notPrimeFlg=true
                break
            }
        }
        if notPrimeFlg{
            fmt.Println(str+"は素数ではないです") 
        }else{
            fmt.Println(str+"は素数です") 
        }
   }
}