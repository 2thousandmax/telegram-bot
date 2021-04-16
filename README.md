# Telegram bot written in Golang

## Run and build
Set environment variables with the following values:

    TELEGRAM_BOT_TOKEN=<your telegram token>
    PORT=<port>
    PUBLIC_URL=<url>
    # if you use heroku
    PUBLIC_URL=<your heroku app name>.herokuapp.com

Create `config.yaml` file in the **configs** folder. For reference see `config.yaml.sample`

To build project as binary, you can run:

    make build

# TODO
- [ ] Docker
  - [ ] Dockerfile
  - [ ] Deploy container to heroku
- [ ] Logger
- [ ] Write unit tests
- [ ] New features 
  - [x] Inline buttons for tomorrow and yesterday classes
  - [ ] Send whole schedule
  - [ ] Send classes and lecturers together