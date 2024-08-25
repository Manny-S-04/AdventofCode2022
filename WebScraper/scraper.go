package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
    "github.com/joho/godotenv"
)


func main(){

    err := godotenv.Load("variables.env")
    if err != nil {
        fmt.Println("Error loading .env file")
        return
    }

    dayFlag := flag.Int("day", 1, "DayFlag")
    yearFlag := flag.Int("year", 2022, "YearFlag")
    dirToSaveTo := flag.String("dir", "C:/Users/man20/Documents/", "DirToSaveToFlag")
    
    flag.Parse()

    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
    defer cancel()

    baseAddress := "https://adventofcode.com/"
    email := os.Getenv("USER_NAME")
    password := os.Getenv("PASSWORD")

    err = Login(&ctx, baseAddress, email, password)
    if err != nil {
        panic("Failed to login")
    }

    dayString := strconv.Itoa(*dayFlag)
    yearString := strconv.Itoa(*yearFlag)


    fmt.Println(baseAddress + yearString + "/day/" + dayString + "/input")

    var res string
    err = chromedp.Run(ctx,
        chromedp.Navigate(baseAddress + yearString + "/day/" + dayString + "/input"),
        chromedp.Text(`pre`, &res),
    )
    if err != nil {
        panic("Failed to get input")
    }

    if(strings.Contains(res, "Puzzle")){
        panic("Failed to get input check authorization")
    }
    if(strings.Contains(res, "404")){
        panic("Failed to get input check address")
    }

      if _, err := os.Stat(*dirToSaveTo); os.IsNotExist(err) {
        file, err := os.Create(*dirToSaveTo)
        if err != nil {
            fmt.Println("Error creating file:", err)
            return
        }
        defer file.Close()
    }

    err = os.WriteFile(*dirToSaveTo, []byte(res), 0644)
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    err = chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
        return chromedp.Cancel(ctx)
    }))
    if err != nil {
        panic("Could not close browser session")
    }
}


func Login(ctx *context.Context, baseAddress, email, password string) error {
    err := chromedp.Run(*ctx,
        chromedp.ActionFunc(func(ctx context.Context) error {
            fmt.Println("Authenticating:", baseAddress+"auth/github")
            return nil
        }),
        chromedp.Navigate(baseAddress+"auth/github"),
        chromedp.Sleep(2*time.Second),

        chromedp.ActionFunc(func(ctx context.Context) error {
            fmt.Println("Waiting for login input to be visible")
            return nil
        }),
        chromedp.WaitVisible(`#login_field`, chromedp.ByID),

        chromedp.ActionFunc(func(ctx context.Context) error {
            fmt.Printf("Entering email \n")
            return nil
        }),
        chromedp.SendKeys(`#login_field`, email),

        chromedp.ActionFunc(func(ctx context.Context) error {
            fmt.Println("Waiting for password input to be visible")
            return nil
        }),

        chromedp.ActionFunc(func(ctx context.Context) error {
            fmt.Printf("Entering password\n")
            return nil
        }),
        chromedp.SendKeys(`#password`, password),

        chromedp.ActionFunc(func(ctx context.Context) error {
            fmt.Println("Clicking submit button")
            return nil
        }),
        chromedp.Click(`input[type="submit"]`),
        chromedp.Sleep(2*time.Second),
    )

    if err != nil{
        return fmt.Errorf("There was an error, %v", err)
    }
    return nil
}

