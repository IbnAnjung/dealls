## Project Structure
      .
      ├── build                         # build & ci/cd content 
      ├   ├── http.Dockerfile 
      ├── cmd                           # main / bootstrapping services 
      ├   ├── http                      # http service 
      ├   ├   ├── config                 
      ├   ├   ├── handler                      
      ├   ├   ├── router                 
      ├   ├   ├── server.go             # load all dependency & start http service                 
      ├── database                      # initiate db migration
      ├── entity                        # entity layer
      ├   ├── enauth                
      ├   ├── anuser                
      ├── mock                          # all mock entity interface 
      ├   ├── entity                
      ├   ├── pkg                
      ├── pkg                           # helper, driver
      ├   ├── cache                
      ├   ├── crypt                
      ├   ├── error                
      ├   ├── http                
      ├   ├── jwt                
      ├   ├── orm                
      ├   ├── redis                
      ├   ├── sql                
      ├   ├── string                
      ├   ├── structvalidator                
      ├   ├── time                
      ├── repository                      # adapter to get data from source
      ├   ├── gorm                        # source SQL with gorm              
      ├   ├   ├── model                            
      ├── usecase                         # bussiness logic
      ├   ├── auth
      ├   ├── user
      ├── .env.example                  
      ├── .gitignore                   
      ├── docker-compose.yaml
      ├── go.mod
      ├── main.go
      ├── mocker.yaml
      └── README.md

## How To run This Project
```bash
# clone repository
git clone git@github.com:IbnAnjung/dealls.git

cd dealls

docker-compose up 
```

## Acknowledgements
 - [golangci-lint for linter](https://golangci-lint.run/)
 - [osv-scanner to keep dependency are safe](https://google.github.io/osv-scanner/)

