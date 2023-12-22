# Pixelart

This little project is a homage to `r\place`.

It's an API written in Go which can be used to create your own pixel-placing event.
Over this API you can create a pixel canvas and send requests to place single pixels using coordinates and RGB values on it.

Example using the sample frontend and Python script:

![Sample picture rendering](/example/sample-picture-rendering.gif)

## Structure

- `cmd` & `internal`: Pixelart API - Written in Go
- `frontend`: Sample frontend using HTML, CSS and JavaScript
- `example`: Example Python script which reads a picture and sends it pixel by pixel to the API

## Run the API

1. `cd` into the `backend` folder
2. Run `go mod tidy` to install all needed dependencies
3. Edit `.env.list` file according to your needs:
    - Adapt canvas width and height
    - Specify if you want to use a backup file; if you want set the value to `true`
    - Specify the path where backup files should be saved to; if location does not exist, folders will be generated
automatically
    - For debugging, you can change Gin's mode
4. Run `go run cmd/pixelart/main.go`
5. The API and webserver is now running - Open `localhost:8080/web` to access the frontend

## Using the API & Endpoints

- `/web` (GET)
  - Method: **GET**
  - Get the frontend in this project - Of course the code can be modified, so you can use your own frontend.
At the moment it's a very simple HTML canvas


- `/pixels`
  - Method: **GET**
  - Get the canvas information. The response contains the size of the canvas and a three-dimensional array for the canvas representation.
    - First dimension: Rows
    - Second dimension: Column
    - Third dimension: RGB values for a pixel
  - Example response (3x3 canvas):
```json
{
    "width": 3,
    "height": 3,
    "pixels": [
        [
            [0, 0, 0],
            [0, 0, 0],
            [0, 0, 0]
        ],
        [
            [0, 0, 0],
            [0, 0, 0],
            [0, 0, 0]
        ],
        [
            [0, 0, 0],
            [0, 0, 0],
            [0, 0, 0]
        ]
    ]
}
```


- `/pixel`
  - Method: **POST**
  - Modify the color of a single pixel - Coordinates are zero-indexed
  - Example request body (Set pixel color at (0|2) to green):
```json
{
    "x": 0,
    "y": 2,
    "rgb": [0, 255, 0]
}
```
