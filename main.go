/*
* Compare various Image resize algorithms for the Go language
* https://github.com/fawick/speedtest-resize/blob/master/README.md
*
*
* Thumbor
* https://news.ycombinator.com/item?id=7931744
*
* SpeedTest
* https://github.com/fawick/speedtest-resize
*
*
*
*
* https://github.com/nfnt/resize
* install go get github.com/nfnt/resize
*/

package main

import (
    "fmt"
	"github.com/nfnt/resize"
    "os"
    "path/filepath"
    "image/jpeg"    
    "log"    
    )

    func resizer(c string, width uint){
        //for {
            filetoconvert := c

            //Open    
            fmt.Printf("File ->" +  filetoconvert +" \n")
            file, err := os.Open(filetoconvert)
            check(err)

            // decode jpeg into image.Image
            //fmt.Printf("Decode file\n")
            img, err := jpeg.Decode(file)
            check(err)
            file.Close()   
            
            // resize to width 1000 using Lanczos resampling
            // and preserve aspect ratio
            m := resize.Resize(width, 0, img, resize.Lanczos3)

            fileresized := "Resized_"+ filetoconvert+ ".jpg"    
            out, err := os.Create(fileresized)
            fmt.Println("File ->" +  fileresized)
            
            if err != nil {
                log.Fatal(err)
            }
            defer out.Close()

            // write new image to file
            jpeg.Encode(out, m, nil)        
    }

    func GetPngFile() ([]os.FileInfo, error){
        dirname := "." + string(filepath.Separator)

        d, err := os.Open(dirname)
        check(err)
        defer d.Close()

        files, err := d.Readdir(-1)
        check(err)

        slice := make([]os.FileInfo, 0, 1)

        for _, file := range files {
            if file.Mode().IsRegular() {
                if filepath.Ext(file.Name()) == ".png" {
                    slice = append(slice, file)
                }
            }
        }

        size := len(slice)

        fmt.Println("nb file: %d   ",size)

        return slice, nil
    }
 
    func main() {

        files,err := GetPngFile()
        check(err)

        //var c chan string = make(chan string)
        for _, file := range files {        
                go resizer(file.Name(), 100)
        }

        var input string
        fmt.Scanln(&input)
    }

    func check(e error){
        if e != nil {
            log.Fatal(e)
            fmt.Printf("%s", e)
            os.Exit(1)
            panic(e)
        }
    }

