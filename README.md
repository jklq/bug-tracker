# template for go + htmx
1. First fill out the env file (see .env.example)
2. Run `dbrocket up` to push the migrations
3. Run `air` to start the go server
4. Run `npx tailwindcss -i ./public/tailwind.css -o ./public/styles.css --watch` if you're going to modify tailwind code