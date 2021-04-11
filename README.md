# Telegram bot written in Golang

## Run and build
Set environment variables with the following values:

    TELEGRAM_BOT_TOKEN=<your telegram token>
    PORT=<port>
    PUBLIC_URL=<url> or <your heroku app name>.herokuapp.com    # if you use heroku

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
  - [ ] Inline buttons for tomorrow and yesterday classes
  - [ ] Send whole schedule
  - [ ] Send classes and lecturers together