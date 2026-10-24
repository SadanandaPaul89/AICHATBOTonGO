A terminal Based AIchatbot made using GOLang with Langchain on top of Gemini.

How to run:
clone or fork the repository: https://github.com/SadanandaPaul89/AICHATBOTonGO.git
cd LangChainBot/bot
go mod init

### Install the dependencies with the following commands:
go get github.com/joho/godotenv
go get github.com/tmc/langchaingo/llms
go get github.com/tmc/langchaingo/llms/googleai

## Next:
go mod tidy

## Initialize the environment
Use the variable given in the .example.env and create a new .env file
Paste your gemini api key obtained from [Google Ai Studio](https://aistudio.google.com/)

### Run this project
from the root, run 
go run .

# THE PROJECT IS READY

# Reference Documentation
GOLang [LangChainGo](https://pkg.go.dev/github.com/tmc/langchaingo)
