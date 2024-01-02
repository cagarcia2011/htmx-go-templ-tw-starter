# HTMX, GO & Tailwind Starter

This minimal starter prepares an HTMX, GO, A-H.Templ, Echo and Tailwindcss project.

## Prerequisites

Make sure you have the following already installed:

-   [Go](https://go.dev/doc/install)
-   [Air](https://github.com/cosmtrek/air): Live reload for Go apps
-   [Templ](https://templ.guide/quick-start/installation): For building HTML with Go
-   [Tailwindcss](https://tailwindcss.com/docs/installation):

    -   For this starter I used the [standalone CLI build](https://tailwindcss.com/blog/standalone-cli) to use without Node.
    -   I added the build to my path and renamed the file to `twcss`. If you decide to name the file differently make sure to change the `Makefile` with the right name:

        ```bash
        css:
          twcss build -o static/css/tw.css --minify

        css-watch:
            twcss build -o static/css/tw.css --watch
        ```

## Run locally

Considerations to run locally:

1. Modify the `Makefile` file if you decide to go a different route with Tailwindcss, or with the folder structure.
2. Modify the `.are.toml` file to suite your needs. The configuration used in this starter was for a Mac. Be sure to check the [Air](https://github.com/cosmtrek/air) documentation for any changes needed.
3. Modify the `tailwind.config.js` if you need to change/add folders where tailwindcss will be used. For example:
    ```
    content: ["./templates/**/*.{html,js,templ,go}", "./someotherfolder/**/*.{file extensions where tailwindcss will live}"],
    ```
4. (Optional) Change your editor's tailwindcss intellisense configuration to add Templ as an included language. For example, for VS Code add the following setting to the User Settings (JSON):
    ```
     "tailwindCSS.includeLanguages": {
             "templ": "html",
         }
    ```
5. Run `make dev` on your terminal at the root of the project.

This will run the server and you can access the app at `localhost:3000` (if default host was not changed).
