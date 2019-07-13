package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
    //"unicode/utf8"
)

func main() {
    bookmarkGo()
}


func bookmarkGo(){
    displayList()
    fmt.Print("どのテキストを選択しますか")
    var i int
    fmt.Scan(&i)
    choseOpenVim(i)
}


func choseOpenVim(i int){
    ret := getBookmarkList()
    slice := strings.Split(ret[i], ":")

    // TODO 指定したindexがない場合の処理
    cmd := exec.Command("vim", slice[1])
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}


func getBookmarkList()[]string{
    // TODO 絶対パスを入れておかないとうまく動かない
    out, err := exec.Command("cat", "bookmark.txt").Output()
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    ret := strings.Split(string(out), "\n")
    //ret = ret[:len(ret)-1] // 不要な空行が入るので削除
    return ret
}


func displayList(){
    ret := getBookmarkList()
    for i, s := range ret {
        if len(s) > 70{
            tmp := strings.Split(ret[i], ":")
            tmp2 :=tmp[1]
            fmt.Printf("%s ",tmp[0])
            fmt.Printf(":...%s \n",tmp2[len(tmp2)-70:])

        }else{
            fmt.Printf("%d: %s\n", i, s)
        }
    }
}

