// Go in Action
// @jeffotoni
// 2019-03-10

package main

import (
    "strconv"

    "github.com/jeffotoni/gologs/repo/postgres"
)

func worker(id int, jobs <-chan string, results chan<- string) {
    for j := range jobs {
        //fmt.Println("worker", id, "started  job", j)
        //time.Sleep(time.Second)
        //fmt.Println("worker", id, "finished job", j)
        results <- j
    }
}

func main() {

    produc := "1"
    jobs := make(chan string, 500000)
    results := make(chan string, 500000)

    for w := 1; w <= 2000; w++ {
        go worker(w, jobs, results)
    }

    for i := 1; i <= 500000; i++ {
        jobs <- `{"versão": "1.1","host": "exemplo.org","key":"producer_` + produc + `_` + strconv.Itoa(i) + `","level":"info","project":"my-project-here","short_message":"one msg here...","nível": 5,"some_info":"foo jeff"}`
    }
    close(jobs)

    // Finally we collect all the results of the work.
    for a := 1; a <= 500000; a++ {
        cmsgJson := <-results
        postgres.Insert5Log(cmsgJson)
    }
}
