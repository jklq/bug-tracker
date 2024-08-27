![Bug Tracker Banner](banner.png)

# Bug Tracker 🐛

This repository contains the source code for a Bug Tracker web application. Built with Go and HTMX. It is currently under construction.

1. First fill out the env file (see `.env.example`)
2. Run `dbmate up` to push the migrations
3. Run `templ generate` to generate template code
4. Run `air` to start the go server
5. Run `npx tailwindcss -i ./public/tailwind.css -o ./public/styles.css --watch` if you're going to modify tailwind code

## Todo

- Fix error where if you have a cookie that appears logged in, you cant log in, even if the user does not actually exist

## Inspiration

- ["Admin dashboard: analytics UX" by Halo Product](https://dribbble.com/shots/19687516-Admin-dashboard-analytics-UX])
  ![Image of Admin Dashboard](https://cdn.dribbble.com/userupload/3831213/file/original-c8996d294ff916cb9d0e3f3991cefdb9.png?resize=1024x768)
