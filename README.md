![Bug Tracker Banner](banner.png)
# Bug Tracker 🐛
This repository contains the source code for a Bug Tracker web application. Built with Go and HTMX. It is currently under construction.

1. First fill out the env file (see .env.example)
2. Run `dbrocket up` to push the migrations
3. Run `air` to start the go server
4. Run `npx tailwindcss -i ./public/tailwind.css -o ./public/styles.css --watch` if you're going to modify tailwind code


## Todo

- Fix errror where if you have a cookie that appears logged in in the session, you cant log in, even if the user does not actually exist