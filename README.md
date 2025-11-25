SEG2105 – Lab 6
University of Ottawa

This program implements a concurrent web scraper using goroutines, channels, and a fixed-size worker pool, following the course lecture slides and lab requirements

---

## What the Program Does

- Uses a **worker pool** (fixed number of workers).
- Each worker is a **goroutine** that receives URLs from a `jobs` channel.
- Workers:
  - Fetch each URL
  - Record the **status code**
  - Measure the **size of the downloaded content**
  - Return everything through the `results` channel
- The main goroutine waits for **all** results before exiting.

---

## Code Used

The full implementation (from earlier) is in **main.go** and includes:

- `FetchResult` struct  
- `worker()` function  
- `fetchURL()` helper  
- goroutine creation  
- job distribution and result collection  

This matches the worker-pool pattern shown in the Lab 6 slides.

---

## How to Run

```bash
go run main.go
