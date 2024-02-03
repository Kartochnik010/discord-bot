# Discord bot

## Objective
Develop a Discord bot using the Go programming language. The bot should exhibit unique functionalities that are engaging and useful for Discord users.

## About
Discord bot listens to commands once active parsing user input and asynchronously handling incoming requests.

#### List of commands
```
hello - Bot will answer with Hi! Kind of a healthcheck
yt - link to random favourite yt videos
help - print this message
```



You can also type `moris ...` to ask the bot anything. It will answer in a characteristic voice of cartoon character Moris from Madagascar.

## Usage

1. clone repo
```Bash
   git clone https://github.com/Kartochnik010/discord-bot.git
```
2. Enter work directory:
```Bash
    cd discord-bot
```
3. Tidy dependencies:
```Go
    go mod tidy
```
4. Add env variables. (Ask me for the api keys, but you can try adding yours.)
```Bash
    # .env
    BOT_TOKEN="*****"
    GPT_TOKEN="sk-******"
    GPT_MAX_TOKENS="256"
    GPT_MODEL="gpt-4"
```
5. Run the bot:
```Go
    make run
```


## Core Features

Your bot should include the following core features:

> Command Parsing: Ability to parse and respond to user commands.
- Bot parses incoming messages for the specific prefix. This is how bot knows that message is a command.
> Help Command: A command that lists all available commands and their descriptions.

1. Check if the bot is active.
2. Type `moris help`

> Asynchronous Processing: Ability to perform tasks asynchronously without blocking the main thread.

- Handlers handle incoming user requests asyncrounosly from the main thread(goroutine).

## Unique Feature Ideas

> Language Translation: Translate user messages into a specified language using an external API.
 
- You can talk to the bot as if it is a cartoon character. Powered by OpenAI's API. 
 
## 5. Documentation

> Code Documentation: Properly comment your code for readability and maintainability.
- Some people say that code should be self explanotary. Added comments in the code and documentation in this readme.

> User Guide: Write a brief user guide explaining how to interact with your bot and its features.
- `moris help` command prints out 

## 6. Submission

 - [ ] Submit your source code via Github/GitLab/etc.
 - [ ] Include a README file with setup instructions and a brief description of your bot’s functionality.

## 7. Evaluation Criteria

 - [ ] Functionality: Does the bot work as expected without errors?
 - [ ] Code Quality: Is the code well-structured, commented, and adhering to Go best practices?
 - [ ] Creativity: How innovative and useful are the bot's features?
    Documentation: Quality of the code documentation and user guide.
