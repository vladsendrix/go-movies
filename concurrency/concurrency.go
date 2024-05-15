package concurrency

import (
    "database/sql"
    "log"
    "sync"
    "time"
)

func DirtyRead(db *sql.DB) {
    var wg sync.WaitGroup
    wg.Add(2)

    // Transaction 1: Update a movie's release year
    go func() {
        defer wg.Done()

        tx, err := db.Begin()
        if err != nil {
            log.Fatal(err)
        }

        _, err = tx.Exec("UPDATE movies SET release_year = 2022 WHERE director_id = 1")
        if err != nil {
            log.Fatal(err)
        }

        // Don't commit the transaction yet
    }()

    // Transaction 2: Read the movie's release year
    go func() {
        defer wg.Done()

        tx, err := db.Begin()
        if err != nil {
            log.Fatal(err)
        }

        var releaseYear int
        err = tx.QueryRow("SELECT release_year FROM movies WHERE director_id = 1").Scan(&releaseYear)
        if err != nil {
            log.Fatal(err)
        }

        log.Println("Release Year:", releaseYear)

        if err := tx.Commit(); err != nil {
            log.Fatal(err)
        }
    }()

    wg.Wait()
}

func LostUpdate(db *sql.DB) {
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        tx, err := db.Begin()
        if err != nil {
            log.Fatal(err)
        }
        _, err = tx.Exec("UPDATE movies SET rating = rating + 1 WHERE id = 1")
        if err != nil {
            log.Fatal(err)
        }
        tx.Commit()
    }()

    go func() {
        defer wg.Done()
        tx, err := db.Begin()
        if err != nil {
            log.Fatal(err)
        }
        _, err = tx.Exec("UPDATE movies SET rating = rating + 1 WHERE id = 1")
        if err != nil {
            log.Fatal(err)
        }
        tx.Commit()
    }()

    wg.Wait()
}

func PhantomReads(db *sql.DB) {
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        tx, err := db.Begin()
        if err != nil {
            log.Fatal(err)
        }
        rows, err := tx.Query("SELECT * FROM movies WHERE rating > 8")
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()

        count := 0
        for rows.Next() {
            count++
        }
        if err = rows.Err(); err != nil {
            log.Fatal(err)
        }

        log.Println("Number of movies with rating > 8:", count)
        tx.Commit()
    }()

    go func() {
        defer wg.Done()
        tx, err := db.Begin()
        if err != nil {
            log.Fatal(err)
        }
        _, err = tx.Exec("INSERT INTO movies (id, title, rating) VALUES (999, 'New Movie', 9)")
        if err != nil {
            log.Fatal(err)
        }
        tx.Commit()
    }()

    wg.Wait()
}

func UnrepeatableReads(db *sql.DB) {
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        tx, err := db.Begin()
        if err != nil {
            log.Fatal(err)
        }
        var rating int
        err = tx.QueryRow("SELECT rating FROM movies WHERE id = 1").Scan(&rating)
        if err != nil {
            log.Fatal(err)
        }
        log.Println("Initial rating:", rating)
        time.Sleep(2 * time.Second)
        err = tx.QueryRow("SELECT rating FROM movies WHERE id = 1").Scan(&rating)
        if err != nil {
            log.Fatal(err)
        }
        log.Println("Final rating:", rating)
        tx.Commit()
    }()

    go func() {
        defer wg.Done()
        tx, err := db.Begin()
        if err != nil {
            log.Fatal(err)
        }
        _, err = tx.Exec("UPDATE movies SET rating = rating + 1 WHERE id = 1")
        if err != nil {
            log.Fatal(err)
        }
        tx.Commit()
    }()

    wg.Wait()
}

func TestConcurrency(db *sql.DB) {
    DirtyRead(db)
    LostUpdate(db)
    PhantomReads(db)
    UnrepeatableReads(db)
}