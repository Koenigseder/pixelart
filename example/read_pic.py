import time

from PIL import Image
import requests


def read_pic():
    im = Image.open('beer.png')
    width, height = im.size

    for y in range(height):
        for x in range(width):
            r, g, b, a = im.getpixel((y, x))

            if a == 0:
                continue

            requests.post("http://localhost:8080/api/pixel", json={
                "x": x,
                "y": y,
                "rgb": [r, g, b]
            })

            time.sleep(0.1)


def main():
    read_pic()


if __name__ == "__main__":
    main()
