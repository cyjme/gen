**Auto create golang api project (gin+gorm+swagger), quickly build restful api and document.**

![gen-demo-gif](https://raw.githubusercontent.com/cyjme/gen/master/gen-demo.gif?raw=true)

## Example:
gen new blog-project

gen add api --model article --fields title:string,content:string,userId:int

gen add api -m user -f name:string,email:string,password:string


## Detail:

### step-0 install gen
go get -u github.com/cyjme/gen

go get -u github.com/swaggo/swag/cmd/swag

### step-1 create project
gen new blog-project

### step-2 edit config.yml
cd blog-project

vim config/config.yml

create database

### step-3 install dependency, recommend use `go mod`
go mod tidy

### step-4 add api
gen add api --model article --fields title:string,content:string,userId:int

### step-5 run
go run main.go

Api document: `your_url:port/swagger/index.html`

Api: `your_url:port/articles`    [post,get]

Api: `your_url:port/articles/{id}`    [get,put,patch,delete]

### more
get articles api, can use these params default:

  `your_url:port/articles?page=1&pageSize=10&query=title:like:%search_word%,user_id:=:1&order=created_at:desc,userId:asc`

  query param is an array split by `,` , the element is `{field}:{=|>|<|like}:{value}`

  order param is an array split by `,` , the element is `{field}:{asc|desc}`
