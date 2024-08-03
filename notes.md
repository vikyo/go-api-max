This is tutorial for developing api using go.
Udemy - Go - The Complete Guide(Maxi)

`binding: "required"` is used to tell server that the body of request must contain these feilds.
ShouldBindJSON is the function that matches the request body with the model struct and extracts it from request body.

L167:
    To add sqlite first go get it from https://github.com/mattn/go-sqlite3
    Import it, but since it is not directly added but under the hood, so we use _ symbol to tell go
    that include it in the final build.
    Since this library did not work, we will use go get modernc.org/sqlite as suggested in https://www.udemy.com/course/go-the-complete-guide/learn/lecture/40959806#questions/20950586

L170:
    If you have query that fetches data the use .Query, if updates/creates data then .Exec
    To see sqlite tables in vscode install extn: SQLite.

L172:
    Query() gives rows, QueryRow gives one single row.

go.mod vs go.sum:
    https://golangbyexample.com/go-mod-sum-module/