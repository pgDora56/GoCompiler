package main

import (
    "fmt"
    "os"
    "io/ioutil"
)

func main(){
    fmt.Println("Let's make 417!")

    f, err := os.Open("D:\\GoLang\\Compiler\\Shiina\\tes.417")

    if err != nil {
        fmt.Println(err)
        fmt.Println("error")
    }

    defer f.Close() // 遅延実行: プログラムの最後に実行する

    // ======= Read 1: Not use ioutil

    // var buf []byte = make([]byte, 1024) // バイト型スライスの作成
    // // buf := make([]byte, 1024} // これも同義
    //
    // i := 1
    // for {
    //     n, err := f.Read(buf) // とりあえず1024byte分読めるだけ読む
    //
    //     if n == 0 {
    //         break
    //     }
    //     if err != nil {
    //         break
    //     }
    //
    //     fmt.Println(string(buf[:n]))
    // }


    // ======= Read 2: Use ioutil
    // こっちのほうがかんたんだ！
    b, err := ioutil.ReadAll(f)
    fmt.Println(string(b))
}
