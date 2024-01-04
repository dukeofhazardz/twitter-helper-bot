## Welcome to Twitter-Helper-Bot
    Note: 
    This project is created for fun and learning using https://github.com/dghubble/go-twitter
    
    For developing a twitter bot, you should have Prerequisite of having a Twitter Developer Account. 
---

## Apply for Twitter Developer Account
    
#### https://developer.twitter.com/en/apply-for-access
---
# Essential packages
    go get github.com/dghubble/go-twitter/twitter
    go get github.com/dghubble/oauth1.
    
---
## Keys and Tokens
    Once you are accessed to the developer account, click the Keys and tokens option in the developer portal.
---
### Take a note of the keys and token
 -  Consumer Keys(API keys)
 -  Authentication Tokens   
---
    create a .env file and store the client credentials
    Note : Replace your credentials in xxx

    TWITTER_CONSUMER_KEY= xxx
    TWITTER_CONSUMER_SECRET= xxx
    TWITTER_ACCESS_TOKEN=xxx
    TWITTER_ACCESS_SECRET= xxx  
---
## How to Start Twitter-Helper Bot


- For liunx environment

**Prerequisites** 

    Golang version  go1.16.6 required

Before Run this bot, please follow the steps,

    Install golang  (if your already install golang please skip the step)
    
    Note:  User - root
   
    [Download  Golang package] (https://golang.org/doc/install)
	
   	   1. unzip the download folder 
   	   2. tar -xzf golangpackage name 
	   3. Copy go folder 
	   4. cp -rf  go /user/local
	   5. Set golang path 
 	   6. vim ~/.bashrc
	   
	   Add the below line in the bashrc 
		1. export GOROOT=/usr/local/go
		2. export GOPATH=$HOME/Goprojects
		3 .export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
		
	   7. Save the file
	   8.source  ~/.bashrc

	Check go version in comment line
		
		go version 

	 Output 
		go version go1.16.6 linux/amd64

	
     Run the program  
    
    make sure, your go project has the following project struture,
    Inside the GOPATH, Create three folders
    1. bin
    2. pkg
    3. src

    Go to src folder and clone the project repository from the https://github.com/yuvan11/go-twitter-helper-bot.git
    
    clone the repository from project URL 
    
    Run the command,
    make run

---


- For Windows 64-bit environment

**Prerequisites** 

    Golang version  go1.16.6 required

Before Run this bot, please follow the steps,

    Install golang  (if your already install golang please skip the step)
    
   
    [Download  Golang package] (https://golang.org/doc/install)
	
   	  1. Open the MSI file you downloaded and follow the prompts to install Go.
         
      2.By default, the installer will install Go to Program Files or Program Files (x86). You can change the location as needed. After installing, you will need to close and reopen any open command prompts so that changes to the environment made by the installer are reflected at the command prompt.

        3. Verify that you've installed Go.
        In Windows, click the Start menu.
        In the menu's search box, type cmd, then press the Enter key.

        In the Command Prompt window that appears, type the following command:
        $ go version
        Confirm that the command prints the installed version of Go.
	   
	   Add the below path in the environmental variable
       	1. set GOROOT = C:\Go
		2. set GOPATH = C:\Projects\Goprojects\
        3. set GOBIN =  C:\Projects\Goprojects\bin

Run the program  
    
    make sure, your go project has the following project struture,
    Inside the GOPATH, Create three folders
    1. bin
    2. pkg
    3. src

    Go to src folder and clone the project repository from the https://github.com/yuvan11/go-twitter-helper-bot.git
    
    Run the command,
    make run

---
## Specialities of this bot
    The bot is designed using [go-twitter](https://github.com/dghubble/go-twitter).

    1. posting a new tweet
    2. Fetching the current top trending hastags from region based(INDIA) 
    3. Fetching tweets of the users whoever use the trending hastags
    4. Filtering the tweet status Id of the user from the user's tweet.
    5. Reweeting the tweet with help of the filtered user status ID
---
## Benefits 
    Once you starting the bot, your bot will successfully post a new tweet and also retweeting the other users tweet with trending tags.
    
`
This project is still in development...  
`
---
## LICENSE
[MIT LICENSE](https://github.com/yuvan11/go-twitter-helper-bot/blob/master/LICENSE)




